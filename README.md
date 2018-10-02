# KaleidoProject
This is a forked geth 1.8 w/ step-by-step logging placed to trace the mining process. 

* `/go-ethereum` - modified geth with mining logging
* `/devnet` - PoA geth network containing (2) nodes and (1) bootnode. Configured using the modified geth bin 
* `miner.log` - rolling log of mining process

## Fire it up
* Make geth:
  * `cd /go-ethereum; make geth`
  * Make note of path to the newly compiled geth binary. (i.e. `/Users/Mac/Documents/kaleidotest/go-ethereum/build/bin/geth`). We'll reference this binary in future steps as `$CustomGeth`
* Initialize bootnode:
  * In `/devnet` run `/Users/Mac/Documents/kaleidotest/go-ethereum/build/bootnode -nodekey boot.key -verbosity -9 -addr :30310`
* Initialize node 1:
  * Open 2nd terminal window
  * In `/devnet` run `$CustomGeth --datadir node1/ --syncmode 'full' --port 30311 --rpc --rpcaddr 'localhost' --rpcport 8501 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --bootnodes 'enode://a0cd5bf310b8c30396c80cb82d053c522a3bb5c0e3f483f78abeddfc64b351d2401222f7b8f3425670dfca0c395c11d3b1290d8d5e6655e0fcba957c73cbd679@127.0.0.1:30310' --networkid 1515 --gasprice '1' -unlock '0xdb97df08c187fb9c7f46b34b00200eaa95e321c3' --password node1/password.txt --mine`
* Initialize node 2:
  * Open 3rd terminal window
  * In `/devnet` run `$CustomGeth --datadir node2/ --syncmode 'full' --port 30312 --rpc --rpcaddr 'localhost' --rpcport 8502 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --bootnodes 'enode://a0cd5bf310b8c30396c80cb82d053c522a3bb5c0e3f483f78abeddfc64b351d2401222f7b8f3425670dfca0c395c11d3b1290d8d5e6655e0fcba957c73cbd679@127.0.0.1:30310' --networkid 1515 --gasprice '1' -unlock '0xfd15d8dc8a53f07fdad6b980e291ed790f864255' --password node2/password.txt --mine`
* Send an Ether transaction between nodes:
  * In new terminal, run `$CustomGeth attach 'http://localhost:8501'`
  * In the Eth JSON RPC console, run `eth.sendTransaction({'from':eth.coinbase, 'to':'db97df08c187fb9c7f46b34b00200eaa95e321c3', 'value':web3.toWei(3, 'ether')})`

* Open `miner.logs` and watch as the transaction goes from:
  * pool queue
  * validation
  * promotion
  * discovery by worker, categorized as remote or local
  * sorted by nonce and gas price 
  * commit
  * execution
