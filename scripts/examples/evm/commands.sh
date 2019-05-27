./rtool --backend="tcp://localhost:46657" evm create --callf /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample.json --abif /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample.abi --nonce 0

./rtool --backend="tcp://localhost:46657" evm exist  --callf /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample_exist.json

./rtool --backend="tcp://localhost:46657" query nonce --address aaf40b3b7d103b01e89c7aa489ed186c5467dabf

./rtool --backend="tcp://localhost:46657" evm call --callf /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample_execute.json --abif /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample.abi --nonce 1

./rtool --backend="tcp://localhost:46657" evm read --callf /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample_read.json --abif /Users/za/workspace/src/github.com/dappledger/AnnChainvous/scripts/examples/evm/sample.abi --nonce 2
