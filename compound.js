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
const EMPTY_ADDRESS = "0x0000000000000000000000000000000000000000";
const COMP_CONTROLLER = "0x2eaa9d77ae4d8f9cdd9faacd44016e746485bddb";
const COMPOUND_ETH_ADDRESS = "0xd6801a1dffcd0a410336ef88def4320d6df1883e";
const COMPOUND_DAI_ADDRESS = "0x6d7f0754ffeb405d23c51ce938289d4835be3b14";

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
let Compound = resolver.require("./Compound.sol");
Compound.setProvider(options.provider());


function CompoundContract() {
	return new web3.eth.Contract(Compound.abi, cached.compound);
}

async function balanceOf(ctoken) {
	return await CompoundContract().methods.balanceOf(cToken).call();
}

async function deploy() {
	let rs = await Compound.new(COMP_CONTROLLER, COMPOUND_ETH_ADDRESS, {from: account.address});
	console.log(`Compound address=${rs.address} txHash=${rs.transactionHash}`);
	cached.compound = rs.address;
}

async function borrow(borrower, collateralToken, collateralAmount, borrowedToken, borrowedAmount) {
	if (collateralToken === EMPTY_ADDRESS) {
		return await CompoundContract().methods.borrow(borrower, collateralToken, collateralAmount, borrowedToken, borrowedAmount).send({from: account.address, value: collateralAmount});
	}
	return await CompoundContract().methods.borrow(borrower, collateralToken, collateralAmount, borrowedToken, borrowedAmount).send({from: account.address});
}

async function repay(cToken) {
	return await CompoundContract().methods.payback(account.address, cToken).send({from: account.address});
}

async function flow() {
	await deploy();

	// borrow DAI from ETH
	// let res = await borrow(account.address, EMPTY_ADDRESS, "1000000000000000000", COMPOUND_DAI_ADDRESS, "100000000000000000000");
	// // print all events returned by above rs.
	// for (let k in res.events) {
	// 	console.log(res.events[k].returnValues);
	// }

	await CompoundContract().methods.mintEth().send({from: account.address, value: "1000000000000000000"});

	// // check balance of DAI
	// res = await balanceOf(COMPOUND_DAI_ADDRESS);
	// console.log(`balanceOf DAI is ${res}`);

	// // repay DAI to get back ETH
	// res = await repay(COMPOUND_DAI_ADDRESS);
	// console.log(res);
}

flow().then(function() {
	console.log("done");
}).catch(function(err) {
	console.error(err);
})


