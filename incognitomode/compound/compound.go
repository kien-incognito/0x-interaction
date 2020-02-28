// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compound

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CompoundABI is the input ABI used to generate the binding from.
const CompoundABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"},{\"internalType\":\"contractComptroller\",\"name\":\"_comptroller\",\"type\":\"address\"},{\"internalType\":\"contractCEther\",\"name\":\"_cEther\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasBorrowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"investors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrowedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"payback\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"investedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"investedAmount\",\"type\":\"uint256\"}],\"name\":\"invest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"investor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"investedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemedAmount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"balanceOfCToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CompoundBin is the compiled bytecode used for deploying new contracts.
var CompoundBin = "0x608060405234801561001057600080fd5b50604051612e1d380380612e1d8339818101604052606081101561003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050612cf18061012c6000396000f3fe6080604052600436106100915760003560e01c806372e94bf61161005957806372e94bf6146103f95780637aeec947146104505780638283e3e2146104b5578063b42a644b1461053e578063ca89346a1461059557610091565b80630e6dfcd5146100935780633fa744ad14610155578063574b895c146102005780635e9bef43146102bf578063653ac42214610344575b005b34801561009f57600080fd5b5061010c600480360360608110156100b657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610674565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b6101b76004803603604081101561016b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107fc565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561020c57600080fd5b5061026f6004803603604081101561022357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f5e565b604051808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390f35b3480156102cb57600080fd5b5061032e600480360360408110156102e257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fb5565b6040518082815260200191505060405180910390f35b6103b06004803603606081101561035a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610fda565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561040557600080fd5b5061040e611474565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561045c57600080fd5b5061049f6004803603602081101561047357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611479565b6040518082815260200191505060405180910390f35b3480156104c157600080fd5b50610524600480360360408110156104d857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611650565b604051808215151515815260200191505060405180910390f35b34801561054a57600080fd5b5061055361167f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61062b600480360360a08110156105ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116a4565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106d057600080fd5b82600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561075957600080fd5b82600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506107f08460648503611e8d565b91509150935093915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461085857600080fd5b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166108eb57600080fd5b6000600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905060008111610a2057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610cc657610add8573ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610a9d57600080fd5b505afa158015610ab1573d6000803e3d6000fd5b505050506040513d6020811015610ac757600080fd5b810190808051906020019092919050505061226b565b811115610ae957600080fd5b60008573ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610b3157600080fd5b505afa158015610b45573d6000803e3d6000fd5b505050506040513d6020811015610b5b57600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b387846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610bf557600080fd5b505af1158015610c09573d6000803e3d6000fd5b505050506040513d6020811015610c1f57600080fd5b8101908080519060200190929190505050508573ffffffffffffffffffffffffffffffffffffffff16630e752702836040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015610c8457600080fd5b505af1158015610c98573d6000803e3d6000fd5b505050506040513d6020811015610cae57600080fd5b81019080805190602001909291905050505050610d63565b803414610cd257600080fd5b47811115610cdf57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634e4d9fea826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610d4957600080fd5b505af1158015610d5d573d6000803e3d6000fd5b50505050505b6000610d6e83611479565b9050610d7d8360648303611e8d565b50506000600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415610ecd5760008294509450505050610f57565b8573ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610f1357600080fd5b505afa158015610f27573d6000803e3d6000fd5b505050506040513d6020811015610f3d57600080fd5b810190808051906020019092919050505082945094505050505b9250929050565b6003602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b6005602052816000526040600020602052806000526040600020600091509150505481565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461103657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561124b57600034101561107957600080fd5b6000611083612369565b905080600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156111d957600080fd5b505af11580156111ed573d6000803e3d6000fd5b505050506040513d602081101561120357600080fd5b810190808051906020019092919050505061121d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681925092505061146c565b826112d58573ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561129557600080fd5b505afa1580156112a9573d6000803e3d6000fd5b505050506040513d60208110156112bf57600080fd5b810190808051906020019092919050505061226b565b10156112e057600080fd5b60006112ec85856124cd565b905080600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561142057600080fd5b505af1158015611434573d6000803e3d6000fd5b505050506040513d602081101561144a57600080fd5b810190808051906020019092919050505061146457600080fd5b848192509250505b935093915050565b600081565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561159057600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561154e57600080fd5b505afa158015611562573d6000803e3d6000fd5b505050506040513d602081101561157857600080fd5b8101908080519060200190929190505050905061164b565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561160d57600080fd5b505afa158015611621573d6000803e3d6000fd5b505050506040513d602081101561163757600080fd5b810190808051906020019092919050505090505b919050565b60046020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461170057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561176e5761173d612369565b50611769600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612749565b611798565b846117788761226b565b101561178357600080fd5b61178d86866124cd565b5061179786612749565b5b6117a28484612956565b600061182d8573ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156117ed57600080fd5b505afa158015611801573d6000803e3d6000fd5b505050506040513d602081101561181757600080fd5b810190808051906020019092919050505061226b565b9050600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611ab4578673ffffffffffffffffffffffffffffffffffffffff16600360008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461199557600080fd5b80600360008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000828254019250508190555085600360008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282540192505081905550611c55565b60405180606001604052808781526020018281526020018873ffffffffffffffffffffffffffffffffffffffff16815250600360008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050506001600460008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8473ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611c9b57600080fd5b505afa158015611caf573d6000803e3d6000fd5b505050506040513d6020811015611cc557600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611d7d57600080fd5b505af1158015611d91573d6000803e3d6000fd5b505050506040513d6020811015611da757600080fd5b810190808051906020019092919050505050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415611dfb576000849250925050611e83565b8473ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611e4157600080fd5b505afa158015611e55573d6000803e3d6000fd5b505050506040513d6020811015611e6b57600080fd5b81019080805190602001909291905050508492509250505b9550959350505050565b600080600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611fea57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015611f3a57600080fd5b505af1158015611f4e573d6000803e3d6000fd5b505050506040513d6020811015611f6457600080fd5b8101908080519060200190929190505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015611fdd573d6000803e3d6000fd5b5060008391509150612264565b8373ffffffffffffffffffffffffffffffffffffffff1663db006a75846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561203d57600080fd5b505af1158015612051573d6000803e3d6000fd5b505050506040513d602081101561206757600080fd5b8101908080519060200190929190505050508373ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156120bf57600080fd5b505afa1580156120d3573d6000803e3d6000fd5b505050506040513d60208110156120e957600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156121a157600080fd5b505af11580156121b5573d6000803e3d6000fd5b505050506040513d60208110156121cb57600080fd5b8101908080519060200190929190505050508373ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561222357600080fd5b505afa158015612237573d6000803e3d6000fd5b505050506040513d602081101561224d57600080fd5b810190808051906020019092919050505083915091505b9250929050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156122a957479050612364565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561232657600080fd5b505afa15801561233a573d6000803e3d6000fd5b505050506040513d602081101561235057600080fd5b810190808051906020019092919050505090505b919050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631249c58b346040518263ffffffff1660e01b81526004016000604051808303818588803b1580156123d557600080fd5b505af11580156123e9573d6000803e3d6000fd5b5050505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561248d57600080fd5b505afa1580156124a1573d6000803e3d6000fd5b505050506040513d60208110156124b757600080fd5b8101908080519060200190929190505050905090565b6000816125598473ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561251957600080fd5b505afa15801561252d573d6000803e3d6000fd5b505050506040513d602081101561254357600080fd5b810190808051906020019092919050505061226b565b101561256457600080fd5b6125ef8373ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156125ad57600080fd5b505afa1580156125c1573d6000803e3d6000fd5b505050506040513d60208110156125d757600080fd5b81019080805190602001909291905050508484612aeb565b60008373ffffffffffffffffffffffffffffffffffffffff1663a0712d68846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561264457600080fd5b505af1158015612658573d6000803e3d6000fd5b505050506040513d602081101561266e57600080fd5b81019080805190602001909291905050501461268957600080fd5b8273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561270657600080fd5b505afa15801561271a573d6000803e3d6000fd5b505050506040513d602081101561273057600080fd5b8101908080519060200190929190505050905092915050565b6060600160405190808252806020026020018201604052801561277b5781602001602082028038833980820191505090505b509050818160008151811061278c57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2998238826040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019060200280838360005b8381101561285757808201518184015260208101905061283c565b5050505090500192505050600060405180830381600087803b15801561287c57600080fd5b505af1158015612890573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156128ba57600080fd5b81019080805160405193929190846401000000008211156128da57600080fd5b838201915060208201858111156128f057600080fd5b825186602082028301116401000000008211171561290d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015612944578082015181840152602081019050612929565b50505050905001604052505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612a4c576000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5ebeaec836040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015612a0257600080fd5b505af1158015612a16573d6000803e3d6000fd5b505050506040513d6020811015612a2c57600080fd5b810190808051906020019092919050505014612a4757600080fd5b612ae7565b60008273ffffffffffffffffffffffffffffffffffffffff1663c5ebeaec836040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015612aa157600080fd5b505af1158015612ab5573d6000803e3d6000fd5b505050506040513d6020811015612acb57600080fd5b810190808051906020019092919050505014612ae657600080fd5b5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614612cb7578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015612ba757600080fd5b505af1158015612bbb573d6000803e3d6000fd5b505050506040513d6020811015612bd157600080fd5b8101908080519060200190929190505050612beb57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015612c7257600080fd5b505af1158015612c86573d6000803e3d6000fd5b505050506040513d6020811015612c9c57600080fd5b8101908080519060200190929190505050612cb657600080fd5b5b50505056fea265627a7a723158209becf7afab15761b0642fed0a0bf8e56ec5199724e852e53e90d56ec1b2a95a064736f6c63430005100032"

