package main

import (
	blockdata "github.com/AviTheBrown/Go_Blockchain/block_data"
)

func main() {
	gensis := blockdata.CreateGenesisBlock()
	gensis.PrintBlockInfo()

	secondBlock := blockdata.CreateNewBlock(gensis, 2)
	secondBlock.PrintBlockInfo()

	thirdBlock := blockdata.CreateNewBlock(secondBlock, 3)
	thirdBlock.PrintBlockInfo()

	fourthBlock := blockdata.CreateNewBlock(thirdBlock, 4)
	fourthBlock.PrintBlockInfo()

	fithBlock := blockdata.CreateNewBlock(fourthBlock, 5)
	fithBlock.PrintBlockInfo()
}
