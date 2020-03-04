# harmony-tests
harmony-tests executes a set of test cases for testing that regular and staking transactions work properly on Harmony's blockchain.

It uses the harmony-tf testing framework under the hood.

## Installation

```
mkdir -p harmony-tests && cd harmony-tests
bash <(curl -s -S -L https://raw.githubusercontent.com/SebastianJ/harmony-tests/master/scripts/install.sh)
```

## Usage
If you want to use private keys:
e.g. for testnet -> `nano keys/testnet/private_keys.txt` - and save your list of private keys to this file

If you want to use keystore files:
e.g. for testnet -> `cp -r keystore-folder keys/testnet`

Harmony TF will automatically identify keyfiles no matter what you call the folders or what the files are called - as long as they reside under `keys/testnet` (or whatever network you're using) they'll be identified.

## Usage
To run all test cases:
`./harmony-tests --network NETWORK`

Only run staking tests: `./harmony-tests --network NETWORK --test staking`
Only run staking -> create validators: `./harmony-tests --network NETWORK --test staking/validators/create`
Only run staking -> edit validators: `./harmony-tests --network NETWORK --test staking/validators/edit`
Only run staking -> delegation: `./harmony-tests --network NETWORK --test staking/delegation/delegate`
Only run staking -> undelegation: `./harmony-tests --network NETWORK --test staking/delegation/undelegate`

Only run tx tests: `./harmony-tests --network NETWORK --test transactions`
Only run tx -> cross app shard tests: `./harmony-tests --network NETWORK --test transactions/cross_app_shard`
Only run tx -> cross beacon shard tests: `./harmony-tests --network NETWORK --test transactions/cross_beacon_shard`
Only run tx -> same app shard tests: `./harmony-tests --network NETWORK --test transactions/same_app_shard`
Only run tx -> same beacon shard tests: `./harmony-tests --network NETWORK --test transactions/same_beacon_shard`

## Writing test cases
Test cases are defined as YAML files and are placed in testcases/ - see this folder for existing test cases and how to impelement test cases.

There are some requirements for these files - this README will be updated with a list of these eventually :)
