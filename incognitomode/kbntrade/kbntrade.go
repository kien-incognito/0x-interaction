// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbntrade

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

// KbntradeABI is the input ABI used to generate the binding from.
const KbntradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// KbntradeBin is the compiled bytecode used for deploying new contracts.
const KbntradeBin = `0x608060405234801561001057600080fd5b506040516112443803806112448339818101604052604081101561003357600080fd5b81019080805190602001909291908051906020019092919050505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050611164806100e06000396000f3fe60806040526004361061004a5760003560e01c80630aea81881461004c5780630e32db52146100e257806372e94bf614610197578063785250da146101ee578063b42a644b14610245575b005b34801561005857600080fd5b506100c56004803603606081101561006f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061029c565b604051808381526020018281526020019250505060405180910390f35b61014e600480360360608110156100f857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103ca565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156101a357600080fd5b506101ac610581565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101fa57600080fd5b50610203610586565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025157600080fd5b5061025a6105ac565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558685876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b15801561037957600080fd5b505afa15801561038d573d6000803e3d6000fd5b505050506040513d60408110156103a357600080fd5b81019080805190602001909291908051906020019092919050505091509150935093915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461042657600080fd5b83610430866105d1565b101561043b57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141561047457600080fd5b6000809050600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610546576104da86600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16876106e6565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161461052a57600061051b8787876108b7565b1161052557600080fd5b610541565b60006105368787610b06565b1161054057600080fd5b5b61055d565b60006105528587610d34565b1161055c57600080fd5b5b610566846105d1565b90506105728482610f80565b83819250925050935093915050565b600081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610626573073ffffffffffffffffffffffffffffffffffffffff163190506106e1565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156106a357600080fd5b505afa1580156106b7573d6000803e3d6000fd5b505050506040513d60208110156106cd57600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146108b2578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156107a257600080fd5b505af11580156107b6573d6000803e3d6000fd5b505050506040513d60208110156107cc57600080fd5b81019080805190602001909291905050506107e657600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561086d57600080fd5b505af1158015610881573d6000803e3d6000fd5b505050506040513d602081101561089757600080fd5b81019080805190602001909291905050506108b157600080fd5b5b505050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558685876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b15801561099457600080fd5b505afa1580156109a8573d6000803e3d6000fd5b505050506040513d60408110156109be57600080fd5b810190808051906020019092919080519060200190929190505050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637409e2eb868686856040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050602060405180830381600087803b158015610ac157600080fd5b505af1158015610ad5573d6000803e3d6000fd5b505050506040513d6020811015610aeb57600080fd5b81019080805190602001909291905050509150509392505050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e558573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b158015610bf757600080fd5b505afa158015610c0b573d6000803e3d6000fd5b505050506040513d6040811015610c2157600080fd5b810190808051906020019092919080519060200190929190505050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633bba21dc8585846040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015610cf057600080fd5b505af1158015610d04573d6000803e3d6000fd5b505050506040513d6020811015610d1a57600080fd5b810190808051906020019092919050505091505092915050565b6000813073ffffffffffffffffffffffffffffffffffffffff16311015610d5a57600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663809a9e5573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee86346040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604080518083038186803b158015610e4a57600080fd5b505afa158015610e5e573d6000803e3d6000fd5b505050506040513d6040811015610e7457600080fd5b810190808051906020019092919080519060200190929190505050509050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637a2a04568486846040518463ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506020604051808303818588803b158015610f3b57600080fd5b505af1158015610f4f573d6000803e3d6000fd5b50505050506040513d6020811015610f6657600080fd5b810190808051906020019092919050505091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561104657803073ffffffffffffffffffffffffffffffffffffffff16311015610fd957600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611040573d6000803e3d6000fd5b5061112b565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156110ee57600080fd5b505af1158015611102573d6000803e3d6000fd5b505050506040513d602081101561111857600080fd5b8101908080519060200190929190505050505b505056fea265627a7a72315820a34f7d2c03ffc48c17225eaa863b587fa41353ba2b99045a29df9fa1755f276964736f6c634300050c0032`

