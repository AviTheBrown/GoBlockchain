package hashing

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	mathRand "math/rand"
)

func generateMagicNum() int32 {
	max := big.NewInt(int64(mathRand.Int31()))
	randNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatalln("there was problem generating a rand number")
	}
	if randNum.Cmp(big.NewInt(int64(int32(1<<31-1)))) > 0 {
		log.Fatalln("generated number is out of int32 range")
	}
	return int32(randNum.Int64())
}

func GeneratHash(id int, timestamp int64, magicNumber int32, preHash string) string {
	data := fmt.Sprintf("%d%d%s", id, timestamp, generateMagicNum(), preHash)
	fmt.Println(data)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
