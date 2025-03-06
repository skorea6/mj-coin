package main

import (
	"mjcoin/explorer"
	"mjcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