// DeployCompound deploys a new Ethereum contract, binding an instance of Compound to it.
func DeployCompound(auth *bind.TransactOpts, backend bind.ContractBackend, _incognitoSmartContract common.Address, _comptroller common.Address, _cEther common.Address) (common.Address, *types.Transaction, *Compound, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompoundBin), backend, _incognitoSmartContract, _comptroller, _cEther)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// Compound is an auto generated Go binding around an Ethereum contract.
type Compound struct {
	CompoundCaller     // Read-only binding to the contract
	CompoundTransactor // Write-only binding to the contract
	CompoundFilterer   // Log filterer for contract events
}

// CompoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundSession struct {
	Contract     *Compound         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundCallerSession struct {
	Contract *CompoundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundTransactorSession struct {
	Contract     *CompoundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundRaw struct {
	Contract *Compound // Generic contract binding to access the raw methods on
}

// CompoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundCallerRaw struct {
	Contract *CompoundCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundTransactorRaw struct {
	Contract *CompoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompound creates a new instance of Compound, bound to a specific deployed contract.
func NewCompound(address common.Address, backend bind.ContractBackend) (*Compound, error) {
	contract, err := bindCompound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// NewCompoundCaller creates a new read-only instance of Compound, bound to a specific deployed contract.
func NewCompoundCaller(address common.Address, caller bind.ContractCaller) (*CompoundCaller, error) {
	contract, err := bindCompound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundCaller{contract: contract}, nil
}

// NewCompoundTransactor creates a new write-only instance of Compound, bound to a specific deployed contract.
func NewCompoundTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundTransactor, error) {
	contract, err := bindCompound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundTransactor{contract: contract}, nil
}

// NewCompoundFilterer creates a new log filterer instance of Compound, bound to a specific deployed contract.
func NewCompoundFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundFilterer, error) {
	contract, err := bindCompound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundFilterer{contract: contract}, nil
}

// bindCompound binds a generic wrapper to an already deployed contract.
func bindCompound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.CompoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Compound *CompoundCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Compound *CompoundSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Compound.Contract.ETHCONTRACTADDRESS(&_Compound.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Compound *CompoundCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Compound.Contract.ETHCONTRACTADDRESS(&_Compound.CallOpts)
}

// BalanceOfCToken is a free data retrieval call binding the contract method 0x7aeec947.
//
// Solidity: function balanceOfCToken(address cToken) constant returns(uint256)
func (_Compound *CompoundCaller) BalanceOfCToken(opts *bind.CallOpts, cToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "balanceOfCToken", cToken)
	return *ret0, err
}

// BalanceOfCToken is a free data retrieval call binding the contract method 0x7aeec947.
//
// Solidity: function balanceOfCToken(address cToken) constant returns(uint256)
func (_Compound *CompoundSession) BalanceOfCToken(cToken common.Address) (*big.Int, error) {
	return _Compound.Contract.BalanceOfCToken(&_Compound.CallOpts, cToken)
}

// BalanceOfCToken is a free data retrieval call binding the contract method 0x7aeec947.
//
// Solidity: function balanceOfCToken(address cToken) constant returns(uint256)
func (_Compound *CompoundCallerSession) BalanceOfCToken(cToken common.Address) (*big.Int, error) {
	return _Compound.Contract.BalanceOfCToken(&_Compound.CallOpts, cToken)
}

// Borrowers is a free data retrieval call binding the contract method 0x574b895c.
//
// Solidity: function borrowers(address , address ) constant returns(uint256 collateralAmount, uint256 borrowedAmount, address collateralToken)
func (_Compound *CompoundCaller) Borrowers(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	CollateralAmount *big.Int
	BorrowedAmount   *big.Int
	CollateralToken  common.Address
}, error) {
	ret := new(struct {
		CollateralAmount *big.Int
		BorrowedAmount   *big.Int
		CollateralToken  common.Address
	})
	out := ret
	err := _Compound.contract.Call(opts, out, "borrowers", arg0, arg1)
	return *ret, err
}

