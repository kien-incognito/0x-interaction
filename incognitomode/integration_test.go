package incognitomode

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/kien-incognito/0x-interaction/incognitomode/compound"
	"github.com/kien-incognito/0x-interaction/incognitomode/incmodemock"
	"github.com/kien-incognito/0x-interaction/incognitomode/kbntrade"
	"github.com/kien-incognito/0x-interaction/incognitomode/zrxtrade"
	s "github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SampleData struct {
	tradeType        string
	srcName          string
	srcTokenAddress  common.Address
	destName         string
	destTokenAddress common.Address
	srcQuantity      *big.Int
}

const EMPTY_ADDRESS = "0x0000000000000000000000000000000000000000"

var (
	VERIFIER    = common.HexToAddress("0x0858434298202ce0d76dbE20Ef5DA035CDEFc664")
	SIG         = common.Hex2Bytes("53550cb7de64582b075cb387ebb3cc6391f0e98f63236d869a628b2ca1541e4e1f3d2a6d88a088f48a9e5d4d14eba6bb4c2bbbbe62e6d95958d51c97daf193f500")
	HASH        = common.Hex2Bytes("6921b72d23cc590aba33512a55b1c3c84da1c03e6d8a588b9a1975a5470dbe13")
	MAIN_NET_DAI_ADDRESS = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	MAIN_NET_SAI_ADDRESS = common.HexToAddress("0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359")
	RINKEBY_DAI_ADDRESS = common.HexToAddress("0x5592EC0cfb4dbc12D3aB100b257153436a1f0FEa")
	RINKEBY_COMPTROLLER = common.HexToAddress("0x2eaa9d77ae4d8f9cdd9faacd44016e746485bddb")
	RINKEBY_CETHER      = common.HexToAddress("0xd6801a1dffcd0a410336ef88def4320d6df1883e")
	ETH_ADDRESS = common.HexToAddress(EMPTY_ADDRESS)

	RINKEBY_CDAI_ADDRESS = common.HexToAddress("0x6d7f0754ffeb405d23c51ce938289d4835be3b14")

	oneEth      = big.NewInt(1000000000000000000)
	aHundredEth = big.NewInt(0).Mul(oneEth, big.NewInt(100))
	samples     = []SampleData{
		{
			"Kyber",
			"ETH",
			ETH_ADDRESS,
			"DAI",
			MAIN_NET_DAI_ADDRESS,
			oneEth,
		},
		{
			"Kyber",
			"DAI",
			MAIN_NET_DAI_ADDRESS,
			"ETH",
			ETH_ADDRESS,
			oneEth, // 1 DAI
		},
		{
			"Kyber",
			"DAI",
			MAIN_NET_DAI_ADDRESS,
			"SAI",
			MAIN_NET_SAI_ADDRESS,
			oneEth, // 1 DAI
		},
		//{
		//	"0x",
		//	"ETH",
		//	ETH_ADDRESS,
		//	"DAI",
		//	DAI_ADDRESS,
		//	oneEth, // 1 ETH
		//},
		//{
		//	"0x",
		//	"DAI",
		//	DAI_ADDRESS,
		//	"ETH",
		//	ETH_ADDRESS,
		//	oneEth, // 1 DAI
		//},
		//{
		//	"0x",
		//	"DAI",
		//	DAI_ADDRESS,
		//	"SAI",
		//	SAI_ADDRESS,
		//	oneEth, // 1 DAI
		//},
	}
)

type IntegrationTestSuite struct {
	s.Suite
	privKey                                               *ecdsa.PrivateKey
	client                                                *ethclient.Client
	admin, incAddr, vaultAddr, kbnTradeAddr, zrxTradeAddr common.Address
	compoundAddr                                          common.Address
	incMode                                               *incmodemock.Incmodemock
	compound                                              *compound.Compound
}

