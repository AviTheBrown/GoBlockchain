package blockdata

import (
	"fmt"
	"github.com/AviTheBrown/Go_Blockchain/hashing"
	"time"
)

type Block struct {
	Id         int
	Time_stamp int64
	// Data       string
	PrevHash     string
	Hash         string
	Magic_number int32
}

func (b Block) GetPrevBlockHash(prevBlock Block) string {
	return prevBlock.PrevHash

}
func CreateGenesisBlock() Block {
	gBlock := Block{
		Id:           1,
		Time_stamp:   time.Now().UTC().UnixNano(),
		Magic_number: hashing.GenerateMagicNum(),
		PrevHash:     "0",
	}

	gBlock.Hash = hashing.GeneratHash(gBlock.Id, gBlock.Time_stamp, gBlock.Magic_number, gBlock.PrevHash)
	return gBlock
}

func CreateNewBlock(prevBlock Block, id int) Block {
	block := Block{
		Id:           id,
		Time_stamp:   time.Now().UTC().UnixNano(),
		Magic_number: hashing.GenerateMagicNum(),
		PrevHash:     prevBlock.Hash,
	}
	block.Hash = hashing.GeneratHash(block.Id, block.Time_stamp, block.Magic_number, block.PrevHash)
	return block
}

func (b Block) PrintBlockInfo() {
	if b.Id == 1 {
		fmt.Printf("Genesis Block:\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)
	} else {
		fmt.Printf("\nBlock:\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)
		// fmt.Printf("Data: %s\n", b.Data)
	}
}
