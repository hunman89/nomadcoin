package main

import (
	"github.com/hunman89/nomadcoin/cli"
	"github.com/hunman89/nomadcoin/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
