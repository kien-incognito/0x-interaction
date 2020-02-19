pragma solidity ^0.5.12;
pragma experimental ABIEncoderV2;

import "./IERC20.sol";
import "./pause.sol";

/**
 * @dev Interface of the contract capable of checking if an instruction is
 * confirmed over at Incognito Chain
 */
contract Incognito {
    function instructionApproved(
        bool,
        bytes32,
        uint,
        bytes32[] memory,
        bool[] memory,
        bytes32,
        bytes32,
        uint[] memory,
        uint8[] memory,
        bytes32[] memory,
        bytes32[] memory
    ) public view returns (bool);
}

/**
 * @dev Interface of the previous Vault contract to query burn proof status
 */
contract Withdrawable {
    function isWithdrawed(bytes32) public view returns (bool);
}

/**
 * @dev Responsible for holding the assets and issue minting instruction to
 * Incognito Chain. Also, when presented with a burn proof created over at
 * Incognito Chain, releases the tokens back to user
 */
contract IncognitoMode is AdminPausable {
    address constant public ETH_TOKEN = 0x0000000000000000000000000000000000000000;
    mapping(bytes32 => bool) public withdrawed;

    // withdrawRequests store pending withdrawn requests which specify how many token amount that an address can withdraw.
    mapping(address => mapping(address => uint)) public withdrawRequests;

    Incognito public incognito;
    Withdrawable public prevVault;
    address payable public newVault;

    event Deposit(address token, string incognitoAddress, uint amount);
    event Withdraw(address token, address to, uint amount);
    event WithdrawRequest(address token, string incognitoAddress, uint amount);
    event Migrate(address newVault);
    event MoveAssets(address[] assets);
    event UpdateIncognitoProxy(address newIncognitoProxy);
    event Trade(string incognitoAddress, address token, uint amount);
    
    struct BurnInst {
        uint8 flag;
        uint8 meta;
        uint8 shard;
        address token;
        address payable to;
        uint burned;
    }

    /**
     * @dev Creates new Vault to hold assets for Incognito Chain
     * @param admin: authorized address to Pause and migrate contract
     * @param incognitoProxyAddress: contract containing Incognito's committees
     * @param _prevVault: previous version of the Vault to refer back if necessary
     * After migrating all assets to a new Vault, we still need to refer
     * back to previous Vault to make sure old withdrawals aren't being reused
     */
    constructor(address admin, address incognitoProxyAddress, address _prevVault) public AdminPausable(admin) {
        incognito = Incognito(incognitoProxyAddress);
        prevVault = Withdrawable(_prevVault);
        newVault = address(0);
    }

    /**
     * @dev Makes a ETH deposit to the vault to mint pETH over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice The maximum amount to deposit is capped since Incognito balance is stored as uint64
     * @param incognitoAddress: Incognito Address to receive pETH
     */
    function deposit(string memory incognitoAddress) public payable isNotPaused {
        // require((msg.value + address(this).balance) <= 10 ** 27, "Balance of this contract has been reaching to its uint's maximum.");
        require(address(this).balance <= 10 ** 27);
        emit Deposit(ETH_TOKEN, incognitoAddress, msg.value);
    }

    /**
     * @dev Makes a ERC20 deposit to the vault to mint pERC20 over at Incognito Chain
     * @notice This only works when the contract is not Paused
     * @notice The maximum amount to deposit is capped since Incognito balance is stored as uint64
     * @notice Before calling this function, enough ERC20 must be allowed to
     * tranfer from msg.sender to this contract
     * @param token: address of the ERC20 token
     * @param amount: to deposit to the vault and mint on Incognito Chain
     * @param incognitoAddress: Incognito Address to receive pERC20
     */
    function depositERC20(address token, uint amount, string memory incognitoAddress) public payable isNotPaused {
        ERC20 erc20Interface = ERC20(token);
        uint8 decimals = getDecimals(address(token));
        uint tokenBalance = erc20Interface.balanceOf(address(this));
        uint emitAmount = amount;
        if (decimals > 9) {
            emitAmount = emitAmount / (10 ** (uint(decimals) - 9));
            tokenBalance = tokenBalance / (10 ** (uint(decimals) - 9));
        }
        require(emitAmount <= 10 ** 18 && tokenBalance <= 10 ** 18 && emitAmount + tokenBalance <= 10 ** 18);

        erc20Interface.transferFrom(msg.sender, address(this), amount);
        require(checkSuccess());
        emit Deposit(token, incognitoAddress, emitAmount);
    }

    /**
     * @dev Checks if a burn proof has been used before
     * @notice First, we check inside the storage of this contract itself. If the
     * hash has been used before, we return the result. Otherwise, we query
     * previous vault recursively until the first Vault (prevVault address is 0x0)
     * @param hash: of the burn proof
     * @return bool: whether the proof has been used or not
     */
    function isWithdrawed(bytes32 hash) public view returns(bool) {
        if (withdrawed[hash]) {
            return true;
        } else if (address(prevVault) == address(0)) {
            return false;
        }
        return prevVault.isWithdrawed(hash);
    }

    /**
     * @dev Parses a burn instruction and returns individual components
     * @param inst: the full instruction, containing both metadata and body
     * @return flag: 
     * @return meta: type of the instruction, 72 for burning instruction
     * @return shard: ID of the Incognito shard containing the instruction, must be 1
     * @return token: ETH address of the token contract (0x0 for ETH)
     * @return to: ETH address of the receiver of the token
     * @return amount: of tokens to return
     */
    function parseBurnInst(bytes memory inst) public pure returns (uint8, uint8, address, address payable, uint) {
        uint8 meta = uint8(inst[0]);
        uint8 shard = uint8(inst[1]);
        address token;
        address payable to;
        uint amount;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x22)) // [3:34]
            to := mload(add(inst, 0x42)) // [34:66]
            amount := mload(add(inst, 0x62)) // [66:98]
        }
        return (meta, shard, token, to, amount);
    }
    
    /**
     * @dev returns BurnInst object. This function is used in submitBurnProof function.
     * @param inst: the full instruction, containing both metadata and body
     */
    function getBurnInst(bytes memory inst) internal pure returns (BurnInst memory) {
        uint8 flag = uint8(inst[0]);
        uint8 meta = uint8(inst[1]);
        uint8 shard = uint8(inst[2]);
        address token;
        address payable to;
        uint amount;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x23)) // [3:35]
            to := mload(add(inst, 0x43)) // [35:67]
            amount := mload(add(inst, 0x63)) // [67:99]
        }
        BurnInst memory burnInst = BurnInst({flag: flag, meta: meta, shard: shard, token: token, to: to, burned: amount});
        return burnInst;
    }

    /**
     * @dev Verifies that a burn instruction is valid
     * @notice All params except inst are the list of 2 elements corresponding to
     * the proof on beacon and bridge
     * @notice All params are the same as in `withdraw`
     */
    function verifyInst(
        bytes memory inst,
        uint[2] memory heights,
        bytes32[][2] memory instPaths,
        bool[][2] memory instPathIsLefts,
        bytes32[2] memory instRoots,
        bytes32[2] memory blkData,
        uint[][2] memory sigIdxs,
        uint8[][2] memory sigVs,
        bytes32[][2] memory sigRs,
        bytes32[][2] memory sigSs
    ) internal {
        // Each instruction can only by redeemed once
        bytes32 instHash = keccak256(inst);
        bytes32 beaconInstHash = keccak256(abi.encodePacked(inst, heights[0]));
        bytes32 bridgeInstHash = keccak256(abi.encodePacked(inst, heights[1]));
        require(!isWithdrawed(instHash));

        // Verify instruction on beacon
        require(incognito.instructionApproved(
            true,
            beaconInstHash,
            heights[0],
            instPaths[0],
            instPathIsLefts[0],
            instRoots[0],
            blkData[0],
            sigIdxs[0],
            sigVs[0],
            sigRs[0],
            sigSs[0]
        ));

        // Verify instruction on bridge
        require(incognito.instructionApproved(
            false,
            bridgeInstHash,
            heights[1],
            instPaths[1],
            instPathIsLefts[1],
            instRoots[1],
            blkData[1],
            sigIdxs[1],
            sigVs[1],
            sigRs[1],
            sigSs[1]
        ));
        withdrawed[instHash] = true;
    }

    /**
     * @dev Withdraws pETH/pERC20 by providing a burn proof over at Incognito Chain
     * @notice This function takes a burn instruction on Incognito Chain, checks
     * for its validity and returns the token back to ETH chain
     * @notice This only works when the contract is not Paused
     * @notice All params except inst are the list of 2 elements corresponding to
     * the proof on beacon and bridge
     * @param inst: the decoded instruction as a list of bytes
     * @param heights: the blocks containing the instruction
     * @param instPaths: merkle path of the instruction
     * @param instPathIsLefts: whether each node on the path is the left or right child
     * @param instRoots: root of the merkle tree contains all instructions
     * @param blkData: merkle has of the block body
     * @param sigIdxs: indices of the validators who signed this block
     * @param sigVs: part of the signatures of the validators
     * @param sigRs: part of the signatures of the validators
     * @param sigSs: part of the signatures of the validators
     */
    function withdraw(
        bytes memory inst,
        uint[2] memory heights,
        bytes32[][2] memory instPaths,
        bool[][2] memory instPathIsLefts,
        bytes32[2] memory instRoots,
        bytes32[2] memory blkData,
        uint[][2] memory sigIdxs,
        uint8[][2] memory sigVs,
        bytes32[][2] memory sigRs,
        bytes32[][2] memory sigSs
    ) public isNotPaused {
        (uint8 meta, uint8 shard, address token, address payable to, uint burned) = parseBurnInst(inst);
        require(meta == 72 && shard == 1); // Check instruction type

        // Check if balance is enough
        if (token == ETH_TOKEN) {
            require(address(this).balance >= burned);
        } else {
            uint8 decimals = getDecimals(token);
            if (decimals > 9) {
                burned = burned * (10 ** (uint(decimals) - 9));
            }
            require(ERC20(token).balanceOf(address(this)) >= burned);
        }

        verifyInst(
            inst,
            heights,
            instPaths,
            instPathIsLefts,
            instRoots,
            blkData,
            sigIdxs,
            sigVs,
            sigRs,
            sigSs
        );

        // Send and notify
        if (token == ETH_TOKEN) {
            to.transfer(burned);
        } else {
            ERC20(token).transfer(to, burned);
            require(checkSuccess());
        }
        emit Withdraw(token, to, burned);
    }

    /**
     * @dev Burnt Proof is submited to store burnt amount of p-token/p-ETH and receiver's address
     * Receiver then can call withdrawRequest to withdraw these token to he/she incognito address.
     * @notice This function takes a burn instruction on Incognito Chain, checks
     * for its validity and returns the token back to ETH chain
     * @notice This only works when the contract is not Paused
     * @notice All params except inst are the list of 2 elements corresponding to
     * the proof on beacon and bridge
     * @param inst: the decoded instruction as a list of bytes
     * @param heights: the blocks containing the instruction
     * @param instPaths: merkle path of the instruction
     * @param instPathIsLefts: whether each node on the path is the left or right child
     * @param instRoots: root of the merkle tree contains all instructions
     * @param blkData: merkle has of the block body
     * @param sigIdxs: indices of the validators who signed this block
     * @param sigVs: part of the signatures of the validators
     * @param sigRs: part of the signatures of the validators
     * @param sigSs: part of the signatures of the validators
     */
    function submitBurnProof(
        bytes memory inst,
        uint[2] memory heights,
        bytes32[][2] memory instPaths,
        bool[][2] memory instPathIsLefts,
        bytes32[2] memory instRoots,
        bytes32[2] memory blkData,
        uint[][2] memory sigIdxs,
        uint8[][2] memory sigVs,
        bytes32[][2] memory sigRs,
        bytes32[][2] memory sigSs
    ) public isNotPaused {
        BurnInst memory burnInst = getBurnInst(inst);
        require(burnInst.meta == 72 && burnInst.shard == 1); // Check instruction type
        // Check if balance is enough
        if (burnInst.token == ETH_TOKEN) {
            require(address(this).balance >= burnInst.burned);
        } else {
            uint8 decimals = getDecimals(burnInst.token);
            if (decimals > 9) {
                burnInst.burned = burnInst.burned * (10 ** (uint(decimals) - 9));
            }
            require(ERC20(burnInst.token).balanceOf(address(this)) >= burnInst.burned);
        }

        verifyInst(
            inst,
            heights,
            instPaths,
            instPathIsLefts,
            instRoots,
            blkData,
            sigIdxs,
            sigVs,
            sigRs,
            sigSs
        );

        withdrawRequests[burnInst.to][burnInst.token] += burnInst.burned;
    }
    
    /**
     * @dev generate address from signature data and hash.
     */
    function sigToAddress(bytes memory signData, bytes32 hash) public pure returns (address) {
        bytes32 s;
        bytes32 r;
        uint8 v;
        assembly {
            r := mload(add(signData, 0x20))
            s := mload(add(signData, 0x40))
        }
        v = uint8(signData[64]) + 27;
        return ecrecover(hash, v, r, s);
    }

    /**
     * @dev User requests withdraw token contains in withdrawRequests.
     * WithdrawRequest event will be emitted to let incognito recognize and mint new p-tokens for the user.
     * @param incognitoAddress: incognito's address that will receive minted p-tokens.
     */
    function requestWithdraw(string memory incognitoAddress, address token, uint amount, bytes memory signData, bytes32 hash) public {
        address verifier = sigToAddress(signData, hash);
        require(withdrawRequests[verifier][token] >= amount);
        withdrawRequests[verifier][token] -= amount;
        emit Deposit(token, incognitoAddress, amount);
    }

    /**
     * @dev execute is used when users want to trade their tokens to other tokens.
     * @param recipientToken: received token address.
     * @param exchangeAddress: exchange address that execute trade process.
     * @param callData: encoded with signature and params of trade function.
     * @param hash:
     * @param signData:
     */
    function execute(
        address token,
        uint amount,
        address recipientToken,
        address exchangeAddress,
        bytes memory callData,
        bytes32 hash,
        bytes memory signData
    ) public payable {
        address verifier = sigToAddress(signData, hash);
        
        require(withdrawRequests[verifier][token] >= amount);
        require(token != recipientToken);

        // define number of eth spent for forwarder.
        uint ethAmount = msg.value;
        if (token == ETH_TOKEN) {
            ethAmount += amount;
        } else {
            // transfer token to exchangeAddress.
            require(ERC20(token).balanceOf(address(this)) >= amount);
            require(ERC20(token).transfer(exchangeAddress, amount));
        }

        trade(recipientToken, ethAmount, callData, exchangeAddress);

        // update withdrawRequests
        withdrawRequests[verifier][token] -= amount;
        withdrawRequests[verifier][recipientToken] += amount;
    }
    
    function trade(address recipientToken, uint ethAmount, bytes memory callData, address exchangeAddress) internal {
         // get balance of recipient token before trade to compare after trade.
        uint balanceBeforeTrade = balanceOf(recipientToken);
        if (recipientToken == ETH_TOKEN) {
            balanceBeforeTrade -= msg.value;
        }
        require(address(this).balance >= ethAmount);
        (bool success, bytes memory result) = exchangeAddress.call.value(ethAmount)(callData);

        require(success);
        (address returnedTokenAddress, uint returnedAmount) = abi.decode(result, (address, uint));

        require(returnedTokenAddress == recipientToken && balanceOf(recipientToken) - balanceBeforeTrade == returnedAmount);
    }

    /**
     * NOTE: This function is used for testing purpose only, remove/comment this function after used.
     */
    function setAmount(address verifier, address sellToken, uint amount) public {
        withdrawRequests[verifier][sellToken] = amount;
    }

    /**
     * @dev Saves the address of the new Vault to migrate assets to
     * @notice In case of emergency, Admin will Pause the contract, shutting down
     * all incoming transactions. After a new contract with the fix is deployed,
     * they will migrate assets to it and allow normal operations to resume
     * @notice This only works when the contract is Paused
     * @notice This can only be called by Admin
     * @param _newVault: address to save
     */
    function migrate(address payable _newVault) public onlyAdmin isPaused {
        require(_newVault != address(0));
        newVault = _newVault;
        emit Migrate(_newVault);
    }

    /**
     * @dev Move some assets to newVault
     * @notice This only works when the contract is Paused
     * @notice This can only be called by Admin
     * @param assets: address of the ERC20 tokens to move, 0x0 for ETH
     */
    function moveAssets(address[] memory assets) public onlyAdmin isPaused {
        require(newVault != address(0));
        for (uint i = 0; i < assets.length; i++) {
            if (assets[i] == ETH_TOKEN) {
                newVault.transfer(address(this).balance);
            } else {
                uint bal = ERC20(assets[i]).balanceOf(address(this));
                if (bal > 0) {
                    ERC20(assets[i]).transfer(newVault, bal);
                    require(checkSuccess());
                }
            }
        }
        emit MoveAssets(assets);
    }

    /**
     * @dev Changes the IncognitoProxy to use
     * @notice If the IncognitoProxy contract malfunctioned, Admin could config
     * the Vault to use a new fixed IncognitoProxy contract
     * @notice This only works when the contract is Paused
     * @notice This can only be called by Admin
     * @param newIncognitoProxy: address of the new contract
     */
    function updateIncognitoProxy(address newIncognitoProxy) public onlyAdmin isPaused {
        require(newIncognitoProxy != address(0));
        incognito = Incognito(newIncognitoProxy);
        emit UpdateIncognitoProxy(newIncognitoProxy);
    }

    /**
     * @dev Payable Fallback function to receive Ether from oldVault when migrating
     */
    function () external payable {}

    /**
     * @dev Check if transfer() and transferFrom() of ERC20 succeeded or not
     * This check is needed to fix https://github.com/ethereum/solidity/issues/4116
     * This function is copied from https://github.com/AdExNetwork/adex-protocol-eth/blob/master/contracts/libs/SafeERC20.sol
     */
    function checkSuccess() private pure returns (bool) {
        uint256 returnValue = 0;

        assembly {
            // check number of bytes returned from last function call
            switch returndatasize

            // no bytes returned: assume success
            case 0x0 {
                returnValue := 1
            }

            // 32 bytes returned: check if non-zero
            case 0x20 {
                // copy 32 bytes into scratch space
                returndatacopy(0x0, 0x0, 0x20)

                // load those bytes into returnValue
                returnValue := mload(0x0)
            }

            // not sure what was returned: don't mark as success
            default { }
        }
        return returnValue != 0;
    }

    /**
     * @dev Get the decimals of an ERC20 token, return 0 if it isn't defined
     * We check the returndatasize to covert both cases that the token has
     * and doesn't have the function decimals()
     */
    function getDecimals(address token) public view returns (uint8) {
        ERC20 erc20 = ERC20(token);
        uint256 returnValue = 0;
        erc20.decimals();

        assembly {
            // check number of bytes returned from last function call
            switch returndatasize

            // no bytes returned: zero decimals
            case 0x0 {
                returnValue := 0
            }

            // For uint8, the returndatasize is still 32 bytes
            // 32 bytes returned: check if non-zero
            case 0x20 {
                // copy 32 bytes into scratch space
                returndatacopy(0x0, 0x0, 0x20)

                // load those bytes into returnValue
                returnValue := mload(0x0)
            }

            // not sure what was returned: don't mark as success
            default {
                returnValue := 0
            }
        }
        return uint8(returnValue);
    }

    function balanceOf(address token) public view returns (uint) {
        if (token == ETH_TOKEN) {
            return address(this).balance;
        }
        return ERC20(token).balanceOf(address(this));
    }
}
