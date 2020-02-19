package incognitomode

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/incognitomode/incmode"
	"github.com/incognitochain/incognitomode/kbntrade"
	"github.com/incognitochain/incognitomode/zrxtrade"
	s "github.com/stretchr/testify/suite"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
)

type SampleData struct {
	tradeType string
	srcName string
	srcTokenAddress common.Address
	destName string
	destTokenAddress common.Address
	srcQuantity *big.Int
}

const EMPTY_ADDRESS = "0x0000000000000000000000000000000000000000"


var (
	VERIFIER = common.HexToAddress("0x0858434298202ce0d76dbE20Ef5DA035CDEFc664")
	SIG = common.Hex2Bytes("53550cb7de64582b075cb387ebb3cc6391f0e98f63236d869a628b2ca1541e4e1f3d2a6d88a088f48a9e5d4d14eba6bb4c2bbbbe62e6d95958d51c97daf193f500")
	HASH = common.Hex2Bytes("6921b72d23cc590aba33512a55b1c3c84da1c03e6d8a588b9a1975a5470dbe13")
	DAI_ADDRESS = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	SAI_ADDRESS = common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359")
	ETH_ADDRESS = common.HexToAddress(EMPTY_ADDRESS)
	oneEth = big.NewInt(1000000000000000000)
	aHundredEth = oneEth.Mul(oneEth, big.NewInt(100))
	samples = []SampleData{
		{
			"Kyber",
			"ETH",
			ETH_ADDRESS,
			"DAI",
			DAI_ADDRESS,
			oneEth,
		},
		{
			"Kyber",
			"DAI",
			DAI_ADDRESS,
			"ETH",
			ETH_ADDRESS,
			oneEth, // 1 DAI
		},
		//{
		//	"Kyber",
		//	"DAI",
		//	DAI_ADDRESS,
		//	"SAI",
		//	SAI_ADDRESS,
		//	oneEth, // 1 DAI
		//},
		{
			"0x",
			"ETH",
			ETH_ADDRESS,
			"DAI",
			DAI_ADDRESS,
			oneEth, // 1 ETH
		},
		{
			"0x",
			"DAI",
			DAI_ADDRESS,
			"ETH",
			ETH_ADDRESS,
			oneEth, // 1 DAI
		},
		{
			"0x",
			"DAI",
			DAI_ADDRESS,
			"SAI",
			SAI_ADDRESS,
			oneEth, // 1 DAI
		},
	}
)

type IntegrationTestSuite struct {
	s.Suite
	privKey *ecdsa.PrivateKey
	client *ethclient.Client
	admin, incAddr, vaultAddr, kbnTradeAddr, zrxTradeAddr common.Address
	incMode *incmode.Incmode
}

func (suite *IntegrationTestSuite) Test_A_Deploy() {
	var (
		err error
		tx *types.Transaction
		ethBalance *big.Int
	)
	suite.admin = common.HexToAddress(Admin)

	prevVault := common.Address{}
	kyberNetworkProxy := common.HexToAddress("0x818E6FECD516Ecc3849DAf6845e3EC868087B755")
	wETHAddr := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	zrxProxy := common.HexToAddress("0x95e6f48254609a6ee006f7d493c8e5fb97094cef")

	beaconComm := make([]common.Address, 0)
	bridgeComm := make([]common.Address, 0)

	if suite.privKey, suite.client, err = connect(); err != nil {
		suite.T().Fatal(err)
	}
	defer suite.client.Close()

	auth := getAuth(suite.privKey)

	// deploy incognito proxy
	if suite.incAddr, tx, _, err = incognito_proxy.DeployIncognitoProxy(auth, suite.client, suite.admin, beaconComm, bridgeComm); err != nil {
		suite.T().Fatal(err)
	}

	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	// deploy incognito mode
	if suite.vaultAddr, tx, suite.incMode, err = incmode.DeployIncmode(auth, suite.client, suite.admin, suite.incAddr, prevVault); err != nil {
		suite.T().Fatal(err)
	}

	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	// deploy kbn trade
	if suite.kbnTradeAddr, tx, _, err = kbntrade.DeployKbntrade(auth, suite.client, kyberNetworkProxy, suite.vaultAddr); err != nil {
		suite.T().Fatal(err)
	}

	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	// deploy zrx trade
	if suite.zrxTradeAddr, tx, _, err = zrxtrade.DeployZrxtrade(auth, suite.client, wETHAddr, zrxProxy, suite.vaultAddr); err != nil {
		suite.T().Fatal(err)
	}

	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	// deposit
	// set auth.Value to be 100000000000000000000 (100ETH)
	auth.Value = aHundredEth
	if tx, err = suite.incMode.Deposit(auth, EMPTY_ADDRESS); err != nil {
		suite.T().Fatal(err)
	}
	// check DestToken balance of incognito mode smart contract
	incModeRaw := incmode.IncmodeRaw{Contract: suite.incMode}
	if err = incModeRaw.Call(nil, &ethBalance, "balanceOf", common.HexToAddress(EMPTY_ADDRESS)); err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(auth.Value.String(), ethBalance.String())
}

