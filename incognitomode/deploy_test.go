package incognitomode

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/incognitomode/incmode"
	"github.com/incognitochain/incognitomode/incognito_proxy"
	"github.com/incognitochain/incognitomode/kbntrade"
)

func TestDeployProxyAndVault(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	admin := common.HexToAddress(Admin)
	fmt.Println("Admin address:", admin.Hex())

	// Genesis committee
	// cmtee := getCommitteeHardcoded()
	beaconComm, bridgeComm, err := getCommittee("http://127.0.0.1:9338/")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(10000000000)
	auth.GasLimit = 4000000
	// incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, admin, cmtee.beacons, cmtee.bridges)
	incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, admin, beaconComm, bridgeComm)
	if err != nil {
		t.Fatal(err)
	}
	// incAddr := common.HexToAddress(IncognitoProxyAddress)
	fmt.Println("deployed incognito_proxy")
	fmt.Printf("addr: %s\n", incAddr.Hex())

	// Wait until tx is confirmed
	if err := wait(client, tx.Hash()); err != nil {
		t.Fatal(err)
	}

	// Deploy vault
	prevVault := common.Address{}
	vaultAddr, _, _, err := incmode.DeployIncmode(auth, client, admin, incAddr, prevVault)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddr.Hex())

	// Deploy kbntrade
	kbnAddr := common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200")
	kbnTradeAddr, _, _, err := kbntrade.DeployKbntrade(auth, client, kbnAddr, vaultAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed kbntrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
}

func wait(client *ethclient.Client, tx common.Hash) error {
	ctx := context.Background()
	for range time.Tick(10 * time.Second) {
		_, err := client.TransactionReceipt(ctx, tx)
		if err == nil {
			break
		} else if err == ethereum.NotFound {
			continue
		} else {
			return err
		}
	}
	return nil
}

func connect() (*ecdsa.PrivateKey, *ethclient.Client, error) {
	// privKeyHex := os.Getenv("PRIVKEY")
	privKeyHex := "98CEA7DDDF43C62B7AC37C064E9DCDF6134390B9B06B83254A12A9FA4F38D3D0"
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, nil, err
	}

	return privKey, client, nil
}