func (suite *IntegrationTestSuite) Test_A_Deploy() {
	var (
		err        error
		tx         *types.Transaction
	)
	suite.admin = common.HexToAddress(Admin)

	prevVault := common.Address{}
	//kyberNetworkProxy := common.HexToAddress("0x818E6FECD516Ecc3849DAf6845e3EC868087B755")
	//wETHAddr := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	//zrxProxy := common.HexToAddress("0x95e6f48254609a6ee006f7d493c8e5fb97094cef")

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
	if suite.vaultAddr, tx, suite.incMode, err = incmodemock.DeployIncmodemock(auth, suite.client, suite.admin, suite.incAddr, prevVault); err != nil {
		suite.T().Fatal(err)
	}

	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	//// deploy kbn trade
	//if suite.kbnTradeAddr, tx, _, err = kbntrade.DeployKbntrade(auth, suite.client, kyberNetworkProxy, suite.vaultAddr); err != nil {
	//	suite.T().Fatal(err)
	//}
	//
	//// Wait until tx is confirmed
	//if err := wait(suite.client, tx.Hash()); err != nil {
	//	suite.T().Fatal(err)
	//}
	//
	//// deploy zrx trade
	//if suite.zrxTradeAddr, tx, _, err = zrxtrade.DeployZrxtrade(auth, suite.client, wETHAddr, zrxProxy, suite.vaultAddr); err != nil {
	//	suite.T().Fatal(err)
	//}
	//
	//// Wait until tx is confirmed
	//if err := wait(suite.client, tx.Hash()); err != nil {
	//	suite.T().Fatal(err)
	//}

	if suite.compoundAddr, tx, suite.compound, err = compound.DeployCompound(auth, suite.client, suite.vaultAddr, RINKEBY_COMPTROLLER, RINKEBY_CETHER); err != nil {
		suite.T().Fatal(err)
	}

	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}

	println(fmt.Sprintf("incognito smart contract: %v", suite.vaultAddr.Hex()))
	println(fmt.Sprintf("compound smart contract: %v", suite.compoundAddr.Hex()))

	// deposit
	// TODO(@dnkien): comment these lines if we are using testnet instead of forked version.
	// set auth.Value to be 100000000000000000000 (100ETH)
	//auth.Value = aHundredEth
	//if tx, err = suite.incMode.Deposit(auth, EMPTY_ADDRESS); err != nil {
	//	suite.T().Fatal(err)
	//}
	// check DestToken balance of incognito mode smart contract
	//incModeRaw := incmodemock.IncmodemockRaw{Contract: suite.incMode}
	//if err = incModeRaw.Call(nil, &ethBalance, "balanceOf", common.HexToAddress(EMPTY_ADDRESS)); err != nil {
	//	suite.T().Fatal(err)
	//}
	//suite.Equal(auth.Value.String(), ethBalance.String())
}

//func (suite *IntegrationTestSuite) Test_B_Trade() {
//	for _, sample := range samples {
//		suite.trade(sample.tradeType, sample.srcName, sample.destName, sample.srcTokenAddress, sample.destTokenAddress, sample.srcQuantity)
//	}
//}

func (suite *IntegrationTestSuite) Test_C_Borrow_DAI_From_ETH() {
	oneTenEth := big.NewInt(0).Quo(oneEth, big.NewInt(10))

	// mint 0.1 eth for borrowing 10 DAI (10 * 10**18)
	suite.borrowOrPayback("borrow", ETH_ADDRESS, RINKEBY_CDAI_ADDRESS, RINKEBY_DAI_ADDRESS, oneTenEth, big.NewInt(0).Mul(oneEth, big.NewInt(10)))
}

func (suite *IntegrationTestSuite) Test_D_Payback_DAI_To_Get_ETH() {
	auth := getAuth(suite.privKey)
	expectedAmount := big.NewInt(0).Mul(oneEth, big.NewInt(10))
	// check hasBorrowed
	hasBorrowed, err := suite.compound.HasBorrowed(&bind.CallOpts{From: auth.From}, VERIFIER, RINKEBY_CDAI_ADDRESS)
	suite.NoError(err)
	suite.True(hasBorrowed);

	// get repaidAmount
	borrowInfo, err := suite.compound.Borrowers(&bind.CallOpts{From: auth.From}, VERIFIER, RINKEBY_CDAI_ADDRESS)
	suite.NoError(err)
	suite.Equal(expectedAmount.Cmp(borrowInfo.BorrowedAmount), 0, borrowInfo.BorrowedAmount.String())
	suite.Equal(ETH_ADDRESS.Hex(), borrowInfo.CollateralToken.Hex())

	// send 10 DAI to get back 0.1 eth
	suite.borrowOrPayback("payback", RINKEBY_DAI_ADDRESS, RINKEBY_CDAI_ADDRESS, ETH_ADDRESS, big.NewInt(0).Mul(oneEth, big.NewInt(10)), big.NewInt(0))
}

