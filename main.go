package main

import (
	"github.com/hunman89/nomadcoin/blockchain"
	"github.com/hunman89/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
