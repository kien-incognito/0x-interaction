package incognitomode

//import (
//	"math/big"
//
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/incognitochain/incognitomode/incmode"
//	"github.com/incognitochain/incognitomode/incognito_proxy"
//)
//
//func Withdraw(v *incmode.Incmode, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
//	// auth.GasPrice = big.NewInt(20000000000)
//	tx, err := v.Withdraw(
//		auth,
//		proof.Instruction,
//		proof.Heights,
//
//		proof.InstPaths,
//		proof.InstPathIsLefts,
//		proof.InstRoots,
//		proof.BlkData,
//		proof.SigIdxs,
//		proof.SigVs,
//		proof.SigRs,
//		proof.SigSs,
//	)
//	if err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
//
//func SubmitBurnProof(i *incmode.Incmode, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
//	// auth.GasPrice = big.NewInt(20000000000)
//	tx, err := i.SubmitBurnProof(
//		auth,
//		proof.Instruction,
//		proof.Heights,
//
//		proof.InstPaths,
//		proof.InstPathIsLefts,
//		proof.InstRoots,
//		proof.BlkData,
//		proof.SigIdxs,
//		proof.SigVs,
//		proof.SigRs,
//		proof.SigSs,
//	)
//	if err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
//
//func SwapBridge(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
//	auth.GasPrice = big.NewInt(20000000000)
//	tx, err := inc.SwapBridgeCommittee(
//		auth,
//		proof.Instruction,
//
//		proof.InstPaths,
//		proof.InstPathIsLefts,
//		proof.InstRoots,
//		proof.BlkData,
//		proof.SigIdxs,
//		proof.SigVs,
//		proof.SigRs,
//		proof.SigSs,
//	)
//	if err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
//
//func SwapBeacon(inc *incognito_proxy.IncognitoProxy, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
//	auth.GasPrice = big.NewInt(20000000000)
//	tx, err := inc.SwapBeaconCommittee(
//		auth,
//		proof.Instruction,
//
//		proof.InstPaths[0],
//		proof.InstPathIsLefts[0],
//		proof.InstRoots[0],
//		proof.BlkData[0],
//		proof.SigIdxs[0],
//		proof.SigVs[0],
//		proof.SigRs[0],
//		proof.SigSs[0],
//	)
//	if err != nil {
//		return nil, err
//	}
//	return tx, nil
//}
