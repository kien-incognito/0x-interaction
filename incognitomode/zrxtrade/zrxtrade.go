// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrxtrade

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

// ZrxtradeABI is the input ABI used to generate the binding from.
const ZrxtradeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zeroProxy\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callDataHex\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWrapETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZrxtradeBin is the compiled bytecode used for deploying new contracts.
const ZrxtradeBin = `0x60806040523480156200001157600080fd5b5060405162000eff38038062000eff833981810160405262000037919081019062000130565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505062000202565b6000815190506200011381620001ce565b92915050565b6000815190506200012a81620001e8565b92915050565b6000806000606084860312156200014657600080fd5b6000620001568682870162000102565b9350506020620001698682870162000102565b92505060406200017c8682870162000119565b9150509250925092565b60006200019382620001ae565b9050919050565b6000620001a782620001ae565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620001d98162000186565b8114620001e557600080fd5b50565b620001f3816200019a565b8114620001ff57600080fd5b50565b610ced80620002126000396000f3fe60806040526004361061003f5760003560e01c80630e0569a61461004157806372e94bf61461006a57806398c3615214610095578063b42a644b146100c6575b005b34801561004d57600080fd5b50610068600480360361006391908101906108ea565b6100f1565b005b34801561007657600080fd5b5061007f610181565b60405161008c9190610a8f565b60405180910390f35b6100af60048036036100aa919081019061085b565b610186565b6040516100bd929190610a66565b60405180910390f35b3480156100d257600080fd5b506100db610310565b6040516100e891906109f9565b60405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b815260040161014c9190610aaa565b600060405180830381600087803b15801561016657600080fd5b505af115801561017a573d6000803e3d6000fd5b5050505050565b600081565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e257600080fd5b61020f87600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1688610335565b600083905060008173ffffffffffffffffffffffffffffffffffffffff16348760405161023c91906109c7565b60006040518083038185875af1925050503d8060008114610279576040519150601f19603f3d011682016040523d82523d6000602084013e61027e565b606091505b505090508061028c57600080fd5b60008073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614156102e7576102cc600061049c565b90506102d7816100f1565b6102e2600082610515565b6102fd565b6102f08861049c565b90506102fc8882610515565b5b8781945094505050509550959350505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610497578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b81526004016103a5929190610a3d565b602060405180830381600087803b1580156103bf57600080fd5b505af11580156103d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103f79190810190610832565b61040057600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b815260040161043b929190610a66565b602060405180830381600087803b15801561045557600080fd5b505af1158015610469573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061048d9190810190610832565b61049657600080fd5b5b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610504576104fd600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661068f565b9050610510565b61050d8261068f565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156105db57803073ffffffffffffffffffffffffffffffffffffffff1631101561056e57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156105d5573d6000803e3d6000fd5b5061068b565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610637929190610a14565b602060405180830381600087803b15801561065157600080fd5b505af1158015610665573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506106899190810190610832565b505b5050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156106e4573073ffffffffffffffffffffffffffffffffffffffff16319050610770565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161071d91906109de565b60206040518083038186803b15801561073557600080fd5b505afa158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061076d9190810190610913565b90505b919050565b60008135905061078481610c4e565b92915050565b60008151905061079981610c65565b92915050565b600082601f8301126107b057600080fd5b81356107c36107be82610af2565b610ac5565b915080825260208301602083018583830111156107df57600080fd5b6107ea838284610c0c565b50505092915050565b60008135905061080281610c7c565b92915050565b60008135905061081781610c93565b92915050565b60008151905061082c81610c93565b92915050565b60006020828403121561084457600080fd5b60006108528482850161078a565b91505092915050565b600080600080600060a0868803121561087357600080fd5b6000610881888289016107f3565b955050602061089288828901610808565b94505060406108a3888289016107f3565b935050606086013567ffffffffffffffff8111156108c057600080fd5b6108cc8882890161079f565b92505060806108dd88828901610775565b9150509295509295909350565b6000602082840312156108fc57600080fd5b600061090a84828501610808565b91505092915050565b60006020828403121561092557600080fd5b60006109338482850161081d565b91505092915050565b61094581610ba0565b82525050565b61095481610b46565b82525050565b61096381610b34565b82525050565b600061097482610b1e565b61097e8185610b29565b935061098e818560208601610c1b565b80840191505092915050565b6109a381610bb2565b82525050565b6109b281610bd6565b82525050565b6109c181610b96565b82525050565b60006109d38284610969565b915081905092915050565b60006020820190506109f3600083018461093c565b92915050565b6000602082019050610a0e600083018461094b565b92915050565b6000604082019050610a29600083018561093c565b610a3660208301846109b8565b9392505050565b6000604082019050610a52600083018561095a565b610a5f60208301846109a9565b9392505050565b6000604082019050610a7b600083018561095a565b610a8860208301846109b8565b9392505050565b6000602082019050610aa4600083018461099a565b92915050565b6000602082019050610abf60008301846109b8565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610ae857600080fd5b8060405250919050565b600067ffffffffffffffff821115610b0957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081905092915050565b6000610b3f82610b76565b9050919050565b6000610b5182610b76565b9050919050565b60008115159050919050565b6000610b6f82610b34565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610bab82610be8565b9050919050565b6000610bbd82610bc4565b9050919050565b6000610bcf82610b76565b9050919050565b6000610be182610b96565b9050919050565b6000610bf382610bfa565b9050919050565b6000610c0582610b76565b9050919050565b82818337600083830152505050565b60005b83811015610c39578082015181840152602081019050610c1e565b83811115610c48576000848401525b50505050565b610c5781610b34565b8114610c6257600080fd5b50565b610c6e81610b58565b8114610c7957600080fd5b50565b610c8581610b64565b8114610c9057600080fd5b50565b610c9c81610b96565b8114610ca757600080fd5b5056fea365627a7a72315820156455037c199464176cd3d0ba1b4fbd19d46aab330ade8fd7c35b4309e3c7116c6578706572696d656e74616cf564736f6c634300050c0040`

