package incognitomode
//
//import (
//	"io/ioutil"
//	"errors"
//	"context"
//	"crypto/ecdsa"
//	"fmt"
//	"net/http"
//	"strings"
//	"encoding/json"
//
//	"github.com/incognitochain/incognitomode/incmode"
//	"github.com/incognitochain/incognitomode/incognito_proxy"
//
//	"math/big"
//	"testing"
//	"time"
//
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/ethereum/go-ethereum/params"
//	"github.com/incognitochain/incognitomode/kbntrade"
//	"github.com/incognitochain/incognitomode/zrxtrade"
//	"github.com/ethereum/go-ethereum/rlp"
//	"github.com/incognitochain/incognitomode/common/base58"
//	"github.com/incognitochain/incognitomode/consensus/signatureschemes/bridgesig"
//	"golang.org/x/crypto/sha3"
//	"github.com/ethereum/go-ethereum/accounts/abi"
//)
//
//func TestDeployProxyAndVault(t *testing.T) {
//	privKey, client, err := connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer client.Close()
//
//	admin := common.HexToAddress(Admin)
//	fmt.Println("Admin address:", admin.Hex())
//
//	// Genesis committee
//	// cmtee := getCommitteeHardcoded()
//	beaconComm, bridgeComm, err := getCommittee("http://127.0.0.1:9338/")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// Deploy incognito_proxy
//	auth := bind.NewKeyedTransactor(privKey)
//	auth.Value = big.NewInt(0)
//	// auth.GasPrice = big.NewInt(10000000000)
//	// auth.GasLimit = 4000000
//	// incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, admin, cmtee.beacons, cmtee.bridges)
//	incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, admin, beaconComm, bridgeComm)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// incAddr := common.HexToAddress(IncognitoProxyAddress)
//	fmt.Println("deployed incognito_proxy")
//	fmt.Printf("addr: %s\n", incAddr.Hex())
//
//	// Wait until tx is confirmed
//	if err := wait(client, tx.Hash()); err != nil {
//		t.Fatal(err)
//	}
//
//	// Deploy vault
//	prevVault := common.Address{}
//	vaultAddr, incModeTx, _, err := incmode.DeployIncmode(auth, client, admin, incAddr, prevVault)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("deployed vault")
//	fmt.Printf("addr: %s\n", vaultAddr.Hex())
//
//	// Wait until tx is confirmed
//	if err := wait(client, incModeTx.Hash()); err != nil {
//		t.Fatal(err)
//	}
//
//	// Deploy kbntrade
//	kbnAddr := common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200")
//	kbnTradeAddr, _, _, err := kbntrade.DeployKbntrade(auth, client, kbnAddr, vaultAddr)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("deployed kbntrade")
//	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
//
//	// Deploy 0xTrade
//	wETHAddr := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
//	zeroProxyAddr := common.HexToAddress("0x95e6f48254609a6ee006f7d493c8e5fb97094cef")
//	zrxTradeAddr, _, _, err := zrxtrade.DeployZrxtrade(auth, client, wETHAddr, zeroProxyAddr, vaultAddr)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("deployed zrxTrade")
//	fmt.Printf("addr: %s\n", zrxTradeAddr.Hex())
//}
//
//func TestDeposit(t *testing.T) {
//	privKey, client, err := connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer client.Close()
//
//	// Get contract instance
//	VaultAddress := "0xf94533EfD4a5B70DD0374f8f4457023268036EfF"
//	vaultAddr := common.HexToAddress(VaultAddress)
//	c, err := incmode.NewIncmode(vaultAddr, client)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Deposit
//	auth := bind.NewKeyedTransactor(privKey)
//	// auth.GasPrice = big.NewInt(20000000000)
//	// auth.GasPrice = big.NewInt(10000000000)
//	// auth.GasLimit = 4000000
//	auth.Value = big.NewInt(3 * params.Ether)
//	IncPaymentAddr := "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"
//	tx, err := c.Deposit(auth, IncPaymentAddr)
//	if err != nil {
//		t.Fatal(err)
//	}
//	txHash := tx.Hash()
//	fmt.Printf("deposited, txHash: %x\n", txHash[:])
//}
//
//func TestBurnForDepositToSC(t *testing.T) {
//	// Get proof
//	// proof, err := getAndDecodeBurnProof("59333c998a206e99621faf150f46588bbdfeb6279538266de893cc309e7cf4c5")
//	proof, err := getAndDecodeBurnForDepositToSCProof("6581ed0922f669eb7aeebb9dedf4733e0f1008a752f18ed3b71c12e514a3e56e")
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("Burn for deposit to SC proof: ", proof)
//	// return
//
//	// Connect to ETH
//	privKey, client, err := connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer client.Close()
//
//	// Get contract instance
//	VaultAddress := "0xf94533EfD4a5B70DD0374f8f4457023268036EfF"
//	vaultAddr := common.HexToAddress(VaultAddress)
//	c, err := incmode.NewIncmode(vaultAddr, client)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Burn
//	auth := bind.NewKeyedTransactor(privKey)
//	// auth.GasPrice = big.NewInt(1000000)
//	// auth.GasLimit = 4000000
//	tx, err := SubmitBurnProof(c, auth, proof)
//	if err != nil {
//		t.Fatal(err)
//	}
//	txHash := tx.Hash()
//	fmt.Printf("burned, txHash: %x\n", txHash[:])
//}
//
//func TestGetDepositedBalance(t *testing.T) {
//	_, client, err := connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer client.Close()
//
//	// Get contract instance
//	VaultAddress := "0xf94533EfD4a5B70DD0374f8f4457023268036EfF"
//	vaultAddr := common.HexToAddress(VaultAddress)
//	c, err := incmode.NewIncmode(vaultAddr, client)
//	if err != nil {
//		t.Fatal(err)
//	}
//	token := common.HexToAddress("0x0000000000000000000000000000000000000000")
//	// token := common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
//	owner := common.HexToAddress("0x8224890Cd5A537792d1B8B56c95FAb8a1A5E98B1")
//	bal, err := c.GetDepositedBalance(nil, token, owner)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Printf("deposited balance: %d\n", bal)
//}
//
//func rlpHash(x interface{}) (h common.Hash) {
//	hw := sha3.NewLegacyKeccak256()
//	rlp.Encode(hw, x)
//	hw.Sum(h[:0])
//	return h
//}
//
//func TestRequestWithdraw(t *testing.T) {
//	IncPaymentAddr := "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"
//	incPriKeyStr := "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
//	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(incPriKeyStr)
//	if err != nil {
//		fmt.Println("Decode error: ", err)
//		t.Fatal(err)
//	}
//	newPriKey, newPubKey := bridgesig.KeyGen(incPriKeyBytes)
//	pubKeyStr := crypto.PubkeyToAddress(newPubKey).Hex()
//	fmt.Println("pubKeyStr: ", pubKeyStr)
//
//	tokenID := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
//	hash := rlpHash([]interface{}{
//		tokenID,
//		big.NewInt(1234509876),
//	})
//	// data := toByte32(hash.Bytes())
//	data := hash.Bytes()
//	signBytes, _ := crypto.Sign(data, &newPriKey)
//
//	privKey, client, err := connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer client.Close()
//
//	// Get contract instance
//	VaultAddress := "0xf94533EfD4a5B70DD0374f8f4457023268036EfF"
//	vaultAddr := common.HexToAddress(VaultAddress)
//	c, err := incmode.NewIncmode(vaultAddr, client)
//	if err != nil {
//		t.Fatal(err)
//	}
//	auth := bind.NewKeyedTransactor(privKey)
//
//	token := common.HexToAddress("0x0000000000000000000000000000000000000000")
//	amount := big.NewInt(0.1 * params.Ether)
//	tx, err := c.RequestWithdraw(auth, IncPaymentAddr, token, amount, signBytes, toByte32(data))
//	if err != nil {
//		t.Fatal(err)
//	}
//	txHash := tx.Hash()
//	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
//}
//
//func TestExecuteWith0x(t *testing.T) {
//	incPriKeyStr := "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
//	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(incPriKeyStr)
//	if err != nil {
//		fmt.Println("Decode error: ", err)
//		t.Fatal(err)
//	}
//	newPriKey, newPubKey := bridgesig.KeyGen(incPriKeyBytes)
//	pubKeyStr := crypto.PubkeyToAddress(newPubKey).Hex()
//	fmt.Println("pubKeyStr: ", pubKeyStr)
//
//	tokenID := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
//	hash := rlpHash([]interface{}{
//		tokenID,
//		big.NewInt(1234509876),
//	})
//	// data := toByte32(hash.Bytes())
//	data := hash.Bytes()
//	signBytes, _ := crypto.Sign(data, &newPriKey)
//
//	// for 0x trade
//	tradeAbi, _ := abi.JSON(strings.NewReader(zrxtrade.ZrxtradeABI))
//	tradeAddress := common.HexToAddress("0x12180D9182c300873B2e57cD7FCfD61e736222d3")
//
//	privKey, client, err := connect()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer client.Close()
//	// Get contract instance
//	VaultAddress := "0xf94533EfD4a5B70DD0374f8f4457023268036EfF"
//	vaultAddr := common.HexToAddress(VaultAddress)
//	c, err := incmode.NewIncmode(vaultAddr, client)
//	if err != nil {
//		t.Fatal(err)
//	}
//	auth := bind.NewKeyedTransactor(privKey)
//
//	// quote
//	srcName := "ETH"
//	srcToken := common.HexToAddress("0x0000000000000000000000000000000000000000")
//	destName := "DAI"
//	destToken := common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
//	srcQty := big.NewInt(0.1 * params.Ether)
//
//	quoteData, _ := quote0x(srcName, destName, srcQty)
//	forwarder := common.HexToAddress(quoteData["to"].(string))
//	dt := common.Hex2Bytes(quoteData["data"].(string)[2:])
//	auth.Value, _ = big.NewInt(0).SetString(quoteData["protocolFee"].(string), 10)
//	auth.GasPrice, _ = big.NewInt(0).SetString(quoteData["gasPrice"].(string), 10)
//	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, dt, forwarder)
//
//
//	tx, err := c.Execute(
//		auth,
//		srcToken,
//		srcQty,
//		destToken,
//		tradeAddress,
//		input,
//		toByte32(data),
//		signBytes,
//	)
//	if err != nil {
//		t.Fatal(err)
//	}
//	txHash := tx.Hash()
//	fmt.Printf("0x trade executed , txHash: %x\n", txHash[:])
//}
//
//func quote0x(
//	srcToken, destToken string,
//	srcQty *big.Int,
//) (map[string]interface{}, error) {
//	var (
//		err       error
//		resp      *http.Response
//		bodyBytes []byte
//		result    interface{}
//	)
//	url := fmt.Sprintf("https://api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v", srcToken, destToken, srcQty.String())
//	if resp, err = http.Get(url); err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode != http.StatusOK {
//		return nil, errors.New("Request returns with fucking error!!!")
//	}
//	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
//		return nil, err
//	}
//	if err = json.Unmarshal(bodyBytes, &result); err != nil {
//		return nil, err
//	}
//	return result.(map[string]interface{}), nil
//}
//
//func wait(client *ethclient.Client, tx common.Hash) error {
//	ctx := context.Background()
//	for range time.Tick(10 * time.Second) {
//		_, err := client.TransactionReceipt(ctx, tx)
//		if err == nil {
//			break
//		} else if err == ethereum.NotFound {
//			continue
//		} else {
//			return err
//		}
//	}
//	return nil
//}
//
//func connect() (*ecdsa.PrivateKey, *ethclient.Client, error) {
//	// privKeyHex := os.Getenv("PRIVKEY")
//	privKeyHex := "98cea7dddf43c62b7ac37c064e9dcdf6134390b9b06b83254a12a9fa4f38d3d0"
//	privKey, err := crypto.HexToECDSA(privKeyHex)
//	if err != nil {
//		return nil, nil, err
//	}
//	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())
//
//	network := "development"
//	fmt.Printf("Connecting to network %s\n", network)
//	client, err := ethclient.Dial("http://127.0.0.1:8545")
//	if err != nil {
//		return nil, nil, err
//	}
//
//	return privKey, client, nil
//}