// Borrowers is a free data retrieval call binding the contract method 0x574b895c.
//
// Solidity: function borrowers(address , address ) constant returns(uint256 collateralAmount, uint256 borrowedAmount, address collateralToken)
func (_Compound *CompoundSession) Borrowers(arg0 common.Address, arg1 common.Address) (struct {
	CollateralAmount *big.Int
	BorrowedAmount   *big.Int
	CollateralToken  common.Address
}, error) {
	return _Compound.Contract.Borrowers(&_Compound.CallOpts, arg0, arg1)
}

// Borrowers is a free data retrieval call binding the contract method 0x574b895c.
//
// Solidity: function borrowers(address , address ) constant returns(uint256 collateralAmount, uint256 borrowedAmount, address collateralToken)
func (_Compound *CompoundCallerSession) Borrowers(arg0 common.Address, arg1 common.Address) (struct {
	CollateralAmount *big.Int
	BorrowedAmount   *big.Int
	CollateralToken  common.Address
}, error) {
	return _Compound.Contract.Borrowers(&_Compound.CallOpts, arg0, arg1)
}

// HasBorrowed is a free data retrieval call binding the contract method 0x8283e3e2.
//
// Solidity: function hasBorrowed(address , address ) constant returns(bool)
func (_Compound *CompoundCaller) HasBorrowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "hasBorrowed", arg0, arg1)
	return *ret0, err
}

