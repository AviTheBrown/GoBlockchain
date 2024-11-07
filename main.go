package main

import (
	"fmt"
	"strings"

	blockdata "github.com/AviTheBrown/Go_Blockchain/block_data"
)

func askUser() int {
	var number int
	fmt.Println("Enter how many zeros the hash must start with: ")
	fmt.Scan(&number)
	fmt.Println()
	return number
}
func createZero(count int) string {
	return fmt.Sprintf(strings.Repeat("0", count))
}
func main() {
	getUserNumber := askUser()
	zeros := createZero(getUserNumber)

	gensis := blockdata.CreateGenesisBlock(zeros)
	gensis.PrintBlockInfo()

	secondBlock := blockdata.CreateNewBlock(gensis, 2, zeros)
	secondBlock.PrintBlockInfo()

	thirdBlock := blockdata.CreateNewBlock(secondBlock, 3, zeros)
	thirdBlock.PrintBlockInfo()

	fourthBlock := blockdata.CreateNewBlock(thirdBlock, 4, zeros)
	fourthBlock.PrintBlockInfo()

	fithBlock := blockdata.CreateNewBlock(fourthBlock, 5, zeros)
	fithBlock.PrintBlockInfo()
}
