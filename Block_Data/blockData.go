package blockdata

import (
	"time"
)

type Block struct {
	Id         int
	Time_stamp time.Time
	Hash       string
	Data       string
	PrevHash   string
}

func CreateGenesisBlock() Block {
	return Block{
		Id:         1,
		Time_stamp: time.Now(),
		Hash:       "0",
	}
}
