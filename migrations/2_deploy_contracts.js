const Token = artifacts.require("SimpleTokenSwapContract");

module.exports = function(deployer) {
  deployer.deploy(Token).then(function(result) {
  	console.log(result);
  });
};