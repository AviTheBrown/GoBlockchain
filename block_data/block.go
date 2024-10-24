package blockdata

import (
	"fmt"
	"time"

	"github.com/AviTheBrown/Go_Blockchain/hashing"
)

type Block struct {
	Id         int
	Time_stamp int64
	// Data       string
	PrevHash string
	Hash     string
}

func (b Block) GetPrevBlockHash(prevBlock Block) (string, error) {

}
func CreateGenesisBlock() Block {
	gBlock := Block{
		Id:         1,
		Time_stamp: time.Now().UnixNano(),
		PrevHash:   "0",
	}

	gBlock.Hash = hashing.GeneratHash(gBlock.Id, gBlock.Time_stamp, gBlock.PrevHash)
	return gBlock
}

func (b Block) PrintBlockInfo() {
	if b.Id == 1 {
		fmt.Printf("Block\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)

	} else {

		fmt.Printf("\nBlock\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Print("Hash of the block:\n%s\n", b.Hash)
		// fmt.Printf("Data: %s\n", b.Data)
		// fmt.Printf("Previous Hash: %s\n", b.PrevHash)
		// fmt.Println("=====================")
	}
}
