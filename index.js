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
let utils = require("ethereumjs-util");

dotenv.config();

let account = web3.eth.accounts.privateKeyToAccount(process.env.PRIVATEKEY);
const VERIFIER = "0x0858434298202ce0d76dbE20Ef5DA035CDEFc664";
const SIG = Buffer.from("53550cb7de64582b075cb387ebb3cc6391f0e98f63236d869a628b2ca1541e4e1f3d2a6d88a088f48a9e5d4d14eba6bb4c2bbbbe62e6d95958d51c97daf193f500", "hex");
const HASH = Buffer.from("6921b72d23cc590aba33512a55b1c3c84da1c03e6d8a588b9a1975a5470dbe13", "hex");
const EMPTY_ADDRESS = "0x0000000000000000000000000000000000000000";
const KNC_ADDRESS = "0xdd974d5c2e2928dea5f71b9825b8b646686bd200";
const SAI_ADDRESS = "0x89d24a6b4ccb1b6faa2625fe562bdd9a23260359";
const ABT_ADDRESS = "0xb98d4c97425d9908e66e53a6fdf673acca0be986";
const INCOGNITO_MODE = "./IncognitoModeMock.sol";
const INCOGNITO_PROXY = "./IncognitoProxy.sol";
const KBN_TRADE = "./KBNTrade.sol";
const ZRX_TRADE = "./ZRXTrade.sol";
const EXECUTOR = "./Executor.sol";
const ERC20 = "./ERC20.sol";
const COMPOUND = "./Compound.sol";
const TRADE_FUNC = "trade";
const BORROW_FUNC = "borrow";
const PAYBACK_FUNC = "payback";
const COMP_CONTROLLER = "0x2eaa9d77ae4d8f9cdd9faacd44016e746485bddb";
const COMPOUND_ETH_ADDRESS = "0xd6801a1dffcd0a410336ef88def4320d6df1883e";
const COMPOUND_DAI_ADDRESS = "0x6d7f0754ffeb405d23c51ce938289d4835be3b14";
const DAI_ADDRESS = "0x5592EC0cfb4dbc12D3aB100b257153436a1f0FEa";

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
let Compound = resolver.require(COMPOUND);


IncognitoProxy.setProvider(options.provider());
IncognitoMode.setProvider(options.provider());
KBNTrade.setProvider(options.provider());
ZRXTrade.setProvider(options.provider());
Executor.setProvider(options.provider());
Erc20.setProvider(options.provider());
Compound.setProvider(options.provider());

async function deploy() {
	// deploy smart contract
	let rs = await IncognitoProxy.new(account.address, [], [], {from: account.address});
	console.log(`IncognitoProxy address=${rs.address} txHash=${rs.transactionHash}`);
	cached.incognitoProxy = rs.address;

	rs = await IncognitoMode.new(account.address, cached.incognitoProxy, EMPTY_ADDRESS, {from: account.address});
	console.log(`IncognitoMode address=${rs.address} txHash=${rs.transactionHash}`);
	cached.incognitoMode = rs.address;

	// rs = await KBNTrade.new(process.env.KYBERNETWORK, cached.incognitoMode, {from: account.address});
	// console.log(`KBNTrade address=${rs.address} txHash=${rs.transactionHash}`);
	// cached.KBN = rs.address;

	// rs = await ZRXTrade.new(process.env.WETH, contractAddresses.getContractAddressesForChainOrThrow(1).erc20Proxy, cached.incognitoMode, {from: account.address});
	// console.log(`ZRXTrade address=${rs.address} txHash=${rs.transactionHash}`);
	// cached.ZRX = rs.address;

	rs = await Compound.new(cached.incognitoMode, COMP_CONTROLLER, COMPOUND_ETH_ADDRESS, {from: account.address});
	console.log(`Compound address=${rs.address} txHash=${rs.transactionHash}`);
	cached.Compound = rs.address;
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
	return findFunctionAbi(abi, TRADE_FUNC);
}

function findBorrowFunctionAbi(abi) {
	return findFunctionAbi(abi, BORROW_FUNC);
}

function findPaybackFunctionAbi(abi) {
	return findFunctionAbi(abi, PAYBACK_FUNC);
}

function findFunctionAbi(abi, funcName) {
	for (let i=0; abi.length; i++) {
		if (abi[i].name === funcName) {
			return abi[i];
		}
	}
}

async function setAmount(token, amount) {
	return await IncognitoModeContract().methods.setAmount(VERIFIER, token, amount).send({from: account.address});
}

async function deposit(amount) {
	return await IncognitoModeContract().methods.deposit(EMPTY_ADDRESS).send({from: account.address, value: amount});
}

async function quote(options) {
	let rs = await fetch(`https://api.0x.org/swap/v0/quote?${qs.stringify(options)}`);
	return await rs.json();
}

async function borrowOrPay(funcName, srcToken, destCToken, destToken, srcQty, destQty) {
	let abi = findFunctionAbi(Compound.abi, funcName);
	let params = [];
	if (funcName === BORROW_FUNC) {
		params = [VERIFIER, srcToken, srcQty, destCToken, destQty];
	} else if (funcName === PAYBACK_FUNC) {
		params = [VERIFIER, destCToken];
	}
	let inputData = web3.eth.abi.encodeFunctionCall(abi, params);
	let rs;
	if (srcToken === EMPTY_ADDRESS) {
		rs = await deposit(srcQty);
		console.log(`finish deposit ${srcQty} (wei) txHash=${rs.transactionHash}`);
	} else {
		// TBD.
	}

	// setAmount
	rs = await setAmount(srcToken, srcQty);
	console.log(`finish setAmount to token=${srcToken} amount=${srcQty} tx=${rs.transactionHash}`);

	rs = await IncognitoModeContract().methods.execute(
		srcToken, 
		srcQty, 
		destToken, 
		cached.Compound, 
		Buffer.from(inputData.slice(2, inputData.length), "hex"),
		HASH,
		SIG
	).send({from: account.address});

	console.log(`finish calling execute function at tx=${rs.transactionHash}`);
	console.log(rs);

	if (destToken === EMPTY_ADDRESS) {
		return await web3.eth.getBalance(cached.incognitoMode);
	}
	return balanceOf(destToken);
}

