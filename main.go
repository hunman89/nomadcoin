package main

import (
	"github.com/hunman89/nomadcoin/explorer"
	"github.com/hunman89/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
