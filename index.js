let Config = require("@truffle/config");
let Resolver = require("@truffle/resolver");
let ResolverIntercept = require("@truffle/migrate/resolverintercept");
let Web3 = require("web3");
let fs = require("fs");
let qs = require("qs");
let fetch = require("node-fetch");
let cached = require("./cached.json");
let erc20Contract = require("@0x/contracts-erc20").ERC20TokenContract;
let dotenv = require('dotenv');
let web3 = new Web3('http://localhost:8545');
let BigNumber = require("@0x/utils").BigNumber;
let contractAddresses = require("@0x/contract-addresses");

dotenv.config();

let account = web3.eth.accounts.privateKeyToAccount(process.env.PRIVATEKEY);
const EMPTY_ADDRESS = "0x0000000000000000000000000000000000000000";
const DAI_ADDRESS = "0x6b175474e89094c44da98b954eedeac495271d0f";
const KNC_ADDRESS = "0xdd974d5c2e2928dea5f71b9825b8b646686bd200";
const SAI_ADDRESS = "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359";
const ABT_ADDRESS = "0xb98d4c97425d9908e66e53a6fdf673acca0be986";
const INCOGNITO_MODE = "./IncognitoMode.sol";
const INCOGNITO_PROXY = "./IncognitoProxy.sol";
const KBN_TRADE = "./KBNTrade.sol";
const ZRX_TRADE = "./ZRXTrade.sol";
const EXECUTOR = "./Executor.sol";
const ERC20 = "./ERC20.sol";
const TRADE_FUNC = "trade";

options = {
    working_dir: './',
    network_id: 1,
    resolver: null,
    provider: function() {
        return new Web3.providers.HttpProvider('http://localhost:8545');
    },
    logger: {
      log: function(msg) {
        logger.log("  " + msg);
      }
    },
    basePath: "./"
};

// 'options' is acting like args which are passed to cmd
config = Config.detect(options, null);
options.resolver = new Resolver(config);


// resolver is acting like artifacts in migrations
let resolver = new ResolverIntercept(options.resolver);
let IncognitoProxy = resolver.require(INCOGNITO_PROXY);
let IncognitoMode = resolver.require(INCOGNITO_MODE);
let KBNTrade = resolver.require(KBN_TRADE);
let ZRXTrade = resolver.require(ZRX_TRADE);
let Executor = resolver.require(EXECUTOR);
let Erc20 = resolver.require(ERC20);

IncognitoProxy.setProvider(options.provider());
IncognitoMode.setProvider(options.provider());
KBNTrade.setProvider(options.provider());
ZRXTrade.setProvider(options.provider());
Executor.setProvider(options.provider());
Erc20.setProvider(options.provider());

async function deploy() {
	// deploy smart contract
	let rs = await Executor.new({from: account.address});
	console.log(`Executor address=${rs.address} txHash=${rs.transactionHash}`);
	cached.executor = rs.address;

	rs = await IncognitoProxy.new(account.address, [], [], {from: account.address});
	console.log(`IncognitoProxy address=${rs.address} txHash=${rs.transactionHash}`);
	cached.incognitoProxy = rs.address;

	rs = await IncognitoMode.new(account.address, cached.incognitoProxy, EMPTY_ADDRESS, cached.executor, {from: account.address});
	console.log(`IncognitoMode address=${rs.address} txHash=${rs.transactionHash}`);
	cached.incognitoMode = rs.address;

	rs = await KBNTrade.new(process.env.KYBERNETWORK, cached.incognitoMode, {from: account.address});
	console.log(`KBNTrade address=${rs.address} txHash=${rs.transactionHash}`);
	cached.KBN = rs.address;

	rs = await ZRXTrade.new(process.env.WETH, contractAddresses.getContractAddressesForChainOrThrow(1).erc20Proxy, cached.incognitoMode, {from: account.address});
	console.log(`ZRXTrade address=${rs.address} txHash=${rs.transactionHash}`);
	cached.ZRX = rs.address;
}

