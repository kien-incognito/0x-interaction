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

function toBN(n) {
    let bn = new web3.utils.BN(n);
    return bn
}

// 'options' is acting like args which are passed to cmd
config = Config.detect(options, null);
options.resolver = new Resolver(config);


// resolver is acting like artifacts in migrations
let resolver = new ResolverIntercept(options.resolver);
let SimpleToken = resolver.require("./SimpleTokenSwapContract.sol");
SimpleToken.setProvider(options.provider());

function deploy() {
	SimpleToken.new({from: account.address}).then(function(result) {
    	console.log(`address=${result.address} txHash=${result.transactionHash}`);
    	cached.simpleToken = result.address;
    	// update cached
    	console.log("rewrite cached.json");
        fs.writeFile("cached.json", JSON.stringify(cached), "utf8", function() {
            console.log("Saved cached.json");
        });
    }).catch(function(err) {
    	console.log(err);
    	return;
    });	
}

function quote(options, cb) {
	fetch(`https://api.0x.org/swap/v0/quote?${qs.stringify(options)}`).then(function(res) {
		res.json().then(function(data) {
			cb(data);
		})
	});
}

// triggers trigger0x function which use forwarder to run data.
function trigger0x(forwarder, data, value, gasPrice, cb) {
	if (cached.simpleToken !== undefined) {
		let contract = new web3.eth.Contract(SimpleToken.abi, cached.simpleToken);
		contract.methods.trigger0x(Buffer.from(data.slice(2, data.length), "hex"), forwarder).send({
			from: account.address,
			value: value,
			gasPrice: gasPrice,
			gas: 2000000
		}).on("transactionHash", function(hash) {
			console.log(`triggering trigger0x function - hash=${hash}`);
		}).then(function(result) {
			cb();
		}).catch(function(err) {
			console.log(err);
		})
	}
}

// returns balance of given tokenAddresses.
function balanceOf(tokenAddresses, cb) {
	for (let i=0; i < tokenAddresses.length; i++) {
		let contract = new web3.eth.Contract(SimpleToken.abi, cached.simpleToken);
		if (tokenAddresses[i] === "") {
			// get eth balance
			web3.eth.getBalance(cached.simpleToken).then(function(result) {
				console.log(`balanceOf ${cached.simpleToken} in ETH is ${result}`);
				if (cb !== undefined && i == tokenAddresses.length - 1) {
					cb();
				}
			})
		} else {
			contract.methods.balanceOf(tokenAddresses[i]).call().then(function(data) {
				console.log(`balanceOf ${cached.simpleToken} in ${tokenAddresses[i]} is ${data}`);
				if (cb !== undefined && i == tokenAddresses.length - 1) {
					cb();
				}
			});
		}
	}
}

// implement swapping eth to erc20 token
function eth2Token(buyToken, buyTokenAddress, sellAmount) {
	let options = {
		sellToken: "ETH",
		buyToken: buyToken,
		sellAmount: sellAmount
	}
	balanceOf([buyTokenAddress], function() {
		quote(options, function(res) {
			console.log(`returned quote = ${JSON.stringify(res)}`);
			trigger0x(res.to, res.data, res.value, res.gasPrice, function() {
				balanceOf([buyTokenAddress], undefined);
			});
		});
	});
}

// gets allowed selling amount.
function allowance(tokenAddress, cb) {
	let contracts = contractAddresses.getContractAddressesForChainOrThrow(1);
	let contract = new web3.eth.Contract(SimpleToken.abi, cached.simpleToken);
	contract.methods.allowance("0x6b175474e89094c44da98b954eedeac495271d0f", contracts.erc20Proxy).call().then(function(res) {
		cb(res);
	});
}

// implements swapping between 2 tokens
function token2Token(sellToken, sellTokenAddress, buyToken, buyTokenAddress, sellAmount) {
	let opts = {
		sellToken: sellToken,
		buyToken: buyToken,
		sellAmount: sellAmount
	}
	exchange(opts, sellTokenAddress, buyTokenAddress);
}

// implements swapping token to eth
function token2Eth(sellToken, sellTokenAddress, sellAmount) {
	let opts = {
		sellToken: sellToken,
		buyToken: "ETH",
		sellAmount: sellAmount
	};
	exchange(opts, sellTokenAddress, process.env.WETH, function() {
		// get balance of WETH and withdraw to ETH
		let contract = new web3.eth.Contract(SimpleToken.abi, cached.simpleToken);
		contract.methods.balanceOf(process.env.WETH).call().then(function(data) {
			if (data > 0) {
				// call withdraw
				contract.methods.withdrawWrapETH(process.env.WETH, data).send({from: account.address}).then(function(res) {
					web3.eth.getBalance(cached.simpleToken).then(function(rs) {
						console.log(`ETH's balance of ${cached.simpleToken} is ${rs}`);
					});
				})
			}
		});
	});
}

// common function used for token2Eth and token2Token
function exchange(opts, sellTokenAddress, buyTokenAddress, cb) {
	let contracts = contractAddresses.getContractAddressesForChainOrThrow(1);
	let contract = new web3.eth.Contract(SimpleToken.abi, cached.simpleToken);
	let tokens = [sellTokenAddress, buyTokenAddress];
	let triggerFunc = function() {
		balanceOf(tokens, function() {
			quote(opts, function(r) {
				trigger0x(r.to, r.data, r.value, r.gasPrice, function() {
					balanceOf(tokens, function() {
						// check balance of account
						if (cb !== undefined) {
							cb();
						}
					});
				})
			})
		})
	}

	// check allowance
	allowance(sellTokenAddress, function(amount) {
		console.log(`allowance amount is ${amount}`);
		if (amount < opts.sellAmount) {
			console.log(`approving ${opts.sellAmount}`);
			contract.methods.approve(sellTokenAddress, contracts.erc20Proxy, new BigNumber(opts.sellAmount)).send({
				from: account.address,
				gas: 2000000
			})
			.on("transachtionHash", function(hash) {
				console.log(`approve's hash: ${hash}`);
			})
			.then(function(res) {
				console.log(res);
				triggerFunc();
			});
		} else {
			triggerFunc();
		}
	});
}

const command = process.argv[2];
switch (command) {
	case "deploy:simple_token_swap":
		deploy(); break;
	case "ethToToken":
		eth2Token("DAI", "0x6b175474e89094c44da98b954eedeac495271d0f", 10000000000000); break;
	case "tokenToEth":
		token2Eth("DAI", "0x6b175474e89094c44da98b954eedeac495271d0f", 100000000000000); break;
	case "tokenToToken":
		token2Token("DAI", "0x6b175474e89094c44da98b954eedeac495271d0f", "ZRX", "0xe41d2489571d322189246dafa5ebde1f4699f498", 1000000000000000); break;
	default:
		web3.eth.getBalance(account.address).then(function(rs) {
			console.log(rs);
		});
}


