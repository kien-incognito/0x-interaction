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
var IncmodemockBin = "0x60806040523480156200001157600080fd5b50604051620050bf380380620050bf833981810160405262000037919081019062000185565b82806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160146101000a81548160ff0219169083151502179055506301e1338042016002819055505081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505062000229565b6000815190506200017f816200020f565b92915050565b6000806000606084860312156200019b57600080fd5b6000620001ab868287016200016e565b9350506020620001be868287016200016e565b9250506040620001d1868287016200016e565b9150509250925092565b6000620001e882620001ef565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200021a81620001db565b81146200022657600080fd5b50565b614e8680620002396000396000f3fe6080604052600436106101f95760003560e01c80637e16e6e11161010d578063ce5494bb116100a0578063f75b98ce1161006f578063f75b98ce1461071b578063f851a44014610758578063f8b3974514610783578063fa84702e146107ac578063faf7f494146107d7576101f9565b8063ce5494bb1461063b578063cf54aaa014610664578063dca40d9e146106a1578063e4bd7074146106de576101f9565b80639714378c116100dc5780639714378c146105b15780639e6371ba146105da578063a0c0f9b814610603578063a26e11861461061f576101f9565b80637e16e6e1146105035780638456cb591461054457806388aaf0c81461055b5780638a98453814610586576101f9565b806358bc83371161019057806365b5a00f1161015f57806365b5a00f146103f65780636ff968c31461043357806370a082311461045e578063749c5f861461049b57806379599f96146104d8576101f9565b806358bc8337146103475780635a67cb87146103725780635c975abb1461038e5780636304541c146103b9576101f9565b80633f4ba83a116101cc5780633f4ba83a146102b35780633fec6b40146102ca5780634e71d92d146103075780635654b6c81461031e576101f9565b80630c4f5039146101fb5780631ea1940e146102245780633a51913d146102615780633cb4b58a1461028a575b005b34801561020757600080fd5b50610222600480360361021d9190810190613b7e565b610800565b005b34801561023057600080fd5b5061024b60048036036102469190810190613be8565b610bc7565b6040516102589190614636565b60405180910390f35b34801561026d57600080fd5b5061028860048036036102839190810190613930565b610be7565b005b34801561029657600080fd5b506102b160048036036102ac9190810190613c52565b610d7a565b005b3480156102bf57600080fd5b506102c8611086565b005b3480156102d657600080fd5b506102f160048036036102ec9190810190613dd7565b6111b8565b6040516102fe91906144c5565b60405180910390f35b34801561031357600080fd5b5061031c611246565b005b34801561032a57600080fd5b5061034560048036036103409190810190613e6c565b6113d6565b005b34801561035357600080fd5b5061035c611647565b60405161036991906144c5565b60405180910390f35b61038c60048036036103879190810190613b17565b61164c565b005b34801561039a57600080fd5b506103a361189b565b6040516103b09190614636565b60405180910390f35b3480156103c557600080fd5b506103e060048036036103db9190810190613930565b6118ae565b6040516103ed919061487c565b60405180910390f35b34801561040257600080fd5b5061041d600480360361041891908101906139be565b6118c6565b60405161042a919061487c565b60405180910390f35b34801561043f57600080fd5b506104486118eb565b60405161045591906144c5565b60405180910390f35b34801561046a57600080fd5b5061048560048036036104809190810190613930565b611911565b604051610492919061487c565b60405180910390f35b3480156104a757600080fd5b506104c260048036036104bd9190810190613be8565b6119e0565b6040516104cf9190614636565b60405180910390f35b3480156104e457600080fd5b506104ed611b24565b6040516104fa919061487c565b60405180910390f35b34801561050f57600080fd5b5061052a60048036036105259190810190613c11565b611b2a565b60405161053b9594939291906148b2565b60405180910390f35b34801561055057600080fd5b50610559611ba6565b005b34801561056757600080fd5b50610570611d1c565b60405161057d91906144fb565b60405180910390f35b34801561059257600080fd5b5061059b611d42565b6040516105a89190614786565b60405180910390f35b3480156105bd57600080fd5b506105d860048036036105d39190810190613f13565b611d68565b005b3480156105e657600080fd5b5061060160048036036105fc9190810190613930565b611ec9565b005b61061d60048036036106189190810190613a49565b611fe0565b005b61063960048036036106349190810190613e2b565b6123fd565b005b34801561064757600080fd5b50610662600480360361065d9190810190613959565b6124a5565b005b34801561067057600080fd5b5061068b60048036036106869190810190613930565b612638565b6040516106989190614897565b60405180910390f35b3480156106ad57600080fd5b506106c860048036036106c39190810190613be8565b6126fc565b6040516106d59190614636565b60405180910390f35b3480156106ea57600080fd5b5061070560048036036107009190810190613be8565b61271c565b6040516107129190614636565b60405180910390f35b34801561072757600080fd5b50610742600480360361073d91908101906139be565b612860565b60405161074f919061487c565b60405180910390f35b34801561076457600080fd5b5061076d6128e7565b60405161077a91906144c5565b60405180910390f35b34801561078f57600080fd5b506107aa60048036036107a59190810190613c52565b61290c565b005b3480156107b857600080fd5b506107c1612c9d565b6040516107ce91906147a1565b60405180910390f35b3480156107e357600080fd5b506107fe60048036036107f991908101906139fa565b612cc3565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108869061485c565b60405180910390fd5b600160149054906101000a900460ff166108de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d5906147bc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561093a57600080fd5b60008090505b8151811015610b8c57600073ffffffffffffffffffffffffffffffffffffffff1682828151811061096d57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614156109ff57600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156109f9573d6000803e3d6000fd5b50610b7f565b6000828281518110610a0d57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610a4d91906144e0565b60206040518083038186803b158015610a6557600080fd5b505afa158015610a79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610a9d9190810190613f3c565b90506000811115610b7d57828281518110610ab457fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610b1892919061454d565b602060405180830381600087803b158015610b3257600080fd5b505af1158015610b46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b6a9190810190613bbf565b50610b73612d9f565b610b7c57600080fd5b5b505b8080600101915050610940565b507f492fc8b292f2a2a9b328a366b83745f30c024056d12aa118a15966d26a8ce65881604051610bbc9190614614565b60405180910390a150565b60046020528060005260406000206000915054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6d9061485c565b60405180910390fd5b600160149054906101000a900460ff16610cc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbc906147bc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610cff57600080fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c344681604051610d6f91906144c5565b60405180910390a150565b600160149054906101000a900460ff1615610dca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc19061483c565b60405180910390fd5b6000806000806000610ddb8f611b2a565b9450945094509450945060618560ff16148015610dfb575060018460ff16145b610e0457600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610e8c57600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548101471015610e8757600080fd5b610f8c565b6000610e9784612638565b905060098160ff161115610eb45760098160ff1603600a0a820291505b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482018473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610f2f91906144e0565b60206040518083038186803b158015610f4757600080fd5b505afa158015610f5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f7f9190810190613f3c565b1015610f8a57600080fd5b505b610f9e8f8f8f8f8f8f8f8f8f8f612ddd565b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555080600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505050505050505050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161110c9061485c565b60405180910390fd5b600160149054906101000a900460ff16611164576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115b906147bc565b60405180910390fd5b6000600160146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516111ae91906144e0565b60405180910390a1565b6000806000806020860151915060408601519250601b866040815181106111db57fe5b602001015160f81c60f81b60f81c0190506001858284866040516000815260200160405260405161120f9493929190614741565b6020604051602081039080840390855afa158015611231573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600254421061128a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611281906147dc565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461131a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113119061481c565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516113cc91906144c5565b60405180910390a1565b6113df8161271c565b156113e957600080fd5b60006113f583836111b8565b905083600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561148057600080fd5b6000849050600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156114cf57633b9aca0085816114c757fe5b049050611500565b60006114da87612638565b905060098160ff1611156114fe5760098160ff1603600a0a86816114fa57fe5b0491505b505b60016004600085815260200190815260200160002060006101000a81548160ff02191690831515021790555084600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555084600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e868883604051611636939291906145ad565b60405180910390a150505050505050565b600081565b600160149054906101000a900460ff161561169c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116939061483c565b60405180910390fd5b600083905060006116ac85612638565b905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016116e991906144e0565b60206040518083038186803b15801561170157600080fd5b505afa158015611715573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506117399190810190613f3c565b9050600085905060098360ff1611156117775760098360ff1603600a0a818161175e57fe5b04905060098360ff1603600a0a828161177357fe5b0491505b670de0b6b3a764000081111580156117975750670de0b6b3a76400008211155b80156117ad5750670de0b6b3a764000082820111155b6117b657600080fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330896040518463ffffffff1660e01b81526004016117f393929190614516565b602060405180830381600087803b15801561180d57600080fd5b505af1158015611821573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506118459190810190613bbf565b5061184e612d9f565b61185757600080fd5b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e87868360405161188a939291906145ad565b60405180910390a150505050505050565b600160149054906101000a900460ff1681565b60066020528060005260406000206000915090505481565b6005602052816000526040600020602052806000526040600020600091509150505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561194f574790506119db565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161198891906144e0565b60206040518083038186803b1580156119a057600080fd5b505afa1580156119b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506119d89190810190613f3c565b90505b919050565b60006003600083815260200190815260200160002060009054906101000a900460ff1615611a115760019050611b1f565b600073ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611a715760009050611b1f565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663749c5f86836040518263ffffffff1660e01b8152600401611acc9190614726565b60206040518083038186803b158015611ae457600080fd5b505afa158015611af8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611b1c9190810190613bbf565b90505b919050565b60025481565b60008060008060008086600081518110611b4057fe5b602001015160f81c60f81b60f81c9050600087600181518110611b5f57fe5b602001015160f81c60f81b60f81c9050600080600060228b0151925060428b0151915060628b01519050848484848499509950995099509950505050505091939590929450565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c2c9061485c565b60405180910390fd5b600160149054906101000a900460ff1615611c85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c7c9061483c565b60405180910390fd5b6002544210611cc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cc0906147dc565b60405180910390fd5b60018060146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051611d1291906144e0565b60405180910390a1565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611df7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dee9061485c565b60405180910390fd5b6002544210611e3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e32906147dc565b60405180910390fd5b61016e8110611e7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e76906147fc565b60405180910390fd5b620151808102600254016002819055507f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e881604051611ebe919061487c565b60405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611f58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f4f9061485c565b60405180910390fd5b6002544210611f9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f93906147dc565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611fe98261271c565b15611ff357600080fd5b6000611fff82846111b8565b905086600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561208a57600080fd5b8573ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614156120c357600080fd5b6000349050600073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161415612107578781019050612233565b878973ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161214191906144e0565b60206040518083038186803b15801561215957600080fd5b505afa15801561216d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506121919190810190613f3c565b101561219c57600080fd5b8873ffffffffffffffffffffffffffffffffffffffff1663a9059cbb878a6040518363ffffffff1660e01b81526004016121d79291906145eb565b602060405180830381600087803b1580156121f157600080fd5b505af1158015612205573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506122299190810190613bbf565b61223257600080fd5b5b60006122418883888a613177565b905088600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555088600660008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600660008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050505050505050505050565b600160149054906101000a900460ff161561244d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124449061483c565b60405180910390fd5b6b033b2e3c9fd0803ce800000047111561246657600080fd5b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e6000823460405161249a939291906145ad565b60405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161252b9061485c565b60405180910390fd5b600160149054906101000a900460ff16612583576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161257a906147bc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156125bd57600080fd5b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a8160405161262d91906144e0565b60405180910390a150565b60008082905060008090508173ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561268957600080fd5b505afa15801561269d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506126c19190810190613f3c565b503d600081146126dc57602081146126e557600091506126f1565b600091506126f1565b60206000803e60005191505b508092505050919050565b60036020528060005260406000206000915054906101000a900460ff1681565b60006004600083815260200190815260200160002060009054906101000a900460ff161561274d576001905061285b565b600073ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156127ad576000905061285b565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4bd7074836040518263ffffffff1660e01b81526004016128089190614726565b60206040518083038186803b15801561282057600080fd5b505afa158015612834573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506128589190810190613bbf565b90505b919050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff161561295c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129539061483c565b60405180910390fd5b600080600080600061296d8f611b2a565b9450945094509450945060488560ff1614801561298d575060018460ff16145b61299657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415612a1e57600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548101471015612a1957600080fd5b612b1e565b6000612a2984612638565b905060098160ff161115612a465760098160ff1603600a0a820291505b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482018473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401612ac191906144e0565b60206040518083038186803b158015612ad957600080fd5b505afa158015612aed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612b119190810190613f3c565b1015612b1c57600080fd5b505b612b308f8f8f8f8f8f8f8f8f8f612ddd565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415612bb1578173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015612bab573d6000803e3d6000fd5b50612c51565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401612bec92919061454d565b602060405180830381600087803b158015612c0657600080fd5b505af1158015612c1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612c3e9190810190613bbf565b50612c47612d9f565b612c5057600080fd5b5b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb838383604051612c8493929190614576565b60405180910390a1505050505050505050505050505050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550505050565b600080600090503d60008114612dbc5760208114612dc557612dd1565b60019150612dd1565b60206000803e60005191505b50600081141591505090565b60008a80519060200120905060008b8b600060028110612df957fe5b6020020151604051602001612e0f92919061449d565b60405160208183030381529060405280519060200120905060008c8c600160028110612e3757fe5b6020020151604051602001612e4d92919061449d565b604051602081830303815290604052805190602001209050612e6e836119e0565b15612e7857600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f65d21166001848f600060028110612ec657fe5b60200201518f600060028110612ed857fe5b60200201518f600060028110612eea57fe5b60200201518f600060028110612efc57fe5b60200201518f600060028110612f0e57fe5b60200201518f600060028110612f2057fe5b60200201518f600060028110612f3257fe5b60200201518f600060028110612f4457fe5b60200201518f600060028110612f5657fe5b60200201516040518c63ffffffff1660e01b8152600401612f819b9a99989796959493929190614651565b60206040518083038186803b158015612f9957600080fd5b505afa158015612fad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612fd19190810190613bbf565b612fda57600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f65d21166000838f60016002811061302857fe5b60200201518f60016002811061303a57fe5b60200201518f60016002811061304c57fe5b60200201518f60016002811061305e57fe5b60200201518f60016002811061307057fe5b60200201518f60016002811061308257fe5b60200201518f60016002811061309457fe5b60200201518f6001600281106130a657fe5b60200201518f6001600281106130b857fe5b60200201516040518c63ffffffff1660e01b81526004016130e39b9a99989796959493929190614651565b60206040518083038186803b1580156130fb57600080fd5b505afa15801561310f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506131339190810190613bbf565b61313c57600080fd5b60016003600085815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050505050505050505050565b60008061318386611911565b9050600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156131c05734810390505b844710156131cd57600080fd5b600060608473ffffffffffffffffffffffffffffffffffffffff1687876040516131f79190614486565b60006040518083038185875af1925050503d8060008114613234576040519150601f19603f3d011682016040523d82523d6000602084013e613239565b606091505b50915091508161324857600080fd5b6000808280602001905161325f9190810190613982565b915091508973ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480156132a7575080856132a48c611911565b03145b6132b057600080fd5b8095505050505050949350505050565b6000813590506132cf81614db9565b92915050565b6000813590506132e481614dd0565b92915050565b6000815190506132f981614dd0565b92915050565b600082601f83011261331057600080fd5b813561332361331e82614932565b614905565b9150818183526020840193506020810190508385602084028201111561334857600080fd5b60005b83811015613378578161335e88826132c0565b84526020840193506020830192505060018101905061334b565b5050505092915050565b600082601f83011261339357600080fd5b60026133a66133a18261495a565b614905565b9150818360005b838110156133dd57813586016133c38882613516565b8452602084019350602083019250506001810190506133ad565b5050505092915050565b600082601f8301126133f857600080fd5b600261340b6134068261497c565b614905565b9150818360005b838110156134425781358601613428888261360d565b845260208401935060208301925050600181019050613412565b5050505092915050565b600082601f83011261345d57600080fd5b600261347061346b8261499e565b614905565b9150818360005b838110156134a7578135860161348d8882613704565b845260208401935060208301925050600181019050613477565b5050505092915050565b600082601f8301126134c257600080fd5b60026134d56134d0826149c0565b614905565b9150818360005b8381101561350c57813586016134f28882613787565b8452602084019350602083019250506001810190506134dc565b5050505092915050565b600082601f83011261352757600080fd5b813561353a613535826149e2565b614905565b9150818183526020840193506020810190508385602084028201111561355f57600080fd5b60005b8381101561358f5781613575888261380a565b845260208401935060208301925050600181019050613562565b5050505092915050565b600082601f8301126135aa57600080fd5b60026135bd6135b882614a0a565b614905565b915081838560208402820111156135d357600080fd5b60005b8381101561360357816135e98882613834565b8452602084019350602083019250506001810190506135d6565b5050505092915050565b600082601f83011261361e57600080fd5b813561363161362c82614a2c565b614905565b9150818183526020840193506020810190508385602084028201111561365657600080fd5b60005b83811015613686578161366c8882613834565b845260208401935060208301925050600181019050613659565b5050505092915050565b600082601f8301126136a157600080fd5b60026136b46136af82614a54565b614905565b915081838560208402820111156136ca57600080fd5b60005b838110156136fa57816136e088826138f1565b8452602084019350602083019250506001810190506136cd565b5050505092915050565b600082601f83011261371557600080fd5b813561372861372382614a76565b614905565b9150818183526020840193506020810190508385602084028201111561374d57600080fd5b60005b8381101561377d578161376388826138f1565b845260208401935060208301925050600181019050613750565b5050505092915050565b600082601f83011261379857600080fd5b81356137ab6137a682614a9e565b614905565b915081818352602084019350602081019050838560208402820111156137d057600080fd5b60005b8381101561380057816137e6888261391b565b8452602084019350602083019250506001810190506137d3565b5050505092915050565b60008135905061381981614de7565b92915050565b60008151905061382e81614de7565b92915050565b60008135905061384381614dfe565b92915050565b600082601f83011261385a57600080fd5b813561386d61386882614ac6565b614905565b9150808252602083016020830185838301111561388957600080fd5b613894838284614d5c565b50505092915050565b600082601f8301126138ae57600080fd5b81356138c16138bc82614af2565b614905565b915080825260208301602083018583830111156138dd57600080fd5b6138e8838284614d5c565b50505092915050565b60008135905061390081614e15565b92915050565b60008151905061391581614e15565b92915050565b60008135905061392a81614e2c565b92915050565b60006020828403121561394257600080fd5b6000613950848285016132c0565b91505092915050565b60006020828403121561396b57600080fd5b6000613979848285016132d5565b91505092915050565b6000806040838503121561399557600080fd5b60006139a3858286016132ea565b92505060206139b485828601613906565b9150509250929050565b600080604083850312156139d157600080fd5b60006139df858286016132c0565b92505060206139f0858286016132c0565b9150509250929050565b600080600060608486031215613a0f57600080fd5b6000613a1d868287016132c0565b9350506020613a2e868287016132c0565b9250506040613a3f868287016138f1565b9150509250925092565b600080600080600080600060e0888a031215613a6457600080fd5b6000613a728a828b016132c0565b9750506020613a838a828b016138f1565b9650506040613a948a828b016132c0565b9550506060613aa58a828b016132c0565b945050608088013567ffffffffffffffff811115613ac257600080fd5b613ace8a828b01613849565b93505060a0613adf8a828b01613834565b92505060c088013567ffffffffffffffff811115613afc57600080fd5b613b088a828b01613849565b91505092959891949750929550565b600080600060608486031215613b2c57600080fd5b6000613b3a868287016132c0565b9350506020613b4b868287016138f1565b925050604084013567ffffffffffffffff811115613b6857600080fd5b613b748682870161389d565b9150509250925092565b600060208284031215613b9057600080fd5b600082013567ffffffffffffffff811115613baa57600080fd5b613bb6848285016132ff565b91505092915050565b600060208284031215613bd157600080fd5b6000613bdf8482850161381f565b91505092915050565b600060208284031215613bfa57600080fd5b6000613c0884828501613834565b91505092915050565b600060208284031215613c2357600080fd5b600082013567ffffffffffffffff811115613c3d57600080fd5b613c4984828501613849565b91505092915050565b6000806000806000806000806000806101a08b8d031215613c7257600080fd5b60008b013567ffffffffffffffff811115613c8c57600080fd5b613c988d828e01613849565b9a50506020613ca98d828e01613690565b99505060608b013567ffffffffffffffff811115613cc657600080fd5b613cd28d828e016133e7565b98505060808b013567ffffffffffffffff811115613cef57600080fd5b613cfb8d828e01613382565b97505060a0613d0c8d828e01613599565b96505060e0613d1d8d828e01613599565b9550506101208b013567ffffffffffffffff811115613d3b57600080fd5b613d478d828e0161344c565b9450506101408b013567ffffffffffffffff811115613d6557600080fd5b613d718d828e016134b1565b9350506101608b013567ffffffffffffffff811115613d8f57600080fd5b613d9b8d828e016133e7565b9250506101808b013567ffffffffffffffff811115613db957600080fd5b613dc58d828e016133e7565b9150509295989b9194979a5092959850565b60008060408385031215613dea57600080fd5b600083013567ffffffffffffffff811115613e0457600080fd5b613e1085828601613849565b9250506020613e2185828601613834565b9150509250929050565b600060208284031215613e3d57600080fd5b600082013567ffffffffffffffff811115613e5757600080fd5b613e638482850161389d565b91505092915050565b600080600080600060a08688031215613e8457600080fd5b600086013567ffffffffffffffff811115613e9e57600080fd5b613eaa8882890161389d565b9550506020613ebb888289016132c0565b9450506040613ecc888289016138f1565b935050606086013567ffffffffffffffff811115613ee957600080fd5b613ef588828901613849565b9250506080613f0688828901613834565b9150509295509295909350565b600060208284031215613f2557600080fd5b6000613f33848285016138f1565b91505092915050565b600060208284031215613f4e57600080fd5b6000613f5c84828501613906565b91505092915050565b6000613f718383613ffb565b60208301905092915050565b6000613f8983836141ef565b60208301905092915050565b6000613fa1838361420d565b60208301905092915050565b6000613fb98383614433565b60208301905092915050565b6000613fd18383614468565b60208301905092915050565b613fe681614cde565b82525050565b613ff581614c7f565b82525050565b61400481614c6d565b82525050565b61401381614c6d565b82525050565b600061402482614b6e565b61402e8185614bfc565b935061403983614b1e565b8060005b8381101561406a5781516140518882613f65565b975061405c83614bbb565b92505060018101905061403d565b5085935050505092915050565b600061408282614b79565b61408c8185614c0d565b935061409783614b2e565b8060005b838110156140c85781516140af8882613f7d565b97506140ba83614bc8565b92505060018101905061409b565b5085935050505092915050565b60006140e082614b84565b6140ea8185614c1e565b93506140f583614b3e565b8060005b8381101561412657815161410d8882613f95565b975061411883614bd5565b9250506001810190506140f9565b5085935050505092915050565b600061413e82614b8f565b6141488185614c2f565b935061415383614b4e565b8060005b8381101561418457815161416b8882613fad565b975061417683614be2565b925050600181019050614157565b5085935050505092915050565b600061419c82614b9a565b6141a68185614c40565b93506141b183614b5e565b8060005b838110156141e25781516141c98882613fc5565b97506141d483614bef565b9250506001810190506141b5565b5085935050505092915050565b6141f881614c91565b82525050565b61420781614c91565b82525050565b61421681614c9d565b82525050565b61422581614c9d565b82525050565b600061423682614ba5565b6142408185614c51565b9350614250818560208601614d6b565b80840191505092915050565b61426581614cf0565b82525050565b61427481614d14565b82525050565b600061428582614bb0565b61428f8185614c5c565b935061429f818560208601614d6b565b6142a881614da8565b840191505092915050565b60006142c0601483614c5c565b91507f6e6f7420706175736564207269676874206e6f770000000000000000000000006000830152602082019050919050565b6000614300600783614c5c565b91507f65787069726564000000000000000000000000000000000000000000000000006000830152602082019050919050565b6000614340601a83614c5c565b91507f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e670000000000006000830152602082019050919050565b6000614380600c83614c5c565b91507f756e617574686f72697a656400000000000000000000000000000000000000006000830152602082019050919050565b60006143c0601083614c5c565b91507f706175736564207269676874206e6f77000000000000000000000000000000006000830152602082019050919050565b6000614400600983614c5c565b91507f6e6f742061646d696e00000000000000000000000000000000000000000000006000830152602082019050919050565b61443c81614cc7565b82525050565b61444b81614cc7565b82525050565b61446261445d82614cc7565b614d9e565b82525050565b61447181614cd1565b82525050565b61448081614cd1565b82525050565b6000614492828461422b565b915081905092915050565b60006144a9828561422b565b91506144b58284614451565b6020820191508190509392505050565b60006020820190506144da600083018461400a565b92915050565b60006020820190506144f56000830184613fdd565b92915050565b60006020820190506145106000830184613fec565b92915050565b600060608201905061452b6000830186613fdd565b6145386020830185613fdd565b6145456040830184614442565b949350505050565b60006040820190506145626000830185613fdd565b61456f6020830184614442565b9392505050565b600060608201905061458b600083018661400a565b6145986020830185613fdd565b6145a56040830184614442565b949350505050565b60006060820190506145c2600083018661400a565b81810360208301526145d4818561427a565b90506145e36040830184614442565b949350505050565b6000604082019050614600600083018561400a565b61460d6020830184614442565b9392505050565b6000602082019050818103600083015261462e8184614019565b905092915050565b600060208201905061464b60008301846141fe565b92915050565b600061016082019050614667600083018e6141fe565b614674602083018d61421c565b614681604083018c614442565b8181036060830152614693818b6140d5565b905081810360808301526146a7818a614077565b90506146b660a083018961421c565b6146c360c083018861421c565b81810360e08301526146d58187614133565b90508181036101008301526146ea8186614191565b90508181036101208301526146ff81856140d5565b905081810361014083015261471481846140d5565b90509c9b505050505050505050505050565b600060208201905061473b600083018461421c565b92915050565b6000608082019050614756600083018761421c565b6147636020830186614477565b614770604083018561421c565b61477d606083018461421c565b95945050505050565b600060208201905061479b600083018461425c565b92915050565b60006020820190506147b6600083018461426b565b92915050565b600060208201905081810360008301526147d5816142b3565b9050919050565b600060208201905081810360008301526147f5816142f3565b9050919050565b6000602082019050818103600083015261481581614333565b9050919050565b6000602082019050818103600083015261483581614373565b9050919050565b60006020820190508181036000830152614855816143b3565b9050919050565b60006020820190508181036000830152614875816143f3565b9050919050565b60006020820190506148916000830184614442565b92915050565b60006020820190506148ac6000830184614477565b92915050565b600060a0820190506148c76000830188614477565b6148d46020830187614477565b6148e1604083018661400a565b6148ee6060830185613fec565b6148fb6080830184614442565b9695505050505050565b6000604051905081810181811067ffffffffffffffff8211171561492857600080fd5b8060405250919050565b600067ffffffffffffffff82111561494957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561497157600080fd5b602082029050919050565b600067ffffffffffffffff82111561499357600080fd5b602082029050919050565b600067ffffffffffffffff8211156149b557600080fd5b602082029050919050565b600067ffffffffffffffff8211156149d757600080fd5b602082029050919050565b600067ffffffffffffffff8211156149f957600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115614a2157600080fd5b602082029050919050565b600067ffffffffffffffff821115614a4357600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115614a6b57600080fd5b602082029050919050565b600067ffffffffffffffff821115614a8d57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115614ab557600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115614add57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115614b0957600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000614c7882614ca7565b9050919050565b6000614c8a82614ca7565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000614ce982614d38565b9050919050565b6000614cfb82614d02565b9050919050565b6000614d0d82614ca7565b9050919050565b6000614d1f82614d26565b9050919050565b6000614d3182614ca7565b9050919050565b6000614d4382614d4a565b9050919050565b6000614d5582614ca7565b9050919050565b82818337600083830152505050565b60005b83811015614d89578082015181840152602081019050614d6e565b83811115614d98576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b614dc281614c6d565b8114614dcd57600080fd5b50565b614dd981614c7f565b8114614de457600080fd5b50565b614df081614c91565b8114614dfb57600080fd5b50565b614e0781614c9d565b8114614e1257600080fd5b50565b614e1e81614cc7565b8114614e2957600080fd5b50565b614e3581614cd1565b8114614e4057600080fd5b5056fea365627a7a72315820d0010695a98ffcfb5c8e1ab2935d6671fd16b4163a731a782e55fb3788fa0ce26c6578706572696d656e74616cf564736f6c63430005100040"

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