function IncognitoModeContract() {
	return new web3.eth.Contract(IncognitoMode.abi, cached.incognitoMode);
}

function ERC20Contract(token) {
	return new web3.eth.Contract(Erc20.abi, token);
}

function ExecutorContract() {
	return new web3.eth.Contract(Executor.abi, cached.executor);
}

function findTradeFunctionAbi(abi) {
	for (let i=0; abi.length; i++) {
		if (abi[i].name === TRADE_FUNC) {
			return abi[i];
		}
	}
}

async function setAmount(token, amount) {
	return await IncognitoModeContract().methods.setAmount(token, amount).send({from: account.address});
}

async function deposit(amount) {
	return await IncognitoModeContract().methods.deposit(EMPTY_ADDRESS).send({from: account.address, value: amount});
}

async function quote(options) {
	let rs = await fetch(`https://api.0x.org/swap/v0/quote?${qs.stringify(options)}`);
	return await rs.json();
}

async function trade(_type, srcTokenName, srcToken, srcQty, destTokenName, destToken, incognitoAddress) {
	let encodedData = "";
	let exchangeAddress = "";
	let options = {
		from: account.address,
		gas: 1000000
	}

	if (destToken === EMPTY_ADDRESS) {
		console.log(`eth balance of incognitoMode before trading is ${await web3.eth.getBalance(cached.incognitoMode)}`);
	} else {
		console.log(`balance of token ${destTokenName} before trading of incognitoMode is ${await balanceOf(destToken)}`);
	}

	if (_type === "0x") {
		let rs = await trade0x(srcTokenName, srcToken, srcQty, destTokenName, destToken);
		encodedData = rs[0];
		options.gasPrice = rs[1];
		options.value = rs[2];
		exchangeAddress = cached.ZRX;
	} else {
		encodedData = await tradeKBN(srcToken, srcQty, destToken);
		exchangeAddress = cached.KBN;
	}

	console.log(`encodedData=${encodedData} address=${exchangeAddress}`);
	let rs = await IncognitoModeContract().methods.trade(incognitoAddress, srcToken, srcQty, destToken, exchangeAddress, Buffer.from(encodedData.slice(2, encodedData.length), "hex")).send(options);

	if (destToken === EMPTY_ADDRESS) {
		console.log(`eth balance of incognitoMode after trading is ${await web3.eth.getBalance(cached.incognitoMode)}`);
	} else {
		console.log(`balance of token ${destTokenName} after trading of incognitoMode is ${await balanceOf(destToken)}`);
	}
	return rs;
}

async function balanceOf(token) {
	return await ERC20Contract(token).methods.balanceOf(cached.incognitoMode).call();
}

async function trade0x(srcTokenName, srcToken, srcQty, destTokenName, destToken) {
	let options = {
		sellToken: srcTokenName,
		buyToken: destTokenName,
		sellAmount: srcQty
	}
	let quoteData = await quote(options);

	// find abi of trade function and encode it.
	let abi = findTradeFunctionAbi(ZRXTrade.abi);
	let params = [srcToken, srcQty, destToken, Buffer.from(quoteData.data.slice(2, quoteData.data.length), "hex"), quoteData.to];
	return [web3.eth.abi.encodeFunctionCall(abi, params), quoteData.gasPrice, quoteData.protocolFee];
}

async function tradeKBN(srcToken, srcQty, destToken) {
	let abi = findTradeFunctionAbi(KBNTrade.abi);
	let params = [srcToken, new BigNumber(srcQty), destToken];
	return web3.eth.abi.encodeFunctionCall(abi, params);
}