func (suite *IntegrationTestSuite) Test_B_Trade() {
	for _, sample := range samples {
		suite.trade(sample.tradeType, sample.srcName, sample.destName, sample.srcTokenAddress, sample.destTokenAddress, sample.srcQuantity)
	}
}

func getAuth(privKey *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(10000000000)
	auth.GasLimit = 5000000
	return auth
}

func (suite *IntegrationTestSuite)quote(srcToken, destToken string, srcQty *big.Int) map[string]interface{} {
	var (
		err error
		resp *http.Response
		bodyBytes []byte
		result interface{}
	)
	url := fmt.Sprintf("https://api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v", srcToken, destToken, srcQty.String())
	if resp, err = http.Get(url); err != nil {
		suite.T().Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		suite.T().Fatal(resp.StatusCode)
	}
	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		suite.T().Fatal(err)
	}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		suite.T().Fatal(err)
	}
	return result.(map[string]interface{})
}

func (suite *IntegrationTestSuite)trade(_type string, srcName, destName string, srcToken, destToken common.Address, srcQty *big.Int) {
	var (
		tx *types.Transaction
		err error
		tradeAbi abi.ABI
		input []byte
		balance *big.Int
		forwarder, tradeAddress common.Address
	)

	if _type == "Kyber" {
		tradeAbi, err = abi.JSON(strings.NewReader(kbntrade.KbntradeABI))
		tradeAddress = suite.kbnTradeAddr
	} else {
		tradeAbi, err = abi.JSON(strings.NewReader(zrxtrade.ZrxtradeABI))
		tradeAddress = suite.zrxTradeAddr
	}
	if err != nil {
		suite.T().Fatal(err)
	}

	auth := getAuth(suite.privKey)
	// setAmount
	if tx, err = suite.incMode.SetAmount(auth, VERIFIER, srcToken, srcQty); err != nil {
		suite.T().Fatal(err)
	}
	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	if _type == "Kyber" {
		input, err = tradeAbi.Pack("trade", srcToken, srcQty, destToken)
	} else {
		// quote
		quoteData := suite.quote(srcName, destName, srcQty)
		forwarder = common.HexToAddress(quoteData["to"].(string))
		data := common.Hex2Bytes(quoteData["data"].(string)[2:])
		auth.Value, _ = big.NewInt(0).SetString(quoteData["protocolFee"].(string), 10)
		auth.GasPrice, _ = big.NewInt(0).SetString(quoteData["gasPrice"].(string), 10)
		input, err = tradeAbi.Pack("trade", srcToken, srcQty, destToken, data, forwarder)
	}
	if err != nil {
		suite.T().Fatal(err)
	}
	// trigger execute function from incognitoMode
	if tx, err = suite.incMode.Execute(auth, srcToken, srcQty, destToken, tradeAddress, input, bytes32(HASH), SIG); err != nil {
		println(fmt.Sprintf("execute failed: type=%v src=%v dest=%v quantity=%v", _type, srcName, destName, srcQty.String()))
		suite.T().Fatal(err)
	}
	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}
	// check DestToken balance of incognito mode smart contract
	incModeRaw := incmode.IncmodeRaw{Contract: suite.incMode}
	if err = incModeRaw.Call(nil, &balance, "balanceOf", destToken); err != nil {
		suite.T().Fatal(err)
	}
	println(balance.String())
}

func bytes32(data []byte) (result [32]byte) {
	for i:=0; i < 32; i++ {
		result[i] = data[i]
	}
	return result
}

func TestIntegrationTestSuite(t *testing.T) {
	s.Run(t, new(IntegrationTestSuite))
}



