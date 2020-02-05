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

- Start ganache-cli which fork main net: register/create an account and a project in infura.io and replace <YOUR-PROJECT-ID> with the projectID

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

- Run Eth to Token demo (note: run this before do others swapping functions)


```
npm run ethToToken
```

- Run token to eth demo

```
npm run tokenToEth
```

- Run token to token demo

```
npm run tokenToToken
```

# Known Issues

- Eth's amount in smart contract is not updated after running tokenToEth.
