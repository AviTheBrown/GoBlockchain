package main

import (
	blockdata "github.com/AviTheBrown/Go_Blockchain/block_data"
)

func main() {
	gensis := blockdata.CreateGenesisBlock()
	gensis.PrintBlockInfo()

}
