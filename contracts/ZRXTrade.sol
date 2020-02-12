pragma solidity ^0.5.12;
pragma experimental ABIEncoderV2;

import "@0x/contracts-exchange-forwarder/contracts/src/interfaces/IForwarder.sol";
import "@0x/contracts-utils/contracts/src/LibBytes.sol";
import "./utils.sol";
import "./IERC20.sol";

contract WETH {
    function withdraw(uint wad) public;
}

contract ZRXTrade is TradeUtils
{
    using LibBytes for bytes;

    function() external payable {}

    address zeroProxy;
    address wETH;

    constructor (address _wETH, address _zeroProxy, address payable _incognitoSmartContract)
        public
    {
        zeroProxy = _zeroProxy;
        incognitoSmartContract = _incognitoSmartContract;
        wETH = _wETH;
    }

    function trade(ERC20 srcToken, uint amount, ERC20 destToken, bytes memory callDataHex, address _forwarder) public payable isIncognitoSmartContract returns (uint sentAmount) {
        // do approve if srcToken is not ETH
        approve(srcToken, zeroProxy, amount);

        // trigger 0x forwarder.
        IForwarder forwarder = IForwarder(_forwarder);
        (bool success, ) = address(forwarder).call.value(msg.value)(callDataHex);
        require(success);

        // if destToken is ETH_CONTRACT_ADDRESS then unwrap WETH back to ETH
        if (destToken == ETH_CONTRACT_ADDRESS) {
            sentAmount = getBalance(ETH_CONTRACT_ADDRESS);
            withdrawWrapETH(sentAmount);
            transfer(ETH_CONTRACT_ADDRESS, sentAmount);
        } else {
            sentAmount = getBalance(destToken);
            transfer(destToken, sentAmount);
        }
        return sentAmount;
    }

    function getBalance(ERC20 token) internal view returns (uint) {
        if (token == ETH_CONTRACT_ADDRESS) {
            return balanceOf(ERC20(wETH));
        }
        return balanceOf(token);
    }

    function withdrawWrapETH(uint amount) public {
        WETH(wETH).withdraw(amount);
    }
}
