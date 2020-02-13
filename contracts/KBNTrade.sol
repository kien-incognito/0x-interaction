pragma solidity ^0.5.12;

import './utils.sol';
import './IERC20.sol';

contract KyberNetwork {
    function trade(ERC20 src, uint srcAmount, ERC20 dest, address destAddress, uint maxDestAmount, uint minConversionRate, address walletId) public payable returns(uint);
    function swapTokenToToken(ERC20 src, uint srcAmount, ERC20 dest, uint minConversionRate) public returns(uint);
    function swapEtherToToken(ERC20 token, uint minConversionRate) public payable returns(uint);
    function swapTokenToEther(ERC20 token, uint srcAmount, uint minConversionRate) public returns(uint);
    function getExpectedRate(ERC20 src, ERC20 dest, uint srcQty) public view returns(uint expectedRate, uint slippageRate);
}

contract KBNTrade is TradeUtils {
    // Variables
    KyberNetwork public kyberNetworkProxyContract;
    ERC20 constant KYBER_ETH_TOKEN_ADDRESS = ERC20(0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE);

    // Functions
    /**
     * @dev Contract constructor
     * @param _kyberNetworkProxyContract KyberNetworkProxy contract address
     */
    constructor(KyberNetwork _kyberNetworkProxyContract, address payable _incognitoSmartContract) public {
        kyberNetworkProxyContract = _kyberNetworkProxyContract;
        incognitoSmartContract = _incognitoSmartContract;
    }
    
    // fallback function which allows transfer eth.
    function() external payable {}

    /**
     * @dev Gets the conversion rate for the destToken given the srcQty.
     * @param srcToken source token contract address
     * @param srcQty amount of source tokens
     * @param destToken destination token contract address
     */
    function getConversionRates(ERC20 srcToken, uint srcQty, ERC20 destToken) public view returns (uint, uint) {
        return kyberNetworkProxyContract.getExpectedRate(srcToken, destToken, srcQty);
    }

    function trade(ERC20 srcToken, uint srcQty, ERC20 destToken) public payable isIncognitoSmartContract returns (uint) {
        require(balanceOf(srcToken) >= srcQty);
        require(srcToken != destToken);
        uint amount = 0;
        if (srcToken != ETH_CONTRACT_ADDRESS) {
            // approve
            approve(srcToken, address(kyberNetworkProxyContract), srcQty);
            if (destToken != ETH_CONTRACT_ADDRESS) { // token to token.
                require(tokenToToken(srcToken, srcQty, destToken) > 0);
            } else {
                require(tokenToEth(srcToken, srcQty) > 0);
            }
        } else {
            require(ethToToken(destToken, srcQty) > 0);
        }
        // transfer back to incognito smart contract
        amount = balanceOf(destToken);
        transfer(destToken, amount);
        return amount;
    }
    
    function ethToToken(ERC20 token, uint srcQty) internal returns (uint) {
        // Get the minimum conversion rate
        require(address(this).balance >= srcQty);
        (uint minConversionRate,) = kyberNetworkProxyContract.getExpectedRate(KYBER_ETH_TOKEN_ADDRESS, token, msg.value);
        return kyberNetworkProxyContract.swapEtherToToken.value(srcQty)(token, minConversionRate);
    }
    
    function tokenToEth(ERC20 token, uint amount) internal returns (uint) {
        (uint minConversionRate,) = kyberNetworkProxyContract.getExpectedRate(token, KYBER_ETH_TOKEN_ADDRESS, amount);
        return kyberNetworkProxyContract.swapTokenToEther(token, amount, minConversionRate);
    }
    
    function tokenToToken(ERC20 srcToken, uint srcQty, ERC20 destToken) internal returns (uint) {
        (uint minConversionRate,) = kyberNetworkProxyContract.getExpectedRate(srcToken, destToken, srcQty);
        return kyberNetworkProxyContract.swapTokenToToken(srcToken, srcQty, destToken, minConversionRate);
    }
}
