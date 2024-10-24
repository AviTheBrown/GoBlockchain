package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GeneratHash(id int, timestamp int64, preHash string) string {
	data := fmt.Sprintf("%d%d%s", id, timestamp, preHash)
	fmt.Println(data)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
