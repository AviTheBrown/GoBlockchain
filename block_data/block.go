package blockdata

import (
	"fmt"
	"log"
	"strconv"
	"time"

	hashing "github.com/AviTheBrown/Go_Blockchain/hashing"
)

type Block struct {
	Id           int
	Magic_number int32
	Time_stamp   int64
	PrevHash     string
	Hash         string
	Gen_time     time.Duration
}

func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := seconds / 60
	remainingSeconds := seconds % 60

	if minutes > 0 {
		return fmt.Sprintf("%d %d .", minutes, remainingSeconds)
	}
	return fmt.Sprintf("%d", remainingSeconds)
}

func (b Block) GetPrevBlockHash(prevBlock Block) string {
	return prevBlock.PrevHash

}
func CreateGenesisBlock(zeros string) Block {
	startTime := time.Now()
	gBlock := Block{
		Id:           1,
		Time_stamp:   time.Now().UTC().UnixNano(),
		Magic_number: hashing.GenerateMagicNum(),
		PrevHash:     "0",
	}

	for {
		gBlock.Magic_number = hashing.GenerateMagicNum()
		hash := hashing.GeneratHash(gBlock.Id, gBlock.Time_stamp, gBlock.Magic_number, gBlock.PrevHash, zeros)
		intZeros, err := strconv.Atoi(zeros)
		if err != nil {
			log.Fatal("Couldnt convert the string `zero` to an int.")
		}
		if hashing.ValidateHash(hash, intZeros) {
			gBlock.Hash = hash
			break
		}
	}

	gBlock.Hash = hashing.GeneratHash(gBlock.Id, gBlock.Time_stamp, gBlock.Magic_number, gBlock.PrevHash, zeros)
	gBlock.Gen_time = time.Since(startTime)
	return gBlock
}

func CreateNewBlock(prevBlock Block, id int, zeros string) Block {
	startTime := time.Now()
	block := Block{
		Id:           id,
		Time_stamp:   time.Now().UTC().UnixNano(),
		Magic_number: hashing.GenerateMagicNum(),
		PrevHash:     prevBlock.Hash,
	}

	for {
		block.Magic_number = hashing.GenerateMagicNum()
		hash := hashing.GeneratHash(block.Id, block.Time_stamp, block.Magic_number, block.PrevHash, zeros)
		intZeros, err := strconv.Atoi(zeros)
		if err != nil {
			log.Fatal("Couldnt convert the string `zero` to an int.")
		}
		if hashing.ValidateHash(hash, intZeros) {
			block.Hash = hash
			break
		}
	}

	block.Hash = hashing.GeneratHash(block.Id, block.Time_stamp, block.Magic_number, block.PrevHash, zeros)
	block.Gen_time = time.Since(startTime)
	return block
}

func (b Block) PrintBlockInfo() {
	if b.Id == 1 {
		fmt.Printf("Genesis Block:\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Magic Number: %d\n", b.Magic_number)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)
		fmt.Printf("Block was generating for %s seconds\n", formatDuration(b.Gen_time))
	} else {
		fmt.Printf("\nBlock:\n")
		fmt.Printf("Id: %d\n", b.Id)
		fmt.Printf("Timestamp: %v\n", b.Time_stamp)
		fmt.Printf("Magic Number: %d\n", b.Magic_number)
		fmt.Printf("Hash of the previous block:\n%s\n", b.PrevHash)
		fmt.Printf("Hash of the block:\n%s\n", b.Hash)
		fmt.Printf("Block was generating for %s seconds\n", formatDuration(b.Gen_time))
	}
}