// HasBorrowed is a free data retrieval call binding the contract method 0x8283e3e2.
//
// Solidity: function hasBorrowed(address , address ) constant returns(bool)
func (_Compound *CompoundSession) HasBorrowed(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Compound.Contract.HasBorrowed(&_Compound.CallOpts, arg0, arg1)
}

// HasBorrowed is a free data retrieval call binding the contract method 0x8283e3e2.
//
// Solidity: function hasBorrowed(address , address ) constant returns(bool)
func (_Compound *CompoundCallerSession) HasBorrowed(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Compound.Contract.HasBorrowed(&_Compound.CallOpts, arg0, arg1)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Compound *CompoundCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Compound *CompoundSession) IncognitoSmartContract() (common.Address, error) {
	return _Compound.Contract.IncognitoSmartContract(&_Compound.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Compound *CompoundCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Compound.Contract.IncognitoSmartContract(&_Compound.CallOpts)
}

// Investors is a free data retrieval call binding the contract method 0x5e9bef43.
//
// Solidity: function investors(address , address ) constant returns(uint256)
func (_Compound *CompoundCaller) Investors(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Compound.contract.Call(opts, out, "investors", arg0, arg1)
	return *ret0, err
}

// Investors is a free data retrieval call binding the contract method 0x5e9bef43.
//
// Solidity: function investors(address , address ) constant returns(uint256)
func (_Compound *CompoundSession) Investors(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Compound.Contract.Investors(&_Compound.CallOpts, arg0, arg1)
}

// Investors is a free data retrieval call binding the contract method 0x5e9bef43.
//
// Solidity: function investors(address , address ) constant returns(uint256)
func (_Compound *CompoundCallerSession) Investors(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Compound.Contract.Investors(&_Compound.CallOpts, arg0, arg1)
}

// Borrow is a paid mutator transaction binding the contract method 0xca89346a.
//
// Solidity: function borrow(address borrower, address collateralToken, uint256 srcAmount, address borrowedToken, uint256 borrowedAmount) returns(address, uint256)
func (_Compound *CompoundTransactor) Borrow(opts *bind.TransactOpts, borrower common.Address, collateralToken common.Address, srcAmount *big.Int, borrowedToken common.Address, borrowedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "borrow", borrower, collateralToken, srcAmount, borrowedToken, borrowedAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xca89346a.
//
// Solidity: function borrow(address borrower, address collateralToken, uint256 srcAmount, address borrowedToken, uint256 borrowedAmount) returns(address, uint256)
func (_Compound *CompoundSession) Borrow(borrower common.Address, collateralToken common.Address, srcAmount *big.Int, borrowedToken common.Address, borrowedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Borrow(&_Compound.TransactOpts, borrower, collateralToken, srcAmount, borrowedToken, borrowedAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xca89346a.
//
// Solidity: function borrow(address borrower, address collateralToken, uint256 srcAmount, address borrowedToken, uint256 borrowedAmount) returns(address, uint256)
func (_Compound *CompoundTransactorSession) Borrow(borrower common.Address, collateralToken common.Address, srcAmount *big.Int, borrowedToken common.Address, borrowedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Borrow(&_Compound.TransactOpts, borrower, collateralToken, srcAmount, borrowedToken, borrowedAmount)
}

// Invest is a paid mutator transaction binding the contract method 0x653ac422.
//
// Solidity: function invest(address investor, address investedToken, uint256 investedAmount) returns(address, uint256)
func (_Compound *CompoundTransactor) Invest(opts *bind.TransactOpts, investor common.Address, investedToken common.Address, investedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "invest", investor, investedToken, investedAmount)
}

// Invest is a paid mutator transaction binding the contract method 0x653ac422.
//
// Solidity: function invest(address investor, address investedToken, uint256 investedAmount) returns(address, uint256)
func (_Compound *CompoundSession) Invest(investor common.Address, investedToken common.Address, investedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Invest(&_Compound.TransactOpts, investor, investedToken, investedAmount)
}

// Invest is a paid mutator transaction binding the contract method 0x653ac422.
//
// Solidity: function invest(address investor, address investedToken, uint256 investedAmount) returns(address, uint256)
func (_Compound *CompoundTransactorSession) Invest(investor common.Address, investedToken common.Address, investedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Invest(&_Compound.TransactOpts, investor, investedToken, investedAmount)
}

// Payback is a paid mutator transaction binding the contract method 0x3fa744ad.
//
// Solidity: function payback(address borrower, address cToken) returns(address, uint256)
func (_Compound *CompoundTransactor) Payback(opts *bind.TransactOpts, borrower common.Address, cToken common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "payback", borrower, cToken)
}

// Payback is a paid mutator transaction binding the contract method 0x3fa744ad.
//
// Solidity: function payback(address borrower, address cToken) returns(address, uint256)
func (_Compound *CompoundSession) Payback(borrower common.Address, cToken common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Payback(&_Compound.TransactOpts, borrower, cToken)
}

// Payback is a paid mutator transaction binding the contract method 0x3fa744ad.
//
// Solidity: function payback(address borrower, address cToken) returns(address, uint256)
func (_Compound *CompoundTransactorSession) Payback(borrower common.Address, cToken common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Payback(&_Compound.TransactOpts, borrower, cToken)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address investor, address investedToken, uint256 redeemedAmount) returns(address, uint256)
func (_Compound *CompoundTransactor) Redeem(opts *bind.TransactOpts, investor common.Address, investedToken common.Address, redeemedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "redeem", investor, investedToken, redeemedAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address investor, address investedToken, uint256 redeemedAmount) returns(address, uint256)
func (_Compound *CompoundSession) Redeem(investor common.Address, investedToken common.Address, redeemedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Redeem(&_Compound.TransactOpts, investor, investedToken, redeemedAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address investor, address investedToken, uint256 redeemedAmount) returns(address, uint256)
func (_Compound *CompoundTransactorSession) Redeem(investor common.Address, investedToken common.Address, redeemedAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Redeem(&_Compound.TransactOpts, investor, investedToken, redeemedAmount)
}
