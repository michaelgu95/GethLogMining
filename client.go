package main

import (
    "fmt"
    "log"
    "math/big"

    "github.com/onrik/ethrpc"
)

func main() {
    client := ethrpc.New("http://127.0.0.1:8501")
    coinbase, err := client.EthCoinbase()
    if err != nil {
        log.Fatal(err)
    }

    // Send initial txn
    fmt.Println("Test nonce gap: sending initial txn")

    txid, err := client.EthSendTransaction(ethrpc.T{
        From: coinbase,
        To:    "0xdb97df08c187fb9c7f46b34b00200eaa95e321c3",
        Value: ethrpc.Eth1(),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("initial txid: %s", txid)

    // Send txn with nonce 2
    fmt.Println("Test nonce gap: sending txn with nonce 2")

    txid, err = client.EthSendTransaction(ethrpc.T{
        From: coinbase,
        To:    "0xdb97df08c187fb9c7f46b34b00200eaa95e321c3",
        Value: ethrpc.Eth1(),
	Nonce: 2,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("txid: %s", txid)

    // Send txn with nonce 1
    fmt.Println("Test nonce gap: sending txn with nonce 1")

    txid, err = client.EthSendTransaction(ethrpc.T{
        From: coinbase,
        To:    "0xdb97df08c187fb9c7f46b34b00200eaa95e321c3",
        Value: ethrpc.Eth1(),
	Nonce: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("txid: %s", txid)

    // Test multiple transactions from multiple accounts
    fmt.Println("Test multiple accounts: sending txn with nonce 4, gasPrice=20gwei")

    txid, err = client.EthSendTransaction(ethrpc.T{
        From: coinbase,
        To:    "0xdb97df08c187fb9c7f46b34b00200eaa95e321c3",
        Value: ethrpc.Eth1(),
	Nonce: 4,
	GasPrice: big.NewInt(20000000000),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("txid: %s", txid)

    // Send txn from different account with nonce=4, gasPrice=40gwei
    fmt.Println("Test multiple accounts: sending txn with nonce 4 gasPrice=40gwei")

    txid, err = client.EthSendTransaction(ethrpc.T{
        From: "0xfd15d8dc8a53f07fdad6b980e291ed790f864255",
        To:    "0xdb97df08c187fb9c7f46b34b00200eaa95e321c3",
        Value: ethrpc.Eth1(),
	Nonce: 4,
	GasPrice: big.NewInt(40000000000),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("txid: %s", txid)

    // Send txn with nonce=3 to process the previous 2 queued
    fmt.Println("Test multiple accounts: sending txn with nonce 3 to process the previous 2 queued txns")

    txid, err = client.EthSendTransaction(ethrpc.T{
        From: coinbase,
        To:    "0xdb97df08c187fb9c7f46b34b00200eaa95e321c3",
        Value: ethrpc.Eth1(),
	Nonce: 3,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("txid: %s", txid)
}