// DeployKbntrade deploys a new Ethereum contract, binding an instance of Kbntrade to it.
func DeployKbntrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Kbntrade, error) {
	parsed, err := abi.JSON(strings.NewReader(KbntradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbntradeBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kbntrade{KbntradeCaller: KbntradeCaller{contract: contract}, KbntradeTransactor: KbntradeTransactor{contract: contract}, KbntradeFilterer: KbntradeFilterer{contract: contract}}, nil
}

// Kbntrade is an auto generated Go binding around an Ethereum contract.
type Kbntrade struct {
	KbntradeCaller     // Read-only binding to the contract
	KbntradeTransactor // Write-only binding to the contract
	KbntradeFilterer   // Log filterer for contract events
}

// KbntradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbntradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbntradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbntradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbntradeSession struct {
	Contract     *Kbntrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbntradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbntradeCallerSession struct {
	Contract *KbntradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KbntradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbntradeTransactorSession struct {
	Contract     *KbntradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KbntradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbntradeRaw struct {
	Contract *Kbntrade // Generic contract binding to access the raw methods on
}

// KbntradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbntradeCallerRaw struct {
	Contract *KbntradeCaller // Generic read-only contract binding to access the raw methods on
}

// KbntradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbntradeTransactorRaw struct {
	Contract *KbntradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbntrade creates a new instance of Kbntrade, bound to a specific deployed contract.
func NewKbntrade(address common.Address, backend bind.ContractBackend) (*Kbntrade, error) {
	contract, err := bindKbntrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kbntrade{KbntradeCaller: KbntradeCaller{contract: contract}, KbntradeTransactor: KbntradeTransactor{contract: contract}, KbntradeFilterer: KbntradeFilterer{contract: contract}}, nil
}

// NewKbntradeCaller creates a new read-only instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeCaller(address common.Address, caller bind.ContractCaller) (*KbntradeCaller, error) {
	contract, err := bindKbntrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbntradeCaller{contract: contract}, nil
}

// NewKbntradeTransactor creates a new write-only instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeTransactor(address common.Address, transactor bind.ContractTransactor) (*KbntradeTransactor, error) {
	contract, err := bindKbntrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbntradeTransactor{contract: contract}, nil
}

// NewKbntradeFilterer creates a new log filterer instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeFilterer(address common.Address, filterer bind.ContractFilterer) (*KbntradeFilterer, error) {
	contract, err := bindKbntrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbntradeFilterer{contract: contract}, nil
}

// bindKbntrade binds a generic wrapper to an already deployed contract.
func bindKbntrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbntradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbntrade *KbntradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kbntrade.Contract.KbntradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbntrade *KbntradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbntrade.Contract.KbntradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbntrade *KbntradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbntrade.Contract.KbntradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbntrade *KbntradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kbntrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbntrade *KbntradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbntrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbntrade *KbntradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbntrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Kbntrade.Contract.ETHCONTRACTADDRESS(&_Kbntrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Kbntrade.Contract.ETHCONTRACTADDRESS(&_Kbntrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Kbntrade.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _Kbntrade.Contract.GetConversionRates(&_Kbntrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _Kbntrade.Contract.GetConversionRates(&_Kbntrade.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeSession) IncognitoSmartContract() (common.Address, error) {
	return _Kbntrade.Contract.IncognitoSmartContract(&_Kbntrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Kbntrade.Contract.IncognitoSmartContract(&_Kbntrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _Kbntrade.Contract.KyberNetworkProxyContract(&_Kbntrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _Kbntrade.Contract.KyberNetworkProxyContract(&_Kbntrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x0e32db52.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken) returns(address, uint256)
func (_Kbntrade *KbntradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*types.Transaction, error) {
	return _Kbntrade.contract.Transact(opts, "trade", srcToken, srcQty, destToken)
}

// Trade is a paid mutator transaction binding the contract method 0x0e32db52.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken) returns(address, uint256)
func (_Kbntrade *KbntradeSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*types.Transaction, error) {
	return _Kbntrade.Contract.Trade(&_Kbntrade.TransactOpts, srcToken, srcQty, destToken)
}

// Trade is a paid mutator transaction binding the contract method 0x0e32db52.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken) returns(address, uint256)
func (_Kbntrade *KbntradeTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*types.Transaction, error) {
	return _Kbntrade.Contract.Trade(&_Kbntrade.TransactOpts, srcToken, srcQty, destToken)
}
