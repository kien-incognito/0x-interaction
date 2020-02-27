// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package incmodemock

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

// IncmodemockABI is the input ABI used to generate the binding from.
const IncmodemockABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incognitoProxyAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"MoveAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequest\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"contractIncognito\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newVault\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"heights\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"heights\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newVault\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"moveAssets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"updateIncognitoProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IncmodemockBin is the compiled bytecode used for deploying new contracts.
var IncmodemockBin = "0x60806040523480156200001157600080fd5b5060405162004f1b38038062004f1b833981810160405262000037919081019062000185565b82806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160146101000a81548160ff0219169083151502179055506301e1338042016002819055505081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505062000229565b6000815190506200017f816200020f565b92915050565b6000806000606084860312156200019b57600080fd5b6000620001ab868287016200016e565b9350506020620001be868287016200016e565b9250506040620001d1868287016200016e565b9150509250925092565b6000620001e882620001ef565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200021a81620001db565b81146200022657600080fd5b50565b614ce280620002396000396000f3fe6080604052600436106101f95760003560e01c80637e16e6e11161010d578063ce5494bb116100a0578063f75b98ce1161006f578063f75b98ce1461071b578063f851a44014610758578063f8b3974514610783578063fa84702e146107ac578063faf7f494146107d7576101f9565b8063ce5494bb1461063b578063cf54aaa014610664578063dca40d9e146106a1578063e4bd7074146106de576101f9565b80639714378c116100dc5780639714378c146105b15780639e6371ba146105da578063a0c0f9b814610603578063a26e11861461061f576101f9565b80637e16e6e1146105035780638456cb591461054457806388aaf0c81461055b5780638a98453814610586576101f9565b806358bc83371161019057806365b5a00f1161015f57806365b5a00f146103f65780636ff968c31461043357806370a082311461045e578063749c5f861461049b57806379599f96146104d8576101f9565b806358bc8337146103475780635a67cb87146103725780635c975abb1461038e5780636304541c146103b9576101f9565b80633f4ba83a116101cc5780633f4ba83a146102b35780633fec6b40146102ca5780634e71d92d146103075780635654b6c81461031e576101f9565b80630c4f5039146101fb5780631ea1940e146102245780633a51913d146102615780633cb4b58a1461028a575b005b34801561020757600080fd5b50610222600480360361021d91908101906139da565b610800565b005b34801561023057600080fd5b5061024b60048036036102469190810190613a44565b610bc7565b6040516102589190614492565b60405180910390f35b34801561026d57600080fd5b506102886004803603610283919081019061378c565b610be7565b005b34801561029657600080fd5b506102b160048036036102ac9190810190613aae565b610d7a565b005b3480156102bf57600080fd5b506102c8611086565b005b3480156102d657600080fd5b506102f160048036036102ec9190810190613c33565b6111b8565b6040516102fe9190614321565b60405180910390f35b34801561031357600080fd5b5061031c611246565b005b34801561032a57600080fd5b5061034560048036036103409190810190613cc8565b6113d6565b005b34801561035357600080fd5b5061035c611647565b6040516103699190614321565b60405180910390f35b61038c60048036036103879190810190613973565b61164c565b005b34801561039a57600080fd5b506103a361189b565b6040516103b09190614492565b60405180910390f35b3480156103c557600080fd5b506103e060048036036103db919081019061378c565b6118ae565b6040516103ed91906146d8565b60405180910390f35b34801561040257600080fd5b5061041d6004803603610418919081019061381a565b6118c6565b60405161042a91906146d8565b60405180910390f35b34801561043f57600080fd5b506104486118eb565b6040516104559190614321565b60405180910390f35b34801561046a57600080fd5b506104856004803603610480919081019061378c565b611911565b60405161049291906146d8565b60405180910390f35b3480156104a757600080fd5b506104c260048036036104bd9190810190613a44565b6119e0565b6040516104cf9190614492565b60405180910390f35b3480156104e457600080fd5b506104ed611b24565b6040516104fa91906146d8565b60405180910390f35b34801561050f57600080fd5b5061052a60048036036105259190810190613a6d565b611b2a565b60405161053b95949392919061470e565b60405180910390f35b34801561055057600080fd5b50610559611ba6565b005b34801561056757600080fd5b50610570611d1c565b60405161057d9190614357565b60405180910390f35b34801561059257600080fd5b5061059b611d42565b6040516105a891906145e2565b60405180910390f35b3480156105bd57600080fd5b506105d860048036036105d39190810190613d6f565b611d68565b005b3480156105e657600080fd5b5061060160048036036105fc919081019061378c565b611ec9565b005b61061d600480360361061891908101906138a5565b611fe0565b005b61063960048036036106349190810190613c87565b6122a6565b005b34801561064757600080fd5b50610662600480360361065d91908101906137b5565b61234e565b005b34801561067057600080fd5b5061068b6004803603610686919081019061378c565b6124e1565b60405161069891906146f3565b60405180910390f35b3480156106ad57600080fd5b506106c860048036036106c39190810190613a44565b6125a5565b6040516106d59190614492565b60405180910390f35b3480156106ea57600080fd5b5061070560048036036107009190810190613a44565b6125c5565b6040516107129190614492565b60405180910390f35b34801561072757600080fd5b50610742600480360361073d919081019061381a565b612709565b60405161074f91906146d8565b60405180910390f35b34801561076457600080fd5b5061076d612790565b60405161077a9190614321565b60405180910390f35b34801561078f57600080fd5b506107aa60048036036107a59190810190613aae565b6127b5565b005b3480156107b857600080fd5b506107c1612b46565b6040516107ce91906145fd565b60405180910390f35b3480156107e357600080fd5b506107fe60048036036107f99190810190613856565b612b6c565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610886906146b8565b60405180910390fd5b600160149054906101000a900460ff166108de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d590614618565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561093a57600080fd5b60008090505b8151811015610b8c57600073ffffffffffffffffffffffffffffffffffffffff1682828151811061096d57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614156109ff57600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156109f9573d6000803e3d6000fd5b50610b7f565b6000828281518110610a0d57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610a4d919061433c565b60206040518083038186803b158015610a6557600080fd5b505afa158015610a79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610a9d9190810190613d98565b90506000811115610b7d57828281518110610ab457fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610b189291906143a9565b602060405180830381600087803b158015610b3257600080fd5b505af1158015610b46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b6a9190810190613a1b565b50610b73612c48565b610b7c57600080fd5b5b505b8080600101915050610940565b507f492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce65881604051610bbc9190614470565b60405180910390a150565b60046020528060005260406000206000915054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6d906146b8565b60405180910390fd5b600160149054906101000a900460ff16610cc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbc90614618565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610cff57600080fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c344681604051610d6f9190614321565b60405180910390a150565b600160149054906101000a900460ff1615610dca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc190614698565b60405180910390fd5b6000806000806000610ddb8f611b2a565b9450945094509450945060618560ff16148015610dfb575060018460ff16145b610e0457600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610e8c57600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548101471015610e8757600080fd5b610f8c565b6000610e97846124e1565b905060098160ff161115610eb45760098160ff1603600a0a820291505b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482018473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610f2f919061433c565b60206040518083038186803b158015610f4757600080fd5b505afa158015610f5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f7f9190810190613d98565b1015610f8a57600080fd5b505b610f9e8f8f8f8f8f8f8f8f8f8f612c86565b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555080600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505050505050505050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161110c906146b8565b60405180910390fd5b600160149054906101000a900460ff16611164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115b90614618565b60405180910390fd5b6000600160146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516111ae919061433c565b60405180910390a1565b6000806000806020860151915060408601519250601b866040815181106111db57fe5b602001015160f81c60f81b60f81c0190506001858284866040516000815260200160405260405161120f949392919061459d565b6020604051602081039080840390855afa158015611231573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600254421061128a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128190614638565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461131a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131190614678565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516113cc9190614321565b60405180910390a1565b6113df816125c5565b156113e957600080fd5b60006113f583836111b8565b905083600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561148057600080fd5b6000849050600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156114cf57633b9aca0085816114c757fe5b049050611500565b60006114da876124e1565b905060098160ff1611156114fe5760098160ff1603600a0a86816114fa57fe5b0491505b505b60016004600085815260200190815260200160002060006101000a81548160ff02191690831515021790555084600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555084600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e86888360405161163693929190614409565b60405180910390a150505050505050565b600081565b600160149054906101000a900460ff161561169c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169390614698565b60405180910390fd5b600083905060006116ac856124e1565b905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016116e9919061433c565b60206040518083038186803b15801561170157600080fd5b505afa158015611715573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506117399190810190613d98565b9050600085905060098360ff1611156117775760098360ff1603600a0a818161175e57fe5b04905060098360ff1603600a0a828161177357fe5b0491505b670de0b6b3a764000081111580156117975750670de0b6b3a76400008211155b80156117ad5750670de0b6b3a764000082820111155b6117b657600080fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330896040518463ffffffff1660e01b81526004016117f393929190614372565b602060405180830381600087803b15801561180d57600080fd5b505af1158015611821573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506118459190810190613a1b565b5061184e612c48565b61185757600080fd5b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e87868360405161188a93929190614409565b60405180910390a150505050505050565b600160149054906101000a900460ff1681565b60066020528060005260406000206000915090505481565b6005602052816000526040600020602052806000526040600020600091509150505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561194f574790506119db565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401611988919061433c565b60206040518083038186803b1580156119a057600080fd5b505afa1580156119b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506119d89190810190613d98565b90505b919050565b60006003600083815260200190815260200160002060009054906101000a900460ff1615611a115760019050611b1f565b600073ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611a715760009050611b1f565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663749c5f86836040518263ffffffff1660e01b8152600401611acc9190614582565b60206040518083038186803b158015611ae457600080fd5b505afa158015611af8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611b1c9190810190613a1b565b90505b919050565b60025481565b60008060008060008086600081518110611b4057fe5b602001015160f81c60f81b60f81c9050600087600181518110611b5f57fe5b602001015160f81c60f81b60f81c9050600080600060228b0151925060428b0151915060628b01519050848484848499509950995099509950505050505091939590929450565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c2c906146b8565b60405180910390fd5b600160149054906101000a900460ff1615611c85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c7c90614698565b60405180910390fd5b6002544210611cc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cc090614638565b60405180910390fd5b60018060146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051611d12919061433c565b60405180910390a1565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611df7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dee906146b8565b60405180910390fd5b6002544210611e3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e3290614638565b60405180910390fd5b61016e8110611e7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e7690614658565b60405180910390fd5b620151808102600254016002819055507f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e881604051611ebe91906146d8565b60405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f4f906146b8565b60405180910390fd5b6002544210611f9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f9390614638565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611fe9826125c5565b15611ff357600080fd5b6000611fff82846111b8565b90506000349050600073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614156120455787810190506120dc565b8873ffffffffffffffffffffffffffffffffffffffff1663a9059cbb878a6040518363ffffffff1660e01b8152600401612080929190614447565b602060405180830381600087803b15801561209a57600080fd5b505af11580156120ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506120d29190810190613a1b565b6120db57600080fd5b5b60006120ea8883888a613020565b905088600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555088600660008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600660008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050505050505050505050565b600160149054906101000a900460ff16156122f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ed90614698565b60405180910390fd5b6b033b2e3c9fd0803ce800000047111561230f57600080fd5b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e6000823460405161234393929190614409565b60405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146123dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123d4906146b8565b60405180910390fd5b600160149054906101000a900460ff1661242c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161242390614618565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561246657600080fd5b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a816040516124d6919061433c565b60405180910390a150565b60008082905060008090508173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561253257600080fd5b505afa158015612546573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061256a9190810190613d98565b503d60008114612585576020811461258e576000915061259a565b6000915061259a565b60206000803e60005191505b508092505050919050565b60036020528060005260406000206000915054906101000a900460ff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff16156125f65760019050612704565b600073ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156126565760009050612704565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4bd7074836040518263ffffffff1660e01b81526004016126b19190614582565b60206040518083038186803b1580156126c957600080fd5b505afa1580156126dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506127019190810190613a1b565b90505b919050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff1615612805576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127fc90614698565b60405180910390fd5b60008060008060006128168f611b2a565b9450945094509450945060488560ff16148015612836575060018460ff16145b61283f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156128c757600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481014710156128c257600080fd5b6129c7565b60006128d2846124e1565b905060098160ff1611156128ef5760098160ff1603600a0a820291505b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482018473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161296a919061433c565b60206040518083038186803b15801561298257600080fd5b505afa158015612996573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506129ba9190810190613d98565b10156129c557600080fd5b505b6129d98f8f8f8f8f8f8f8f8f8f612c86565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415612a5a578173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015612a54573d6000803e3d6000fd5b50612afa565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401612a959291906143a9565b602060405180830381600087803b158015612aaf57600080fd5b505af1158015612ac3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612ae79190810190613a1b565b50612af0612c48565b612af957600080fd5b5b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb838383604051612b2d939291906143d2565b60405180910390a1505050505050505050505050505050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505050565b600080600090503d60008114612c655760208114612c6e57612c7a565b60019150612c7a565b60206000803e60005191505b50600081141591505090565b60008a80519060200120905060008b8b600060028110612ca257fe5b6020020151604051602001612cb89291906142f9565b60405160208183030381529060405280519060200120905060008c8c600160028110612ce057fe5b6020020151604051602001612cf69291906142f9565b604051602081830303815290604052805190602001209050612d17836119e0565b15612d2157600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f65d21166001848f600060028110612d6f57fe5b60200201518f600060028110612d8157fe5b60200201518f600060028110612d9357fe5b60200201518f600060028110612da557fe5b60200201518f600060028110612db757fe5b60200201518f600060028110612dc957fe5b60200201518f600060028110612ddb57fe5b60200201518f600060028110612ded57fe5b60200201518f600060028110612dff57fe5b60200201516040518c63ffffffff1660e01b8152600401612e2a9b9a999897969594939291906144ad565b60206040518083038186803b158015612e4257600080fd5b505afa158015612e56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612e7a9190810190613a1b565b612e8357600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f65d21166000838f600160028110612ed157fe5b60200201518f600160028110612ee357fe5b60200201518f600160028110612ef557fe5b60200201518f600160028110612f0757fe5b60200201518f600160028110612f1957fe5b60200201518f600160028110612f2b57fe5b60200201518f600160028110612f3d57fe5b60200201518f600160028110612f4f57fe5b60200201518f600160028110612f6157fe5b60200201516040518c63ffffffff1660e01b8152600401612f8c9b9a999897969594939291906144ad565b60206040518083038186803b158015612fa457600080fd5b505afa158015612fb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612fdc9190810190613a1b565b612fe557600080fd5b60016003600085815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050505050505050505050565b60008061302c86611911565b9050600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156130695734810390505b8447101561307657600080fd5b600060608473ffffffffffffffffffffffffffffffffffffffff1687876040516130a091906142e2565b60006040518083038185875af1925050503d80600081146130dd576040519150601f19603f3d011682016040523d82523d6000602084013e6130e2565b606091505b5091509150816130f157600080fd5b6000808280602001905161310891908101906137de565b915091508095505050505050949350505050565b60008135905061312b81614c15565b92915050565b60008135905061314081614c2c565b92915050565b60008151905061315581614c2c565b92915050565b600082601f83011261316c57600080fd5b813561317f61317a8261478e565b614761565b915081818352602084019350602081019050838560208402820111156131a457600080fd5b60005b838110156131d457816131ba888261311c565b8452602084019350602083019250506001810190506131a7565b5050505092915050565b600082601f8301126131ef57600080fd5b60026132026131fd826147b6565b614761565b9150818360005b83811015613239578135860161321f8882613372565b845260208401935060208301925050600181019050613209565b5050505092915050565b600082601f83011261325457600080fd5b6002613267613262826147d8565b614761565b9150818360005b8381101561329e57813586016132848882613469565b84526020840193506020830192505060018101905061326e565b5050505092915050565b600082601f8301126132b957600080fd5b60026132cc6132c7826147fa565b614761565b9150818360005b8381101561330357813586016132e98882613560565b8452602084019350602083019250506001810190506132d3565b5050505092915050565b600082601f83011261331e57600080fd5b600261333161332c8261481c565b614761565b9150818360005b83811015613368578135860161334e88826135e3565b845260208401935060208301925050600181019050613338565b5050505092915050565b600082601f83011261338357600080fd5b81356133966133918261483e565b614761565b915081818352602084019350602081019050838560208402820111156133bb57600080fd5b60005b838110156133eb57816133d18882613666565b8452602084019350602083019250506001810190506133be565b5050505092915050565b600082601f83011261340657600080fd5b600261341961341482614866565b614761565b9150818385602084028201111561342f57600080fd5b60005b8381101561345f57816134458882613690565b845260208401935060208301925050600181019050613432565b5050505092915050565b600082601f83011261347a57600080fd5b813561348d61348882614888565b614761565b915081818352602084019350602081019050838560208402820111156134b257600080fd5b60005b838110156134e257816134c88882613690565b8452602084019350602083019250506001810190506134b5565b5050505092915050565b600082601f8301126134fd57600080fd5b600261351061350b826148b0565b614761565b9150818385602084028201111561352657600080fd5b60005b83811015613556578161353c888261374d565b845260208401935060208301925050600181019050613529565b5050505092915050565b600082601f83011261357157600080fd5b813561358461357f826148d2565b614761565b915081818352602084019350602081019050838560208402820111156135a957600080fd5b60005b838110156135d957816135bf888261374d565b8452602084019350602083019250506001810190506135ac565b5050505092915050565b600082601f8301126135f457600080fd5b8135613607613602826148fa565b614761565b9150818183526020840193506020810190508385602084028201111561362c57600080fd5b60005b8381101561365c57816136428882613777565b84526020840193506020830192505060018101905061362f565b5050505092915050565b60008135905061367581614c43565b92915050565b60008151905061368a81614c43565b92915050565b60008135905061369f81614c5a565b92915050565b600082601f8301126136b657600080fd5b81356136c96136c482614922565b614761565b915080825260208301602083018583830111156136e557600080fd5b6136f0838284614bb8565b50505092915050565b600082601f83011261370a57600080fd5b813561371d6137188261494e565b614761565b9150808252602083016020830185838301111561373957600080fd5b613744838284614bb8565b50505092915050565b60008135905061375c81614c71565b92915050565b60008151905061377181614c71565b92915050565b60008135905061378681614c88565b92915050565b60006020828403121561379e57600080fd5b60006137ac8482850161311c565b91505092915050565b6000602082840312156137c757600080fd5b60006137d584828501613131565b91505092915050565b600080604083850312156137f157600080fd5b60006137ff85828601613146565b925050602061381085828601613762565b9150509250929050565b6000806040838503121561382d57600080fd5b600061383b8582860161311c565b925050602061384c8582860161311c565b9150509250929050565b60008060006060848603121561386b57600080fd5b60006138798682870161311c565b935050602061388a8682870161311c565b925050604061389b8682870161374d565b9150509250925092565b600080600080600080600060e0888a0312156138c057600080fd5b60006138ce8a828b0161311c565b97505060206138df8a828b0161374d565b96505060406138f08a828b0161311c565b95505060606139018a828b0161311c565b945050608088013567ffffffffffffffff81111561391e57600080fd5b61392a8a828b016136a5565b93505060a061393b8a828b01613690565b92505060c088013567ffffffffffffffff81111561395857600080fd5b6139648a828b016136a5565b91505092959891949750929550565b60008060006060848603121561398857600080fd5b60006139968682870161311c565b93505060206139a78682870161374d565b925050604084013567ffffffffffffffff8111156139c457600080fd5b6139d0868287016136f9565b9150509250925092565b6000602082840312156139ec57600080fd5b600082013567ffffffffffffffff811115613a0657600080fd5b613a128482850161315b565b91505092915050565b600060208284031215613a2d57600080fd5b6000613a3b8482850161367b565b91505092915050565b600060208284031215613a5657600080fd5b6000613a6484828501613690565b91505092915050565b600060208284031215613a7f57600080fd5b600082013567ffffffffffffffff811115613a9957600080fd5b613aa5848285016136a5565b91505092915050565b6000806000806000806000806000806101a08b8d031215613ace57600080fd5b60008b013567ffffffffffffffff811115613ae857600080fd5b613af48d828e016136a5565b9a50506020613b058d828e016134ec565b99505060608b013567ffffffffffffffff811115613b2257600080fd5b613b2e8d828e01613243565b98505060808b013567ffffffffffffffff811115613b4b57600080fd5b613b578d828e016131de565b97505060a0613b688d828e016133f5565b96505060e0613b798d828e016133f5565b9550506101208b013567ffffffffffffffff811115613b9757600080fd5b613ba38d828e016132a8565b9450506101408b013567ffffffffffffffff811115613bc157600080fd5b613bcd8d828e0161330d565b9350506101608b013567ffffffffffffffff811115613beb57600080fd5b613bf78d828e01613243565b9250506101808b013567ffffffffffffffff811115613c1557600080fd5b613c218d828e01613243565b9150509295989b9194979a5092959850565b60008060408385031215613c4657600080fd5b600083013567ffffffffffffffff811115613c6057600080fd5b613c6c858286016136a5565b9250506020613c7d85828601613690565b9150509250929050565b600060208284031215613c9957600080fd5b600082013567ffffffffffffffff811115613cb357600080fd5b613cbf848285016136f9565b91505092915050565b600080600080600060a08688031215613ce057600080fd5b600086013567ffffffffffffffff811115613cfa57600080fd5b613d06888289016136f9565b9550506020613d178882890161311c565b9450506040613d288882890161374d565b935050606086013567ffffffffffffffff811115613d4557600080fd5b613d51888289016136a5565b9250506080613d6288828901613690565b9150509295509295909350565b600060208284031215613d8157600080fd5b6000613d8f8482850161374d565b91505092915050565b600060208284031215613daa57600080fd5b6000613db884828501613762565b91505092915050565b6000613dcd8383613e57565b60208301905092915050565b6000613de5838361404b565b60208301905092915050565b6000613dfd8383614069565b60208301905092915050565b6000613e15838361428f565b60208301905092915050565b6000613e2d83836142c4565b60208301905092915050565b613e4281614b3a565b82525050565b613e5181614adb565b82525050565b613e6081614ac9565b82525050565b613e6f81614ac9565b82525050565b6000613e80826149ca565b613e8a8185614a58565b9350613e958361497a565b8060005b83811015613ec6578151613ead8882613dc1565b9750613eb883614a17565b925050600181019050613e99565b5085935050505092915050565b6000613ede826149d5565b613ee88185614a69565b9350613ef38361498a565b8060005b83811015613f24578151613f0b8882613dd9565b9750613f1683614a24565b925050600181019050613ef7565b5085935050505092915050565b6000613f3c826149e0565b613f468185614a7a565b9350613f518361499a565b8060005b83811015613f82578151613f698882613df1565b9750613f7483614a31565b925050600181019050613f55565b5085935050505092915050565b6000613f9a826149eb565b613fa48185614a8b565b9350613faf836149aa565b8060005b83811015613fe0578151613fc78882613e09565b9750613fd283614a3e565b925050600181019050613fb3565b5085935050505092915050565b6000613ff8826149f6565b6140028185614a9c565b935061400d836149ba565b8060005b8381101561403e5781516140258882613e21565b975061403083614a4b565b925050600181019050614011565b5085935050505092915050565b61405481614aed565b82525050565b61406381614aed565b82525050565b61407281614af9565b82525050565b61408181614af9565b82525050565b600061409282614a01565b61409c8185614aad565b93506140ac818560208601614bc7565b80840191505092915050565b6140c181614b4c565b82525050565b6140d081614b70565b82525050565b60006140e182614a0c565b6140eb8185614ab8565b93506140fb818560208601614bc7565b61410481614c04565b840191505092915050565b600061411c601483614ab8565b91507f6e6f7420706175736564207269676874206e6f770000000000000000000000006000830152602082019050919050565b600061415c600783614ab8565b91507f65787069726564000000000000000000000000000000000000000000000000006000830152602082019050919050565b600061419c601a83614ab8565b91507f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e670000000000006000830152602082019050919050565b60006141dc600c83614ab8565b91507f756e617574686f72697a656400000000000000000000000000000000000000006000830152602082019050919050565b600061421c601083614ab8565b91507f706175736564207269676874206e6f77000000000000000000000000000000006000830152602082019050919050565b600061425c600983614ab8565b91507f6e6f742061646d696e00000000000000000000000000000000000000000000006000830152602082019050919050565b61429881614b23565b82525050565b6142a781614b23565b82525050565b6142be6142b982614b23565b614bfa565b82525050565b6142cd81614b2d565b82525050565b6142dc81614b2d565b82525050565b60006142ee8284614087565b915081905092915050565b60006143058285614087565b915061431182846142ad565b6020820191508190509392505050565b60006020820190506143366000830184613e66565b92915050565b60006020820190506143516000830184613e39565b92915050565b600060208201905061436c6000830184613e48565b92915050565b60006060820190506143876000830186613e39565b6143946020830185613e39565b6143a1604083018461429e565b949350505050565b60006040820190506143be6000830185613e39565b6143cb602083018461429e565b9392505050565b60006060820190506143e76000830186613e66565b6143f46020830185613e39565b614401604083018461429e565b949350505050565b600060608201905061441e6000830186613e66565b818103602083015261443081856140d6565b905061443f604083018461429e565b949350505050565b600060408201905061445c6000830185613e66565b614469602083018461429e565b9392505050565b6000602082019050818103600083015261448a8184613e75565b905092915050565b60006020820190506144a7600083018461405a565b92915050565b6000610160820190506144c3600083018e61405a565b6144d0602083018d614078565b6144dd604083018c61429e565b81810360608301526144ef818b613f31565b90508181036080830152614503818a613ed3565b905061451260a0830189614078565b61451f60c0830188614078565b81810360e08301526145318187613f8f565b90508181036101008301526145468186613fed565b905081810361012083015261455b8185613f31565b90508181036101408301526145708184613f31565b90509c9b505050505050505050505050565b60006020820190506145976000830184614078565b92915050565b60006080820190506145b26000830187614078565b6145bf60208301866142d3565b6145cc6040830185614078565b6145d96060830184614078565b95945050505050565b60006020820190506145f760008301846140b8565b92915050565b600060208201905061461260008301846140c7565b92915050565b600060208201905081810360008301526146318161410f565b9050919050565b600060208201905081810360008301526146518161414f565b9050919050565b600060208201905081810360008301526146718161418f565b9050919050565b60006020820190508181036000830152614691816141cf565b9050919050565b600060208201905081810360008301526146b18161420f565b9050919050565b600060208201905081810360008301526146d18161424f565b9050919050565b60006020820190506146ed600083018461429e565b92915050565b600060208201905061470860008301846142d3565b92915050565b600060a08201905061472360008301886142d3565b61473060208301876142d3565b61473d6040830186613e66565b61474a6060830185613e48565b614757608083018461429e565b9695505050505050565b6000604051905081810181811067ffffffffffffffff8211171561478457600080fd5b8060405250919050565b600067ffffffffffffffff8211156147a557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156147cd57600080fd5b602082029050919050565b600067ffffffffffffffff8211156147ef57600080fd5b602082029050919050565b600067ffffffffffffffff82111561481157600080fd5b602082029050919050565b600067ffffffffffffffff82111561483357600080fd5b602082029050919050565b600067ffffffffffffffff82111561485557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561487d57600080fd5b602082029050919050565b600067ffffffffffffffff82111561489f57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156148c757600080fd5b602082029050919050565b600067ffffffffffffffff8211156148e957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561491157600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561493957600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff82111561496557600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000614ad482614b03565b9050919050565b6000614ae682614b03565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000614b4582614b94565b9050919050565b6000614b5782614b5e565b9050919050565b6000614b6982614b03565b9050919050565b6000614b7b82614b82565b9050919050565b6000614b8d82614b03565b9050919050565b6000614b9f82614ba6565b9050919050565b6000614bb182614b03565b9050919050565b82818337600083830152505050565b60005b83811015614be5578082015181840152602081019050614bca565b83811115614bf4576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b614c1e81614ac9565b8114614c2957600080fd5b50565b614c3581614adb565b8114614c4057600080fd5b50565b614c4c81614aed565b8114614c5757600080fd5b50565b614c6381614af9565b8114614c6e57600080fd5b50565b614c7a81614b23565b8114614c8557600080fd5b50565b614c9181614b2d565b8114614c9c57600080fd5b5056fea365627a7a72315820e4a656dc19b93ef3b4f8a9106adfcdd39c1e7c5fff4799d6d36cad2b0a1f90656c6578706572696d656e74616cf564736f6c63430005100040"

