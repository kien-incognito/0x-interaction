pragma solidity ^0.5.9;
pragma experimental ABIEncoderV2;

import "@0x/contracts-exchange-forwarder/contracts/src/interfaces/IForwarder.sol";
import "@0x/contracts-utils/contracts/src/LibBytes.sol";

contract Token {
    function balanceOf(address addr) public view returns (uint256);
    function transfer(address recipient, uint amount) public;
    function approve(address spender, uint tokens) public returns (bool);
    function allowance(address tokenOwner, address spender) public view returns (uint);
}

contract SimpleTokenSwapContract
{
    using LibBytes for bytes;

    function() external payable {}

    address internal OWNER;
    constructor ()
        public
    {
        OWNER = msg.sender;
    }
    
    function trigger0x(bytes memory callDataHex, address _forwarder) public payable {
        IForwarder forwarder = IForwarder(_forwarder);
        (bool success, ) = address(forwarder).call.value(msg.value)(callDataHex);
        require(success);
    }

    function balanceOf(address tokenAddress) public view returns (uint256) {
        return Token(tokenAddress).balanceOf(address(this));
    }

    function allowance(address tokenAddress, address proxy) public view returns (uint) {
        return Token(tokenAddress).allowance(address(this), proxy);
    }

    function approve(address tokenAddress, address proxy, uint amount) public {
        Token(tokenAddress).approve(proxy, amount);
    }
}