const command = process.argv[2];
let amount = 0;
switch (command) {
	case "deploy":
		deploy().then(function() {
			console.log("rewrite cached.json");
	        fs.writeFile("cached.json", JSON.stringify(cached), "utf8", function() {
	            console.log("Saved cached.json");
	        });
		}).catch(function(e) {
			console.log(e);
		}); 
		break;
	case "setAmount":
		let token = process.argv[3];
		amount = process.argv[4];

		setAmount(token, amount).then(function(rs) {
			console.log(`amount=${amount} has been set for token=${token} at tx=${rs.transactionHash}`);
		}); break;
	case "deposit":
		amount = process.argv[3];
		deposit(amount).then(function(rs) {
			let func = async function() {
				let ethBalance = await web3.eth.getBalance(cached.incognitoMode);
				console.log(`finish depositting amount=${amount} wei into incognitoMode smc at tx=${rs.transactionHash} currentBalance=${ethBalance} wei`);
			};
			func();
		}); break;
	case "trade":
		let tradeType = process.argv[3]
		let srcTokenName = process.argv[4];
		let srcToken = process.argv[5];
		let srcQty = process.argv[6];
		let destTokenName = process.argv[7];
		let destToken = process.argv[8];
		let incognitoAddress = process.argv[9];

		trade(tradeType, srcTokenName, srcToken, srcQty, destTokenName, destToken, incognitoAddress).then(function(res) {
			console.log(res);
			console.log(`finish trading type=${tradeType==="0x" ? tradeType : "KyberNetwork"} fromToken=${srcTokenName} sellAmount=${srcQty} toToken=${destTokenName} incognitoAddress=${incognitoAddress} tx=${res.transactionHash}`);
		}); break;
	default:
		let mode = process.argv[2]; // 0x or KBN
		console.log(`running all with mode=${mode}`);
		let flow = async function() {
			// deploy
			await deploy();

			console.log("\n\n================== TRADE 1 ETH TO DAI ==================\n\n");

			// setAmount 10 ETH for later use
			let rs = await setAmount(EMPTY_ADDRESS, "10000000000000000000");
			console.log(`setAmount = ${10000000000000000000} for ETH with tx=${rs.transactionHash}`);

			// deposit 1 ETH
			rs = await deposit("1000000000000000000");
			console.log(`deposit ${1000000000000000000} to IncognitoModeContract tx=${rs.transactionHash}`);

			// trade 1 ETH to DAI
			rs = await trade(mode, "ETH", "0x0000000000000000000000000000000000000000", "1000000000000000000", "DAI", DAI_ADDRESS, "myIncognitoAddress");
			console.log(`trade ETH to DAI amount=${1000000000000000000} tx=${rs.transactionHash}`);

			// print all events returned by above rs.
			for (let k in rs.events) {
				console.log(rs.events[k].returnValues);
			}
			
			console.log("\n\n================== TRADE 100 DAI TO ETH ==================\n\n");

			// setAmount 10 DAI
			rs = await setAmount(DAI_ADDRESS, "100000000000000000000");
			console.log(`setAmount = ${100000000000000000000} for DAI with tx=${rs.transactionHash}`);

			// trade 10 DAI to ETH
			rs = await trade(mode, "DAI", DAI_ADDRESS, "1000000000000000000", "ETH", EMPTY_ADDRESS, "myIncognitoAddress");
			console.log(`trade DAI to ETH amount=${1000000000000000000} tx=${rs.transactionHash}`);

			// print all events returned by above rs.
			for (let k in rs.events) {
				console.log(rs.events[k].returnValues);
			}

			console.log("\n\n================== TRADE 10 DAI TO KNC ==================\n\n");

			// setAmount 10 DAI
			rs = await setAmount(DAI_ADDRESS, "10000000000000000000");
			console.log(`setAmount = ${10000000000000000000} for DAI with tx=${rs.transactionHash}`);

			// trade 10 DAI to KNC
			rs = await trade(mode, "DAI", DAI_ADDRESS, "10000000000000000000", "ABT", ABT_ADDRESS, "myIncognitoAddress");
			console.log(`trade DAI to KNC amount=${10000000000000000000} tx=${rs.transactionHash}`);

			// print all events returned by above rs.
			for (let k in rs.events) {
				console.log(rs.events[k].returnValues);
			}

		}
		flow().then(function() {
			console.log("finish");
		}).catch(function(e) {
			console.log(e);
		});
}