// DeployZrxtrade deploys a new Ethereum contract, binding an instance of Zrxtrade to it.
func DeployZrxtrade(auth *bind.TransactOpts, backend bind.ContractBackend, _wETH common.Address, _zeroProxy common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Zrxtrade, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrxtradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZrxtradeBin), backend, _wETH, _zeroProxy, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zrxtrade{ZrxtradeCaller: ZrxtradeCaller{contract: contract}, ZrxtradeTransactor: ZrxtradeTransactor{contract: contract}, ZrxtradeFilterer: ZrxtradeFilterer{contract: contract}}, nil
}

// Zrxtrade is an auto generated Go binding around an Ethereum contract.
type Zrxtrade struct {
	ZrxtradeCaller     // Read-only binding to the contract
	ZrxtradeTransactor // Write-only binding to the contract
	ZrxtradeFilterer   // Log filterer for contract events
}

// ZrxtradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrxtradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrxtradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZrxtradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrxtradeSession struct {
	Contract     *Zrxtrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrxtradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrxtradeCallerSession struct {
	Contract *ZrxtradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZrxtradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrxtradeTransactorSession struct {
	Contract     *ZrxtradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZrxtradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrxtradeRaw struct {
	Contract *Zrxtrade // Generic contract binding to access the raw methods on
}

// ZrxtradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrxtradeCallerRaw struct {
	Contract *ZrxtradeCaller // Generic read-only contract binding to access the raw methods on
}

// ZrxtradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrxtradeTransactorRaw struct {
	Contract *ZrxtradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrxtrade creates a new instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtrade(address common.Address, backend bind.ContractBackend) (*Zrxtrade, error) {
	contract, err := bindZrxtrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zrxtrade{ZrxtradeCaller: ZrxtradeCaller{contract: contract}, ZrxtradeTransactor: ZrxtradeTransactor{contract: contract}, ZrxtradeFilterer: ZrxtradeFilterer{contract: contract}}, nil
}

// NewZrxtradeCaller creates a new read-only instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeCaller(address common.Address, caller bind.ContractCaller) (*ZrxtradeCaller, error) {
	contract, err := bindZrxtrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeCaller{contract: contract}, nil
}

// NewZrxtradeTransactor creates a new write-only instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrxtradeTransactor, error) {
	contract, err := bindZrxtrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeTransactor{contract: contract}, nil
}

// NewZrxtradeFilterer creates a new log filterer instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZrxtradeFilterer, error) {
	contract, err := bindZrxtrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeFilterer{contract: contract}, nil
}

// bindZrxtrade binds a generic wrapper to an already deployed contract.
func bindZrxtrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrxtradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zrxtrade *ZrxtradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zrxtrade.Contract.ZrxtradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zrxtrade *ZrxtradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zrxtrade.Contract.ZrxtradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zrxtrade *ZrxtradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zrxtrade.Contract.ZrxtradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zrxtrade *ZrxtradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zrxtrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zrxtrade *ZrxtradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zrxtrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zrxtrade *ZrxtradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zrxtrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zrxtrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Zrxtrade.Contract.ETHCONTRACTADDRESS(&_Zrxtrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Zrxtrade.Contract.ETHCONTRACTADDRESS(&_Zrxtrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zrxtrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeSession) IncognitoSmartContract() (common.Address, error) {
	return _Zrxtrade.Contract.IncognitoSmartContract(&_Zrxtrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Zrxtrade.Contract.IncognitoSmartContract(&_Zrxtrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.contract.Transact(opts, "trade", srcToken, amount, destToken, callDataHex, _forwarder)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeSession) Trade(srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.Contract.Trade(&_Zrxtrade.TransactOpts, srcToken, amount, destToken, callDataHex, _forwarder)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeTransactorSession) Trade(srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.Contract.Trade(&_Zrxtrade.TransactOpts, srcToken, amount, destToken, callDataHex, _forwarder)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeTransactor) WithdrawWrapETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.contract.Transact(opts, "withdrawWrapETH", amount)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeSession) WithdrawWrapETH(amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.Contract.WithdrawWrapETH(&_Zrxtrade.TransactOpts, amount)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeTransactorSession) WithdrawWrapETH(amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.Contract.WithdrawWrapETH(&_Zrxtrade.TransactOpts, amount)
}
