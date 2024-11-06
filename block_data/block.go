package blockdata

import (
	"fmt"
	"github.com/AviTheBrown/Go_Blockchain/hashing"
	"time"
)

type Block struct {
	Id           int
	Magic_number int32
	Time_stamp   int64
	// Data       string
	PrevHash string
	Hash     string
}

func (b Block) GetPrevBlockHash(prevBlock Block) string {
	return prevBlock.PrevHash

}
func CreateGenesisBlock(zeros string) Block {
	gBlock := Block{
		Id:           1,
		Time_stamp:   time.Now().UTC().UnixNano(),
		Magic_number: hashing.GenerateMagicNum(),
		PrevHash:     "0",
	}

	gBlock.Hash = hashing.GeneratHash(gBlock.Id, gBlock.Time_stamp, gBlock.Magic_number, gBlock.PrevHash, zeros)
	return gBlock
}

func CreateNewBlock(prevBlock Block, id int, zeros string) Block {
	block := Block{
		Id:           id,
		Time_stamp:   time.Now().UTC().UnixNano(),
		Magic_number: hashing.GenerateMagicNum(),
		PrevHash:     prevBlock.Hash,
	}
	block.Hash = hashing.GeneratHash(block.Id, block.Time_stamp, block.Magic_number, block.PrevHash, zeros)
	return block
}

func (b Block) PrintBlockInfo() {
	if b.Id == 1 {
		fmt.Printf("Genesis Block:\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Magic Number: %d\n", b.Magic_number)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)
	} else {
		fmt.Printf("\nBlock:\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Magic Number: %d\n", b.Magic_number)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)
		// fmt.Printf("Data: %s\n", b.Data)
	}
}
