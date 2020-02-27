pragma solidity >= 0.5.2 <= 0.6.2;

import './utils.sol';
import './IERC20.sol';

contract CTokenInterface {
    
    uint public totalReserves;
    
    function transfer(address dst, uint amount) external returns (bool);
    function transferFrom(address src, address dst, uint amount) external returns (bool);
    function approve(address spender, uint amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function balanceOfUnderlying(address owner) external returns (uint);
    function getAccountSnapshot(address account) external view returns (uint, uint, uint, uint);
    function borrowRatePerBlock() external view returns (uint);
    function supplyRatePerBlock() external view returns (uint);
    function totalBorrowsCurrent() external view returns (uint);
    function borrowBalanceCurrent(address account) external returns (uint);
    function borrowBalanceStored(address account) public view returns (uint);
    function exchangeRateCurrent() public returns (uint);
    function exchangeRateStored() public view returns (uint);
    function getCash() external view returns (uint);
    function accrueInterest() public returns (uint);
    function seize(address liquidator, address borrower, uint seizeTokens) external returns (uint);
}

contract CEther is CTokenInterface {
	function mint() public payable;
	function redeem(uint redeemTokens) external returns (uint);
	function borrow(uint borrowAmount) external returns (uint);
	function repayBorrow() external payable;
}

contract CErc20 is CTokenInterface {
	address public underlying;
	function mint(uint mintAmount) external returns (uint);
    function redeem(uint redeemTokens) external returns (uint);
    function redeemUnderlying(uint redeemAmount) external returns (uint);
    function borrow(uint borrowAmount) external returns (uint);
    function repayBorrow(uint repayAmount) external returns (uint);
    function repayBorrowBehalf(address borrower, uint repayAmount) external returns (uint);
    function liquidateBorrow(address borrower, uint repayAmount, CTokenInterface cTokenCollateral) external returns (uint);
}

contract Comptroller {
	function enterMarkets(address[] memory cTokens) public returns (uint[] memory);
	function exitMarket(address cTokenAddress) external returns (uint);
	function getAccountLiquidity(address account) public view returns (uint, uint, uint);
	function borrowAllowed(address cToken, address borrower, uint borrowAmount) external returns (uint);
}

contract Compound is TradeUtils {

	Comptroller comptroller;
	CEther cEther;

	constructor(address payable _incognitoSmartContract, Comptroller _comptroller, CEther _cEther) public {
		incognitoSmartContract = _incognitoSmartContract;
		comptroller = _comptroller;
		cEther = _cEther;
	}
	
	struct BorrowInfo {
	    uint collateralAmount;
	    uint borrowedAmount;
	    address collateralToken;
	}
	
	mapping(address=>mapping(address=>BorrowInfo)) public borrowers;
	mapping(address=>mapping(address=>bool)) public hasBorrowed;
	
	mapping(address=>mapping(address=>uint)) public investors;

	// fallback function which allows transfer eth.
    function() external payable {}

    function enterMarket(address cToken) internal {
    	address[] memory tokens = new address[](1);
    	tokens[0] = cToken;
    	comptroller.enterMarkets(tokens);
    }

	/**
	 * @dev mint ethereum into compound smart contract
	 */
	function mint() internal {
		cEther.mint.value(msg.value)();
		enterMarket(address(cEther));
	}

	/**
	 * @dev mint a token into compound smart contract
	 * @param cToken: cToken address which represents to token in compound.
	 * @param mintAmount: minted token amount.
	 */
	function mint(CErc20 cToken, uint mintAmount) internal {
		require(balanceOf(ERC20(cToken.underlying())) >= mintAmount);
		approve(ERC20(cToken.underlying()), address(cToken), mintAmount);
		enterMarket(address(cToken));
		require(cToken.mint(mintAmount) == 0);
	}
	
	/**
	 * @dev this function is used to borrow token.
	 * @param borrower: address of borrower
	 * @param collateralToken: token address which is used to get collateral amount.
	 * @param srcAmount: amount is used to get in collateralToken.
	 * @param borrowedToken: token address that borrower wants to borrow.
	 * @param borrowedAmount: token's amount that borrower wants to borrow.
	 */
	function borrow(address borrower, address collateralToken, uint srcAmount, address borrowedToken, uint borrowedAmount) external payable isIncognitoSmartContract returns (address, uint) {
	    if (ERC20(collateralToken) == ETH_CONTRACT_ADDRESS) {
	        // mint eth 
	        mint();
	    } else {
    	    // check balance of collateralToken
    	    require(balanceOf(ERC20(collateralToken)) >= srcAmount);
    	    
    	    // mint token
    	    mint(CErc20(collateralToken), srcAmount);
	    }
	    // start borrow
	    borrowInternal(borrowedToken, borrowedAmount);
	    uint amount = balanceOf(ERC20(CErc20(borrowedToken).underlying()));
	    
	    // store borrower address
	    if (hasBorrowed[borrower][borrowedToken]) {
	        require(borrowers[borrower][borrowedToken].collateralToken == collateralToken);
	        borrowers[borrower][borrowedToken].borrowedAmount += amount;
	        borrowers[borrower][borrowedToken].collateralAmount += srcAmount;
	    } else {
	        borrowers[borrower][borrowedToken] = BorrowInfo(srcAmount, amount, collateralToken);
	        hasBorrowed[borrower][borrowedToken] = true;
	    }
	    
	    // transfer token to incognito smart contract
		ERC20(CErc20(borrowedToken).underlying()).transfer(incognitoSmartContract, amount);

		if (ERC20(borrowedToken) == ETH_CONTRACT_ADDRESS) {
			return (address(ETH_CONTRACT_ADDRESS), borrowedAmount);
		}
	    return (CErc20(borrowedToken).underlying(), borrowedAmount);
	}
	
	function borrowInternal(address cToken, uint borrowedAmount) internal {
	    if (ERC20(cToken) == ETH_CONTRACT_ADDRESS) {
	        require(cEther.borrow(borrowedAmount) == 0);
	    } else {
	        require(CErc20(cToken).borrow(borrowedAmount) == 0);
	    }
	}
	
	/**
	 * @dev user payback borrowed token
	 * @param borrower: address of borrower
	 * @param cToken: address of cToken that user borrows
	 */
	function payback(address borrower, address cToken) external payable isIncognitoSmartContract returns (address, uint) {
	    require(hasBorrowed[borrower][cToken]);

	    address collateralToken = borrowers[borrower][cToken].collateralToken;
	    uint repaidAmount = borrowers[borrower][cToken].borrowedAmount;
	    require(repaidAmount > 0);
	    
	    // aprove cToken to transfer token to contract.
	    if (ERC20(cToken) != ETH_CONTRACT_ADDRESS) {
	        require(repaidAmount <= balanceOf(ERC20(CErc20(cToken).underlying())));
	        address underlying = CErc20(cToken).underlying();
	        ERC20(underlying).approve(cToken, repaidAmount);
	        
	        // call repayBorrow
	        CErc20(cToken).repayBorrow(repaidAmount);
	    } else { 
	        // compare msg.value with repaidAmount
	        require(msg.value == repaidAmount);
	        require(repaidAmount <= address(this).balance);
	        
	        // call repayBorrow
	        cEther.repayBorrow.value(repaidAmount)();
	    }
	    
	    // redeem collateralToken
	    uint collateralBalance = balanceOfCToken(collateralToken);
	    redeemInternal(collateralToken, collateralBalance-100); // TODO: specify how to calculate the fee instead of hard code 100 wei.
	    
	    borrowers[borrower][cToken].borrowedAmount = 0;
	    borrowers[borrower][cToken].collateralAmount = 0;

	    if (ERC20(cToken) == ETH_CONTRACT_ADDRESS) {
	    	return (address(ETH_CONTRACT_ADDRESS), repaidAmount);
	    }
	    return (CErc20(cToken).underlying(), repaidAmount);
	}
	
	/**
	 * @dev user invests to a token
	 * @param investor: address of investor
	 * @param investedToken: token that is going to be invested
	 * @param investedAmount: amount of token that is invested
	 */
	function invest(address investor, address investedToken, uint investedAmount) external payable isIncognitoSmartContract returns (address, uint) {
	    if (ERC20(investedToken) == ETH_CONTRACT_ADDRESS) {
	        require(msg.value >= 0);
	        mint();
	        investors[investor][investedToken] += msg.value;
	        return (address(ETH_CONTRACT_ADDRESS), 0);
	    } 
        require(balanceOf(ERC20(CErc20(investedToken).underlying())) >= investedAmount);
        mint(CErc20(investedToken), investedAmount);
        investors[investor][investedToken] += investedAmount;
        return (CErc20(investedToken).underlying(), 0);
	}
	
	/**
	 * @dev user redeems invested token
	 * @param investor: address of investor
	 * @param investedToken: token that is going to be redeemed
	 * @param redeemedAmount: amount of token that is redeemed
	 */
	function redeem(address investor, address investedToken, uint redeemedAmount) external isIncognitoSmartContract returns (address, uint) {
	    require(investors[investor][investedToken] >= redeemedAmount);
	    investors[investor][investedToken] -= redeemedAmount;
	    return redeemInternal(investedToken, redeemedAmount-100);
	}

	function redeemInternal(address cToken, uint amount) internal returns (address, uint) {
	    if (ERC20(cToken) == ETH_CONTRACT_ADDRESS) {
	        cEther.redeem(amount);
	        incognitoSmartContract.transfer(amount);
	        return (address(ETH_CONTRACT_ADDRESS), amount);
	    }
        CErc20(cToken).redeem(amount);
        ERC20(CErc20(cToken).underlying()).transfer(incognitoSmartContract, amount);
        return (CErc20(cToken).underlying(), amount);
	}
	
	/**
	 * @dev return balance of current smart contract in cToken
	 * @param cToken: cToken address
	 * @return balance of smart contract in cToken
	 */
	function balanceOfCToken(address cToken) public view returns (uint) {
		if (ERC20(cToken) == ETH_CONTRACT_ADDRESS) {
			return cEther.balanceOf(address(this));
		} 
		return CErc20(cToken).balanceOf(address(this));
	}
}
