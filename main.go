package main

import "./blockchain"

func main() {
	chain := blockchain.GetBlockchain()
	println(chain)
}