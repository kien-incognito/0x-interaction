pragma solidity >= 0.5.12;

import './IERC20.sol';

contract TradeUtils {
	ERC20 constant public ETH_CONTRACT_ADDRESS = ERC20(0x0000000000000000000000000000000000000000);
	address payable public incognitoSmartContract;
	
	modifier isIncognitoSmartContract {
	    require(msg.sender == incognitoSmartContract);
	    _;
	}
	
	// fallback function is used to receive eth.
	function() external payable {}
	
	function balanceOf(ERC20 token) internal view returns (uint256) {
		if (token == ETH_CONTRACT_ADDRESS) {
			return address(this).balance;
		}
        return token.balanceOf(address(this));
    }

	function transfer(ERC20 token, uint amount) internal {
		if (token == ETH_CONTRACT_ADDRESS) {
			require(address(this).balance >= amount);
			incognitoSmartContract.transfer(amount);
		} else {
			token.transfer(incognitoSmartContract, amount);
		}
	}

	function approve(ERC20 token, address proxy, uint amount) internal {
		if (token != ETH_CONTRACT_ADDRESS) {
			require(token.approve(proxy, 0));
			require(token.approve(proxy, amount));
		}
	}
}
