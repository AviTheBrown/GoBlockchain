package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GeneratHash(id int, timestamp int64, preHash string) string {
	data := fmt.Sprintf("%d%s%s", id, timestamp, preHash)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
