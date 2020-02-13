# Requirements

- ganache-cli
- node
- npm
- Infura

# install

- Install package
```
npm i
```

- Start ganache-cli which forks and stimulates main net on your local, register an account then create a project in infura.io and replace <YOUR-PROJECT-ID> with the projectID

```
ganache-cli --fork https://mainnet.infura.io/v3/<YOUR-PROJECT-ID> --account "0x81c85096bc78372f258c804adff8cc0f16f477cc707c366dda02f4a50dd4fe3e,100000000000000000000" --debug
```

# Run

- Compile

```
truffle compile --all
```

- Deploy smart contract (run this before running others function to deploy smart contract into ganache)

```
npm run deploy
```

- Deposit ETH into Incognito Smart Contract


```
npm run deposit 1000000000000000000 // deposit 1 ETH
```

- Set withdraw request amount

```
npm run setAmount 0xdd974d5c2e2928dea5f71b9825b8b646686bd200 1000000000000000000 // set withdraw 1 KNC request 
```

- Trade

```
npm run trade 0x ETH 0x0000000000000000000000000000000000000000 1000000000000000000 DAI 0x6b175474e89094c44da98b954eedeac495271d0f myIncognitoAddress // trade 1 ETH to DAI
```

- Run predefined flow (all steps above)

```
npm run all 0x // run all steps above with 0x mode
```
