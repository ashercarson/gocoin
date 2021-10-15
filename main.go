package main

import (
	"github.com/ashercarson/gocoin/explorer"
	"github.com/ashercarson/gocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