async function trade(_type, srcTokenName, srcToken, srcQty, destTokenName, destToken) {
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
	let rs = await IncognitoModeContract().methods.execute(
		srcToken, 
		srcQty, 
		destToken, 
		exchangeAddress, 
		Buffer.from(encodedData.slice(2, encodedData.length), "hex"),
		HASH,
		SIG
	).send(options);

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
console.log(command);
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

		trade(tradeType, srcTokenName, srcToken, srcQty, destTokenName, destToken).then(function(res) {
			console.log(res);
			console.log(`finish trading type=${tradeType==="0x" ? tradeType : "KyberNetwork"} fromToken=${srcTokenName} sellAmount=${srcQty} toToken=${destTokenName} incognitoAddress=${incognitoAddress} tx=${res.transactionHash}`);
		}); break;
	case "sigToAddress":
		let signData = Buffer.from(process.argv[3], "hex");
		let hash = Buffer.from(process.argv[4], "hex");

		// console.log(`signDataLength=${signData.length}`);
		// let r = signData.subarray(0, 32);
		// let s = signData.subarray(32, 64);
		// let lastBytes = signData.readInt8(64);
		// let v = lastBytes + 27;
		// // console.log(`r=${r.toString("hex")} s=${s.toString("hex")} v=${v}`);
		// let pub = utils.ecrecover(hash, v, "0x" + r.toString("hex"), "0x" + s.toString("hex"));
		// let recoveredAddress = "0x" + utils.pubToAddress(pub).toString('hex');
		// console.log(recoveredAddress);

		let sigToAddress = async function() {
			await deploy();
			return await IncognitoModeContract().methods.sigToAddress(signData, hash).call();
		}

		sigToAddress().then(function(rs) {
			console.log(rs);
		});
		break;
	case "compound":
		console.log(arg);
		let func = async function() {
			await deploy();
			let balance = await borrowOrPay(BORROW_FUNC, EMPTY_ADDRESS, COMPOUND_DAI_ADDRESS, DAI_ADDRESS, "100000000000000000", "10000000000000000000");
			console.log(`balance of DAI after borrow is ${balance}`);

			balance = await borrowOrPay(PAYBACK_FUNC, DAI_ADDRESS, EMPTY_ADDRESS, EMPTY_ADDRESS, "10000000000000000000", "100000000000000000");
			console.log(`balance of ETH after payback is ${balance}`);
		}
		func().then(function() {
			console.log("done");
		});
		break;
	case "compile":
		let file = process.argv[3];
		console.log(`compiling filePath=${file}`);
		let contract = require(file);
		fs.writeFile("abi.json", JSON.stringify(contract.abi), "utf8", function() {
            console.log("Saved abi.json");
        });
        fs.writeFile("bc", contract.bytecode, "utf8", function() {
            console.log("Saved bc");
        });
        break;
	default:
		let mode = process.argv[2]; // 0x or KBN
		let flow = async function() {
			// deploy
			await deploy();

			console.log("\n\n================== TRADE 1 ETH TO DAI ==================\n\n");

			// setAmount 1 ETH
			let rs = await setAmount(EMPTY_ADDRESS, "10000000000000000000");
			console.log(`setAmount = ${10000000000000000000} for ETH with tx=${rs.transactionHash}`);

			// deposit 10 ETH for later user
			rs = await deposit("10000000000000000000");
			console.log(`deposit ${10000000000000000000} to IncognitoModeContract tx=${rs.transactionHash}`);
			console.log(`balance of IncognitoMode=${await web3.eth.getBalance(cached.incognitoMode)}`);

			// trade 1 ETH to DAI
			rs = await trade(mode, "ETH", "0x0000000000000000000000000000000000000000", "1000000000000000000", "DAI", DAI_ADDRESS);
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
			rs = await trade(mode, "DAI", DAI_ADDRESS, "1000000000000000000", "ETH", EMPTY_ADDRESS);
			console.log(`trade DAI to ETH amount=${1000000000000000000} tx=${rs.transactionHash}`);

			// print all events returned by above rs.
			for (let k in rs.events) {
				console.log(rs.events[k].returnValues);
			}

			console.log("\n\n================== TRADE 10 DAI TO KNC ==================\n\n");

			// setAmount 10 DAI
			rs = await setAmount(DAI_ADDRESS, "10000000000000000000");
			console.log(`setAmount = ${10000000000000000000} for DAI with tx=${rs.transactionHash}`);

			// trade 10 DAI to KNC or ABT - KBN works mostly with ABT while 0x does not support this token.
			if (mode === "0x") {
				rs = await trade(mode, "DAI", DAI_ADDRESS, "10000000000000000000", "KNC", KNC_ADDRESS);
				console.log(`trade DAI to KNC amount=${10000000000000000000} tx=${rs.transactionHash}`);
			} else {
				rs = await trade(mode, "DAI", DAI_ADDRESS, "10000000000000000000", "ABT", ABT_ADDRESS);
				console.log(`trade DAI to ABT amount=${10000000000000000000} tx=${rs.transactionHash}`);
			}

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