func (suite *IntegrationTestSuite) Test_E_Invest_ETH() {
	suite.invest(ETH_ADDRESS, RINKEBY_CETHER, oneEth)
}

func (suite *IntegrationTestSuite) Test_F_Redeem_ETH() {
	auth := getAuth(suite.privKey)
	balance, err := suite.incMode.BalanceOf(&bind.CallOpts{From: auth.From}, RINKEBY_CETHER)
	suite.NoError(err)
	suite.redeem(RINKEBY_CETHER, ETH_ADDRESS, ETH_ADDRESS, balance)
}

func getAuth(privKey *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(10000000000)
	auth.GasLimit = 5000000
	return auth
}

func (suite *IntegrationTestSuite) quote(srcToken, destToken string, srcQty *big.Int) map[string]interface{} {
	var (
		err       error
		resp      *http.Response
		bodyBytes []byte
		result    interface{}
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

func (suite *IntegrationTestSuite)borrowOrPayback(funcName string, srcToken, destCToken, destToken common.Address, srcQty, destQty *big.Int) {
	var (
		tx *types.Transaction
		err error
		compoundAbi abi.ABI
		input []byte
		balance *big.Int
	)
	println(fmt.Sprintf("start %v: srcToken=%v destCToken=%v destToken=%v srcQty=%v destQty=%v",
		funcName, srcToken.Hex(), destCToken.Hex(), destToken.Hex(), srcQty.String(), destQty.String()))
	incModeRaw := incmodemock.IncmodemockRaw{Contract: suite.incMode}

	if compoundAbi, err = abi.JSON(strings.NewReader(compound.CompoundABI)); err != nil {
		suite.T().Fatal(err)
	}
	if funcName == "borrow" {
		input, err = compoundAbi.Pack(funcName, VERIFIER, srcToken, srcQty, destCToken, destQty)
	} else if funcName == "payback" {
		input, err = compoundAbi.Pack(funcName, VERIFIER, destCToken)
	} else {
		suite.T().Fatal(fmt.Errorf("invalid funcName"))
	}
	if err != nil {
		suite.T().Fatal(err)
	}
	// define tx.data
	auth := getAuth(suite.privKey)

	// setAmount to stimulate step submitBurntProof
	if tx, err = suite.incMode.SetAmount(auth, VERIFIER, srcToken, srcQty); err != nil {
		suite.T().Fatal(err)
	}
	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}
	println(fmt.Sprintf("Finish executing setAmount at tx:%v", tx.Hash().Hex()))

	// if src token is ETH, send ETH to incognito smart contract before executing
	if srcToken.Hex() == EMPTY_ADDRESS {
		println(fmt.Sprintf("depositing %v into %v", srcQty.String(), suite.incAddr.Hex()))
		auth.Value = srcQty
		// create tx that sends ether to incognito smart contract
		if tx, err = incModeRaw.Transfer(auth); err != nil {
			suite.T().Fatal(err)
		}
		// Wait until tx is confirmed
		if err := wait(suite.client, tx.Hash()); err != nil {
			suite.T().Fatal(err)
		}
		println(fmt.Sprintf(fmt.Sprintf("finish transfer %v wei to %v at tx %v", auth.Value.String(), suite.vaultAddr.Hex(), tx.Hash().Hex())))
		// reset auth
		auth.Value = big.NewInt(0)
	} else {
		// otherwise: TBD.
		// TODO: may send token to incognito smart contract before processing, or run trade test or borrow before testing this function.
		// get balance of srcToken
		// check DestToken balance of incognito mode smart contract
		if err = incModeRaw.Call(nil, &balance, "balanceOf", srcToken); err != nil {
			suite.T().Fatal(err)
		}
		suite.GreaterOrEqual(balance.Cmp(srcQty), 0, fmt.Sprintf("balance less than srcQty: %v", balance.String()))
	}

	// call execute function
	// trigger execute function from incognitoMode
	if tx, err = suite.incMode.Execute(auth, srcToken, srcQty, destToken, suite.compoundAddr, input, bytes32(HASH), SIG); err != nil {
		println(fmt.Sprintf("execute failed: src=%v dest=%v srcQty=%v destQty=%v", srcToken.Hex(), destToken.Hex(), srcQty.String(), destQty.String()))
		suite.T().Fatal(err)
	}
	// Wait until tx is confirmed
	if err := wait(suite.client, tx.Hash()); err != nil {
		suite.T().Fatal(err)
	}
	println(fmt.Sprintf("Finish executing execute at tx:%v", tx.Hash().Hex()))
	// check DestToken balance of incognito mode smart contract
	if err = incModeRaw.Call(nil, &balance, "balanceOf", destToken); err != nil {
		suite.T().Fatal(err)
	}
	println(balance.String())
}

func (suite *IntegrationTestSuite) invest(srcToken, srcCToken common.Address, investAmount *big.Int) {
	var (
		tx *types.Transaction
		err error
		compoundAbi abi.ABI
		input []byte
		balance *big.Int
		amount *big.Int
	)
	incModeRaw := incmodemock.IncmodemockRaw{Contract: suite.incMode}
	compoundAbi, err = abi.JSON(strings.NewReader(compound.CompoundABI))
	suite.NoError(err)

	if srcToken == ETH_ADDRESS {
		input, err = compoundAbi.Pack("invest", VERIFIER, srcToken, investAmount)
	} else {
		input, err = compoundAbi.Pack("invest", VERIFIER, srcCToken, investAmount)
	}
	suite.NoError(err)

	// define tx.data
	auth := getAuth(suite.privKey)

	// setAmount to stimulate step submitBurntProof
	tx, err = suite.incMode.SetAmount(auth, VERIFIER, srcToken, investAmount)
	suite.NoError(err)

	// Wait until tx is confirmed
	err = wait(suite.client, tx.Hash())
	suite.NoError(err)
	println(fmt.Sprintf("Finish executing setAmount at tx:%v", tx.Hash().Hex()))

	// send eth to incognito smart contract to stimulate smc already has eth balance
	if srcToken == ETH_ADDRESS {
		println(fmt.Sprintf("depositing %v into %v", investAmount.String(), suite.incAddr.Hex()))
		auth.Value = investAmount
		// create tx that sends ether to incognito smart contract
		tx, err = incModeRaw.Transfer(auth)
		suite.NoError(err)

		// Wait until tx is confirmed
		err = wait(suite.client, tx.Hash())
		suite.NoError(err)

		println(fmt.Sprintf(fmt.Sprintf("finish transfer %v wei to %v at tx %v", auth.Value.String(), suite.vaultAddr.Hex(), tx.Hash().Hex())))
		// reset auth
		auth.Value = big.NewInt(0)
	} else {
		// TBD.
		// should we do anything here?
	}
	// call execute function
	// trigger execute function from incognitoMode
	tx, err = suite.incMode.Execute(auth, srcToken, investAmount, srcCToken, suite.compoundAddr, input, bytes32(HASH), SIG)
	suite.NoError(err, fmt.Sprintf("execute failed: src=%v srcCToken=%v srcQty=%v", srcToken.Hex(), srcCToken.Hex(), investAmount.String()))

	// Wait until tx is confirmed
	err = wait(suite.client, tx.Hash())
	suite.NoError(err)

	println(fmt.Sprintf("Finish executing execute at tx:%v", tx.Hash().Hex()))

	// get investedAmount from field investors in Compound
	if srcToken == ETH_ADDRESS {
		amount, err = suite.compound.Investors(&bind.CallOpts{From: auth.From}, VERIFIER, ETH_ADDRESS)
	} else {
		amount, err = suite.compound.Investors(&bind.CallOpts{From: auth.From}, VERIFIER, srcCToken)
	}
	suite.NoError(err)

	// check balanceOf
	balance, err = suite.incMode.BalanceOf(&bind.CallOpts{From: auth.From}, srcCToken)
	suite.NoError(err)

	// compare minted amount in compound with balance of incognito smart contract
	suite.Equal(amount.Cmp(balance), 0, fmt.Sprintf("balance mismatched: investedAmount=%v receivedAmount=%v", amount.String(), balance.String()))

	// TODO: Note that returned compound balance is different than amount of input token
	println(fmt.Sprintf("balance of %v is %v", srcCToken.Hex(), balance.String()))
}

func (suite *IntegrationTestSuite) redeem(srcToken, srcCToken, destToken common.Address, redeemAmount *big.Int) {
	var (
		tx *types.Transaction
		err error
		compoundAbi abi.ABI
		input []byte
		balance *big.Int
	)
	compoundAbi, err = abi.JSON(strings.NewReader(compound.CompoundABI))
	suite.NoError(err)

	input, err = compoundAbi.Pack("redeem", VERIFIER, srcCToken, redeemAmount)
	suite.NoError(err)

	// define tx.data
	auth := getAuth(suite.privKey)

	// setAmount to stimulate step submitBurntProof
	tx, err = suite.incMode.SetAmount(auth, VERIFIER, srcToken, redeemAmount)
	suite.NoError(err)

	// Wait until tx is confirmed
	err = wait(suite.client, tx.Hash())
	suite.NoError(err)
	println(fmt.Sprintf("Finish executing setAmount at tx:%v", tx.Hash().Hex()))

	// call execute function
	// trigger execute function from incognitoMode
	tx, err = suite.incMode.Execute(auth, srcToken, redeemAmount, destToken, suite.compoundAddr, input, bytes32(HASH), SIG)
	suite.NoError(err, fmt.Sprintf("execute failed: src=%v srcCToken=%v destToken=%v srcQty=%v", srcToken.Hex(), srcCToken.Hex(), destToken.Hex(), redeemAmount.String()))

	// Wait until tx is confirmed
	err = wait(suite.client, tx.Hash())
	suite.NoError(err)
	println(fmt.Sprintf("Finish executing execute at tx:%v", tx.Hash().Hex()))

	// check balance of returned token's amount
	balance, err = suite.incMode.BalanceOf(&bind.CallOpts{From: auth.From}, destToken)
	suite.NoError(err)
	println(fmt.Sprintf("balance of destToken %v is %v", destToken.Hex(), balance.String()))
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
	incModeRaw := incmodemock.IncmodemockRaw{Contract: suite.incMode}
	if err = incModeRaw.Call(nil, &balance, "balanceOf", destToken); err != nil {
		suite.T().Fatal(err)
	}
	println(balance.String())
}

func bytes32(data []byte) (result [32]byte) {
	for i := 0; i < 32; i++ {
		result[i] = data[i]
	}
	return result
}

func wait(client *ethclient.Client, tx common.Hash) error {
	ctx := context.Background()
	for range time.Tick(10 * time.Second) {
		receipt, err := client.TransactionReceipt(ctx, tx)
		if err != nil {
			if err == ethereum.NotFound {
				continue
			}
			return err
		}
		if receipt.Status == types.ReceiptStatusFailed {
			return fmt.Errorf("tx %v failed", tx.Hex())
		}
		break
	}
	return nil
}

// Run ganache-cli --fork https://rinkeby.infura.io/v3/<your_API_KEY> --account "0x81c85096bc78372f258c804adff8cc0f16f477cc707c366dda02f4a50dd4fe3e,100000000000000000000000"
// before run the test.
func connect() (*ecdsa.PrivateKey, *ethclient.Client, error) {
	// privKeyHex := os.Getenv("PRIVKEY")
	privKeyHex := "81c85096bc78372f258c804adff8cc0f16f477cc707c366dda02f4a50dd4fe3e"
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "local"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial("http://localhost:8545")
	//client, err := ethclient.Dial("https://rinkeby.infura.io/v3/6bc9014389394a1fb3aa1985f323d3c7")
	if err != nil {
		return nil, nil, err
	}

	return privKey, client, nil
}

func TestIntegrationTestSuite(t *testing.T) {
	// TODO: Trade, BorrowOrPay and invest-redeem should be run separately.
	//  for unknown reasons, revert will be thrown in some of the tests while running tests end to end.
	//  research how to solve this problem.

	// TODO: investigate about interest after invest

	// TODO: write compound tests for token to token (ETH) cases.

	s.Run(t, new(IntegrationTestSuite))
}
