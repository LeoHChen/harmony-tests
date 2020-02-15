#!/usr/bin/env bash

echo "Installing Harmony Tx Tests"
rm -rf harmony-tests && curl -LOs http://tools.harmony.one.s3.amazonaws.com/release/linux-x86_64/harmony-tests/harmony-tests && chmod u+x harmony-tests
[ ! -d testcases ] && curl -LOs http://tools.harmony.one.s3.amazonaws.com/release/linux-x86_64/harmony-tests/testcases.tar.gz && tar -xzf testcases.tar.gz && rm -rf testcases.tar.gz
[ ! -f config.yml ] && curl -LOs https://raw.githubusercontent.com/SebastianJ/harmony-tests/master/config.yml
mkdir -p keys/testnet
echo "Harmony Tx Tests have now been downloaded!"
echo ""
echo "Make sure to either add keystore files to keys/testnet or create the file keys/testnet/private_keys.txt and add testnet private keys to it"
echo "The tests can be run on multiple networks - testnet is just used as an example. To run the tests on other networks, e.g. localnet, use --network localnet"
echo ""
echo "When you've added keyfiles or private keys, invoke the tests using ./harmony-tests"
echo ""
echo "To see all available configuration options, run ./harmony-tests --help or check out the configuration specified in config.yml"