// DeployIncmodemock deploys a new Ethereum contract, binding an instance of Incmodemock to it.
func DeployIncmodemock(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, incognitoProxyAddress common.Address, _prevVault common.Address) (common.Address, *types.Transaction, *Incmodemock, error) {
	parsed, err := abi.JSON(strings.NewReader(IncmodemockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncmodemockBin), backend, admin, incognitoProxyAddress, _prevVault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Incmodemock{IncmodemockCaller: IncmodemockCaller{contract: contract}, IncmodemockTransactor: IncmodemockTransactor{contract: contract}, IncmodemockFilterer: IncmodemockFilterer{contract: contract}}, nil
}

// Incmodemock is an auto generated Go binding around an Ethereum contract.
type Incmodemock struct {
	IncmodemockCaller     // Read-only binding to the contract
	IncmodemockTransactor // Write-only binding to the contract
	IncmodemockFilterer   // Log filterer for contract events
}

// IncmodemockCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncmodemockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncmodemockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncmodemockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncmodemockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncmodemockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncmodemockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncmodemockSession struct {
	Contract     *Incmodemock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncmodemockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncmodemockCallerSession struct {
	Contract *IncmodemockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IncmodemockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncmodemockTransactorSession struct {
	Contract     *IncmodemockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IncmodemockRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncmodemockRaw struct {
	Contract *Incmodemock // Generic contract binding to access the raw methods on
}

// IncmodemockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncmodemockCallerRaw struct {
	Contract *IncmodemockCaller // Generic read-only contract binding to access the raw methods on
}

// IncmodemockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncmodemockTransactorRaw struct {
	Contract *IncmodemockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncmodemock creates a new instance of Incmodemock, bound to a specific deployed contract.
func NewIncmodemock(address common.Address, backend bind.ContractBackend) (*Incmodemock, error) {
	contract, err := bindIncmodemock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incmodemock{IncmodemockCaller: IncmodemockCaller{contract: contract}, IncmodemockTransactor: IncmodemockTransactor{contract: contract}, IncmodemockFilterer: IncmodemockFilterer{contract: contract}}, nil
}

// NewIncmodemockCaller creates a new read-only instance of Incmodemock, bound to a specific deployed contract.
func NewIncmodemockCaller(address common.Address, caller bind.ContractCaller) (*IncmodemockCaller, error) {
	contract, err := bindIncmodemock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncmodemockCaller{contract: contract}, nil
}

// NewIncmodemockTransactor creates a new write-only instance of Incmodemock, bound to a specific deployed contract.
func NewIncmodemockTransactor(address common.Address, transactor bind.ContractTransactor) (*IncmodemockTransactor, error) {
	contract, err := bindIncmodemock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncmodemockTransactor{contract: contract}, nil
}

// NewIncmodemockFilterer creates a new log filterer instance of Incmodemock, bound to a specific deployed contract.
func NewIncmodemockFilterer(address common.Address, filterer bind.ContractFilterer) (*IncmodemockFilterer, error) {
	contract, err := bindIncmodemock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncmodemockFilterer{contract: contract}, nil
}

// bindIncmodemock binds a generic wrapper to an already deployed contract.
func bindIncmodemock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncmodemockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incmodemock *IncmodemockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incmodemock.Contract.IncmodemockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incmodemock *IncmodemockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incmodemock.Contract.IncmodemockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incmodemock *IncmodemockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incmodemock.Contract.IncmodemockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incmodemock *IncmodemockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incmodemock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incmodemock *IncmodemockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incmodemock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incmodemock *IncmodemockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incmodemock.Contract.contract.Transact(opts, method, params...)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() constant returns(address)
func (_Incmodemock *IncmodemockCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "ETH_TOKEN")
	return *ret0, err
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() constant returns(address)
func (_Incmodemock *IncmodemockSession) ETHTOKEN() (common.Address, error) {
	return _Incmodemock.Contract.ETHTOKEN(&_Incmodemock.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) ETHTOKEN() (common.Address, error) {
	return _Incmodemock.Contract.ETHTOKEN(&_Incmodemock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Incmodemock *IncmodemockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Incmodemock *IncmodemockSession) Admin() (common.Address, error) {
	return _Incmodemock.Contract.Admin(&_Incmodemock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) Admin() (common.Address, error) {
	return _Incmodemock.Contract.Admin(&_Incmodemock.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) constant returns(uint256)
func (_Incmodemock *IncmodemockCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) constant returns(uint256)
func (_Incmodemock *IncmodemockSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.BalanceOf(&_Incmodemock.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) constant returns(uint256)
func (_Incmodemock *IncmodemockCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.BalanceOf(&_Incmodemock.CallOpts, token)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Incmodemock *IncmodemockCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Incmodemock *IncmodemockSession) Expire() (*big.Int, error) {
	return _Incmodemock.Contract.Expire(&_Incmodemock.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Incmodemock *IncmodemockCallerSession) Expire() (*big.Int, error) {
	return _Incmodemock.Contract.Expire(&_Incmodemock.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) constant returns(uint8)
func (_Incmodemock *IncmodemockCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "getDecimals", token)
	return *ret0, err
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) constant returns(uint8)
func (_Incmodemock *IncmodemockSession) GetDecimals(token common.Address) (uint8, error) {
	return _Incmodemock.Contract.GetDecimals(&_Incmodemock.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) constant returns(uint8)
func (_Incmodemock *IncmodemockCallerSession) GetDecimals(token common.Address) (uint8, error) {
	return _Incmodemock.Contract.GetDecimals(&_Incmodemock.CallOpts, token)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) constant returns(uint256)
func (_Incmodemock *IncmodemockCaller) GetDepositedBalance(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "getDepositedBalance", token, owner)
	return *ret0, err
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) constant returns(uint256)
func (_Incmodemock *IncmodemockSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.GetDepositedBalance(&_Incmodemock.CallOpts, token, owner)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) constant returns(uint256)
func (_Incmodemock *IncmodemockCallerSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.GetDepositedBalance(&_Incmodemock.CallOpts, token, owner)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Incmodemock *IncmodemockCaller) Incognito(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "incognito")
	return *ret0, err
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Incmodemock *IncmodemockSession) Incognito() (common.Address, error) {
	return _Incmodemock.Contract.Incognito(&_Incmodemock.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) Incognito() (common.Address, error) {
	return _Incmodemock.Contract.Incognito(&_Incmodemock.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Incmodemock *IncmodemockCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "isSigDataUsed", hash)
	return *ret0, err
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Incmodemock *IncmodemockSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Incmodemock.Contract.IsSigDataUsed(&_Incmodemock.CallOpts, hash)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) constant returns(bool)
func (_Incmodemock *IncmodemockCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Incmodemock.Contract.IsSigDataUsed(&_Incmodemock.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) constant returns(bool)
func (_Incmodemock *IncmodemockCaller) IsWithdrawed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "isWithdrawed", hash)
	return *ret0, err
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) constant returns(bool)
func (_Incmodemock *IncmodemockSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Incmodemock.Contract.IsWithdrawed(&_Incmodemock.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) constant returns(bool)
func (_Incmodemock *IncmodemockCallerSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Incmodemock.Contract.IsWithdrawed(&_Incmodemock.CallOpts, hash)
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() constant returns(address)
func (_Incmodemock *IncmodemockCaller) NewVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "newVault")
	return *ret0, err
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() constant returns(address)
func (_Incmodemock *IncmodemockSession) NewVault() (common.Address, error) {
	return _Incmodemock.Contract.NewVault(&_Incmodemock.CallOpts)
}

// NewVault is a free data retrieval call binding the contract method 0x88aaf0c8.
//
// Solidity: function newVault() constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) NewVault() (common.Address, error) {
	return _Incmodemock.Contract.NewVault(&_Incmodemock.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Incmodemock *IncmodemockCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Incmodemock.contract.Call(opts, out, "parseBurnInst", inst)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Incmodemock *IncmodemockSession) ParseBurnInst(inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	return _Incmodemock.Contract.ParseBurnInst(&_Incmodemock.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Incmodemock *IncmodemockCallerSession) ParseBurnInst(inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	return _Incmodemock.Contract.ParseBurnInst(&_Incmodemock.CallOpts, inst)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Incmodemock *IncmodemockCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Incmodemock *IncmodemockSession) Paused() (bool, error) {
	return _Incmodemock.Contract.Paused(&_Incmodemock.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Incmodemock *IncmodemockCallerSession) Paused() (bool, error) {
	return _Incmodemock.Contract.Paused(&_Incmodemock.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() constant returns(address)
func (_Incmodemock *IncmodemockCaller) PrevVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "prevVault")
	return *ret0, err
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() constant returns(address)
func (_Incmodemock *IncmodemockSession) PrevVault() (common.Address, error) {
	return _Incmodemock.Contract.PrevVault(&_Incmodemock.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) PrevVault() (common.Address, error) {
	return _Incmodemock.Contract.PrevVault(&_Incmodemock.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Incmodemock *IncmodemockCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "sigDataUsed", arg0)
	return *ret0, err
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Incmodemock *IncmodemockSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Incmodemock.Contract.SigDataUsed(&_Incmodemock.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) constant returns(bool)
func (_Incmodemock *IncmodemockCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Incmodemock.Contract.SigDataUsed(&_Incmodemock.CallOpts, arg0)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Incmodemock *IncmodemockCaller) SigToAddress(opts *bind.CallOpts, signData []byte, hash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "sigToAddress", signData, hash)
	return *ret0, err
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Incmodemock *IncmodemockSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Incmodemock.Contract.SigToAddress(&_Incmodemock.CallOpts, signData, hash)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Incmodemock.Contract.SigToAddress(&_Incmodemock.CallOpts, signData, hash)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Incmodemock *IncmodemockCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Incmodemock *IncmodemockSession) Successor() (common.Address, error) {
	return _Incmodemock.Contract.Successor(&_Incmodemock.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Incmodemock *IncmodemockCallerSession) Successor() (common.Address, error) {
	return _Incmodemock.Contract.Successor(&_Incmodemock.CallOpts)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) constant returns(uint256)
func (_Incmodemock *IncmodemockCaller) TotalDepositedToSCAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "totalDepositedToSCAmount", arg0)
	return *ret0, err
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) constant returns(uint256)
func (_Incmodemock *IncmodemockSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.TotalDepositedToSCAmount(&_Incmodemock.CallOpts, arg0)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) constant returns(uint256)
func (_Incmodemock *IncmodemockCallerSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.TotalDepositedToSCAmount(&_Incmodemock.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) constant returns(uint256)
func (_Incmodemock *IncmodemockCaller) WithdrawRequests(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "withdrawRequests", arg0, arg1)
	return *ret0, err
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) constant returns(uint256)
func (_Incmodemock *IncmodemockSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.WithdrawRequests(&_Incmodemock.CallOpts, arg0, arg1)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) constant returns(uint256)
func (_Incmodemock *IncmodemockCallerSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Incmodemock.Contract.WithdrawRequests(&_Incmodemock.CallOpts, arg0, arg1)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Incmodemock *IncmodemockCaller) Withdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incmodemock.contract.Call(opts, out, "withdrawed", arg0)
	return *ret0, err
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Incmodemock *IncmodemockSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Incmodemock.Contract.Withdrawed(&_Incmodemock.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Incmodemock *IncmodemockCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Incmodemock.Contract.Withdrawed(&_Incmodemock.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Incmodemock *IncmodemockTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Incmodemock *IncmodemockSession) Claim() (*types.Transaction, error) {
	return _Incmodemock.Contract.Claim(&_Incmodemock.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Incmodemock *IncmodemockTransactorSession) Claim() (*types.Transaction, error) {
	return _Incmodemock.Contract.Claim(&_Incmodemock.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Incmodemock *IncmodemockTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "deposit", incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Incmodemock *IncmodemockSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Incmodemock.Contract.Deposit(&_Incmodemock.TransactOpts, incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Incmodemock *IncmodemockTransactorSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Incmodemock.Contract.Deposit(&_Incmodemock.TransactOpts, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Incmodemock *IncmodemockTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Incmodemock *IncmodemockSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Incmodemock.Contract.DepositERC20(&_Incmodemock.TransactOpts, token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Incmodemock *IncmodemockTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Incmodemock.Contract.DepositERC20(&_Incmodemock.TransactOpts, token, amount, incognitoAddress)
}

// Execute is a paid mutator transaction binding the contract method 0xa0c0f9b8.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes32 hash, bytes signData) returns()
func (_Incmodemock *IncmodemockTransactor) Execute(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, hash [32]byte, signData []byte) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "execute", token, amount, recipientToken, exchangeAddress, callData, hash, signData)
}

// Execute is a paid mutator transaction binding the contract method 0xa0c0f9b8.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes32 hash, bytes signData) returns()
func (_Incmodemock *IncmodemockSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, hash [32]byte, signData []byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.Execute(&_Incmodemock.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, hash, signData)
}

// Execute is a paid mutator transaction binding the contract method 0xa0c0f9b8.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes32 hash, bytes signData) returns()
func (_Incmodemock *IncmodemockTransactorSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, hash [32]byte, signData []byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.Execute(&_Incmodemock.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, hash, signData)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Incmodemock *IncmodemockTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Incmodemock *IncmodemockSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Incmodemock.Contract.Extend(&_Incmodemock.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Incmodemock *IncmodemockTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Incmodemock.Contract.Extend(&_Incmodemock.TransactOpts, n)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Incmodemock *IncmodemockTransactor) Migrate(opts *bind.TransactOpts, _newVault common.Address) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "migrate", _newVault)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Incmodemock *IncmodemockSession) Migrate(_newVault common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.Migrate(&_Incmodemock.TransactOpts, _newVault)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Incmodemock *IncmodemockTransactorSession) Migrate(_newVault common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.Migrate(&_Incmodemock.TransactOpts, _newVault)
}

// MoveAssets is a paid mutator transaction binding the contract method 0x0c4f5039.
//
// Solidity: function moveAssets(address[] assets) returns()
func (_Incmodemock *IncmodemockTransactor) MoveAssets(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "moveAssets", assets)
}

// MoveAssets is a paid mutator transaction binding the contract method 0x0c4f5039.
//
// Solidity: function moveAssets(address[] assets) returns()
func (_Incmodemock *IncmodemockSession) MoveAssets(assets []common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.MoveAssets(&_Incmodemock.TransactOpts, assets)
}

// MoveAssets is a paid mutator transaction binding the contract method 0x0c4f5039.
//
// Solidity: function moveAssets(address[] assets) returns()
func (_Incmodemock *IncmodemockTransactorSession) MoveAssets(assets []common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.MoveAssets(&_Incmodemock.TransactOpts, assets)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Incmodemock *IncmodemockTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Incmodemock *IncmodemockSession) Pause() (*types.Transaction, error) {
	return _Incmodemock.Contract.Pause(&_Incmodemock.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Incmodemock *IncmodemockTransactorSession) Pause() (*types.Transaction, error) {
	return _Incmodemock.Contract.Pause(&_Incmodemock.TransactOpts)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x5654b6c8.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes32 hash) returns()
func (_Incmodemock *IncmodemockTransactor) RequestWithdraw(opts *bind.TransactOpts, incognitoAddress string, token common.Address, amount *big.Int, signData []byte, hash [32]byte) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "requestWithdraw", incognitoAddress, token, amount, signData, hash)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x5654b6c8.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes32 hash) returns()
func (_Incmodemock *IncmodemockSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, hash [32]byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.RequestWithdraw(&_Incmodemock.TransactOpts, incognitoAddress, token, amount, signData, hash)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x5654b6c8.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes32 hash) returns()
func (_Incmodemock *IncmodemockTransactorSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, hash [32]byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.RequestWithdraw(&_Incmodemock.TransactOpts, incognitoAddress, token, amount, signData, hash)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Incmodemock *IncmodemockTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Incmodemock *IncmodemockSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.Retire(&_Incmodemock.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Incmodemock *IncmodemockTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.Retire(&_Incmodemock.TransactOpts, _successor)
}

// SetAmount is a paid mutator transaction binding the contract method 0xfaf7f494.
//
// Solidity: function setAmount(address to, address token, uint256 amount) returns()
func (_Incmodemock *IncmodemockTransactor) SetAmount(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "setAmount", to, token, amount)
}

// SetAmount is a paid mutator transaction binding the contract method 0xfaf7f494.
//
// Solidity: function setAmount(address to, address token, uint256 amount) returns()
func (_Incmodemock *IncmodemockSession) SetAmount(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Incmodemock.Contract.SetAmount(&_Incmodemock.TransactOpts, to, token, amount)
}

// SetAmount is a paid mutator transaction binding the contract method 0xfaf7f494.
//
// Solidity: function setAmount(address to, address token, uint256 amount) returns()
func (_Incmodemock *IncmodemockTransactorSession) SetAmount(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Incmodemock.Contract.SetAmount(&_Incmodemock.TransactOpts, to, token, amount)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x3cb4b58a.
//
// Solidity: function submitBurnProof(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Incmodemock *IncmodemockTransactor) SubmitBurnProof(opts *bind.TransactOpts, inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "submitBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x3cb4b58a.
//
// Solidity: function submitBurnProof(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Incmodemock *IncmodemockSession) SubmitBurnProof(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.SubmitBurnProof(&_Incmodemock.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x3cb4b58a.
//
// Solidity: function submitBurnProof(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Incmodemock *IncmodemockTransactorSession) SubmitBurnProof(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.SubmitBurnProof(&_Incmodemock.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Incmodemock *IncmodemockTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Incmodemock *IncmodemockSession) Unpause() (*types.Transaction, error) {
	return _Incmodemock.Contract.Unpause(&_Incmodemock.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Incmodemock *IncmodemockTransactorSession) Unpause() (*types.Transaction, error) {
	return _Incmodemock.Contract.Unpause(&_Incmodemock.TransactOpts)
}

// UpdateIncognitoProxy is a paid mutator transaction binding the contract method 0x3a51913d.
//
// Solidity: function updateIncognitoProxy(address newIncognitoProxy) returns()
func (_Incmodemock *IncmodemockTransactor) UpdateIncognitoProxy(opts *bind.TransactOpts, newIncognitoProxy common.Address) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "updateIncognitoProxy", newIncognitoProxy)
}

// UpdateIncognitoProxy is a paid mutator transaction binding the contract method 0x3a51913d.
//
// Solidity: function updateIncognitoProxy(address newIncognitoProxy) returns()
func (_Incmodemock *IncmodemockSession) UpdateIncognitoProxy(newIncognitoProxy common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.UpdateIncognitoProxy(&_Incmodemock.TransactOpts, newIncognitoProxy)
}

// UpdateIncognitoProxy is a paid mutator transaction binding the contract method 0x3a51913d.
//
// Solidity: function updateIncognitoProxy(address newIncognitoProxy) returns()
func (_Incmodemock *IncmodemockTransactorSession) UpdateIncognitoProxy(newIncognitoProxy common.Address) (*types.Transaction, error) {
	return _Incmodemock.Contract.UpdateIncognitoProxy(&_Incmodemock.TransactOpts, newIncognitoProxy)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Incmodemock *IncmodemockTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Incmodemock.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Incmodemock *IncmodemockSession) Withdraw(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.Withdraw(&_Incmodemock.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Incmodemock *IncmodemockTransactorSession) Withdraw(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Incmodemock.Contract.Withdraw(&_Incmodemock.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// IncmodemockClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Incmodemock contract.
type IncmodemockClaimIterator struct {
	Event *IncmodemockClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockClaim represents a Claim event raised by the Incmodemock contract.
type IncmodemockClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Incmodemock *IncmodemockFilterer) FilterClaim(opts *bind.FilterOpts) (*IncmodemockClaimIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &IncmodemockClaimIterator{contract: _Incmodemock.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Incmodemock *IncmodemockFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IncmodemockClaim) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockClaim)
				if err := _Incmodemock.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Incmodemock *IncmodemockFilterer) ParseClaim(log types.Log) (*IncmodemockClaim, error) {
	event := new(IncmodemockClaim)
	if err := _Incmodemock.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Incmodemock contract.
type IncmodemockDepositIterator struct {
	Event *IncmodemockDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockDeposit represents a Deposit event raised by the Incmodemock contract.
type IncmodemockDeposit struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) FilterDeposit(opts *bind.FilterOpts) (*IncmodemockDepositIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &IncmodemockDepositIterator{contract: _Incmodemock.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IncmodemockDeposit) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockDeposit)
				if err := _Incmodemock.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) ParseDeposit(log types.Log) (*IncmodemockDeposit, error) {
	event := new(IncmodemockDeposit)
	if err := _Incmodemock.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Incmodemock contract.
type IncmodemockExtendIterator struct {
	Event *IncmodemockExtend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockExtend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockExtend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockExtend represents a Extend event raised by the Incmodemock contract.
type IncmodemockExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Incmodemock *IncmodemockFilterer) FilterExtend(opts *bind.FilterOpts) (*IncmodemockExtendIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &IncmodemockExtendIterator{contract: _Incmodemock.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Incmodemock *IncmodemockFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *IncmodemockExtend) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockExtend)
				if err := _Incmodemock.contract.UnpackLog(event, "Extend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Incmodemock *IncmodemockFilterer) ParseExtend(log types.Log) (*IncmodemockExtend, error) {
	event := new(IncmodemockExtend)
	if err := _Incmodemock.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the Incmodemock contract.
type IncmodemockMigrateIterator struct {
	Event *IncmodemockMigrate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockMigrate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockMigrate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockMigrate represents a Migrate event raised by the Incmodemock contract.
type IncmodemockMigrate struct {
	NewVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Incmodemock *IncmodemockFilterer) FilterMigrate(opts *bind.FilterOpts) (*IncmodemockMigrateIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return &IncmodemockMigrateIterator{contract: _Incmodemock.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Incmodemock *IncmodemockFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *IncmodemockMigrate) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockMigrate)
				if err := _Incmodemock.contract.UnpackLog(event, "Migrate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMigrate is a log parse operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Incmodemock *IncmodemockFilterer) ParseMigrate(log types.Log) (*IncmodemockMigrate, error) {
	event := new(IncmodemockMigrate)
	if err := _Incmodemock.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockMoveAssetsIterator is returned from FilterMoveAssets and is used to iterate over the raw logs and unpacked data for MoveAssets events raised by the Incmodemock contract.
type IncmodemockMoveAssetsIterator struct {
	Event *IncmodemockMoveAssets // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockMoveAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockMoveAssets)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockMoveAssets)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockMoveAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockMoveAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockMoveAssets represents a MoveAssets event raised by the Incmodemock contract.
type IncmodemockMoveAssets struct {
	Assets []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMoveAssets is a free log retrieval operation binding the contract event 0x492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658.
//
// Solidity: event MoveAssets(address[] assets)
func (_Incmodemock *IncmodemockFilterer) FilterMoveAssets(opts *bind.FilterOpts) (*IncmodemockMoveAssetsIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "MoveAssets")
	if err != nil {
		return nil, err
	}
	return &IncmodemockMoveAssetsIterator{contract: _Incmodemock.contract, event: "MoveAssets", logs: logs, sub: sub}, nil
}

// WatchMoveAssets is a free log subscription operation binding the contract event 0x492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658.
//
// Solidity: event MoveAssets(address[] assets)
func (_Incmodemock *IncmodemockFilterer) WatchMoveAssets(opts *bind.WatchOpts, sink chan<- *IncmodemockMoveAssets) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "MoveAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockMoveAssets)
				if err := _Incmodemock.contract.UnpackLog(event, "MoveAssets", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoveAssets is a log parse operation binding the contract event 0x492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce658.
//
// Solidity: event MoveAssets(address[] assets)
func (_Incmodemock *IncmodemockFilterer) ParseMoveAssets(log types.Log) (*IncmodemockMoveAssets, error) {
	event := new(IncmodemockMoveAssets)
	if err := _Incmodemock.contract.UnpackLog(event, "MoveAssets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Incmodemock contract.
type IncmodemockPausedIterator struct {
	Event *IncmodemockPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockPaused represents a Paused event raised by the Incmodemock contract.
type IncmodemockPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Incmodemock *IncmodemockFilterer) FilterPaused(opts *bind.FilterOpts) (*IncmodemockPausedIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncmodemockPausedIterator{contract: _Incmodemock.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Incmodemock *IncmodemockFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncmodemockPaused) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockPaused)
				if err := _Incmodemock.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Incmodemock *IncmodemockFilterer) ParsePaused(log types.Log) (*IncmodemockPaused, error) {
	event := new(IncmodemockPaused)
	if err := _Incmodemock.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Incmodemock contract.
type IncmodemockTradeIterator struct {
	Event *IncmodemockTrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockTrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockTrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockTrade represents a Trade event raised by the Incmodemock contract.
type IncmodemockTrade struct {
	IncognitoAddress string
	Token            common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x68ae6030d8e5a42e79b7dd481f18c425a779d1b6816487681116ffca8356ac49.
//
// Solidity: event Trade(string incognitoAddress, address token, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) FilterTrade(opts *bind.FilterOpts) (*IncmodemockTradeIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &IncmodemockTradeIterator{contract: _Incmodemock.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x68ae6030d8e5a42e79b7dd481f18c425a779d1b6816487681116ffca8356ac49.
//
// Solidity: event Trade(string incognitoAddress, address token, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *IncmodemockTrade) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockTrade)
				if err := _Incmodemock.contract.UnpackLog(event, "Trade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrade is a log parse operation binding the contract event 0x68ae6030d8e5a42e79b7dd481f18c425a779d1b6816487681116ffca8356ac49.
//
// Solidity: event Trade(string incognitoAddress, address token, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) ParseTrade(log types.Log) (*IncmodemockTrade, error) {
	event := new(IncmodemockTrade)
	if err := _Incmodemock.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Incmodemock contract.
type IncmodemockUnpausedIterator struct {
	Event *IncmodemockUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockUnpaused represents a Unpaused event raised by the Incmodemock contract.
type IncmodemockUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Incmodemock *IncmodemockFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncmodemockUnpausedIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncmodemockUnpausedIterator{contract: _Incmodemock.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Incmodemock *IncmodemockFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncmodemockUnpaused) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockUnpaused)
				if err := _Incmodemock.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Incmodemock *IncmodemockFilterer) ParseUnpaused(log types.Log) (*IncmodemockUnpaused, error) {
	event := new(IncmodemockUnpaused)
	if err := _Incmodemock.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockUpdateIncognitoProxyIterator is returned from FilterUpdateIncognitoProxy and is used to iterate over the raw logs and unpacked data for UpdateIncognitoProxy events raised by the Incmodemock contract.
type IncmodemockUpdateIncognitoProxyIterator struct {
	Event *IncmodemockUpdateIncognitoProxy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockUpdateIncognitoProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockUpdateIncognitoProxy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockUpdateIncognitoProxy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockUpdateIncognitoProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockUpdateIncognitoProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockUpdateIncognitoProxy represents a UpdateIncognitoProxy event raised by the Incmodemock contract.
type IncmodemockUpdateIncognitoProxy struct {
	NewIncognitoProxy common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncognitoProxy is a free log retrieval operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Incmodemock *IncmodemockFilterer) FilterUpdateIncognitoProxy(opts *bind.FilterOpts) (*IncmodemockUpdateIncognitoProxyIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return &IncmodemockUpdateIncognitoProxyIterator{contract: _Incmodemock.contract, event: "UpdateIncognitoProxy", logs: logs, sub: sub}, nil
}

// WatchUpdateIncognitoProxy is a free log subscription operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Incmodemock *IncmodemockFilterer) WatchUpdateIncognitoProxy(opts *bind.WatchOpts, sink chan<- *IncmodemockUpdateIncognitoProxy) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockUpdateIncognitoProxy)
				if err := _Incmodemock.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateIncognitoProxy is a log parse operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Incmodemock *IncmodemockFilterer) ParseUpdateIncognitoProxy(log types.Log) (*IncmodemockUpdateIncognitoProxy, error) {
	event := new(IncmodemockUpdateIncognitoProxy)
	if err := _Incmodemock.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Incmodemock contract.
type IncmodemockWithdrawIterator struct {
	Event *IncmodemockWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockWithdraw represents a Withdraw event raised by the Incmodemock contract.
type IncmodemockWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) FilterWithdraw(opts *bind.FilterOpts) (*IncmodemockWithdrawIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &IncmodemockWithdrawIterator{contract: _Incmodemock.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IncmodemockWithdraw) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockWithdraw)
				if err := _Incmodemock.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) ParseWithdraw(log types.Log) (*IncmodemockWithdraw, error) {
	event := new(IncmodemockWithdraw)
	if err := _Incmodemock.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncmodemockWithdrawRequestIterator is returned from FilterWithdrawRequest and is used to iterate over the raw logs and unpacked data for WithdrawRequest events raised by the Incmodemock contract.
type IncmodemockWithdrawRequestIterator struct {
	Event *IncmodemockWithdrawRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncmodemockWithdrawRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncmodemockWithdrawRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncmodemockWithdrawRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncmodemockWithdrawRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncmodemockWithdrawRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncmodemockWithdrawRequest represents a WithdrawRequest event raised by the Incmodemock contract.
type IncmodemockWithdrawRequest struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequest is a free log retrieval operation binding the contract event 0x392a40e1ff76f04543dc851cc0c7ad40bb9db5d3f672f27d1aa3e5c089533af5.
//
// Solidity: event WithdrawRequest(address token, string incognitoAddress, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) FilterWithdrawRequest(opts *bind.FilterOpts) (*IncmodemockWithdrawRequestIterator, error) {

	logs, sub, err := _Incmodemock.contract.FilterLogs(opts, "WithdrawRequest")
	if err != nil {
		return nil, err
	}
	return &IncmodemockWithdrawRequestIterator{contract: _Incmodemock.contract, event: "WithdrawRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequest is a free log subscription operation binding the contract event 0x392a40e1ff76f04543dc851cc0c7ad40bb9db5d3f672f27d1aa3e5c089533af5.
//
// Solidity: event WithdrawRequest(address token, string incognitoAddress, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) WatchWithdrawRequest(opts *bind.WatchOpts, sink chan<- *IncmodemockWithdrawRequest) (event.Subscription, error) {

	logs, sub, err := _Incmodemock.contract.WatchLogs(opts, "WithdrawRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncmodemockWithdrawRequest)
				if err := _Incmodemock.contract.UnpackLog(event, "WithdrawRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawRequest is a log parse operation binding the contract event 0x392a40e1ff76f04543dc851cc0c7ad40bb9db5d3f672f27d1aa3e5c089533af5.
//
// Solidity: event WithdrawRequest(address token, string incognitoAddress, uint256 amount)
func (_Incmodemock *IncmodemockFilterer) ParseWithdrawRequest(log types.Log) (*IncmodemockWithdrawRequest, error) {
	event := new(IncmodemockWithdrawRequest)
	if err := _Incmodemock.contract.UnpackLog(event, "WithdrawRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}
