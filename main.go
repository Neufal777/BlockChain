package main

import "fmt"

func main() {

	bc := NewBlockchain()

	bc.AddBlock("0.6 BTC a Ainhoa")
	bc.AddBlock("3 BTC a paula")
	bc.AddBlock("1.6 BTC a miriam")
	bc.AddBlock("7 BTC a Tete")

	for _, block := range bc.blocks {

		fmt.Printf("Hash blocke anterior: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash bloque actual: %x\n\n", block.Hash)
	}

}
