package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	res := hex.EncodeToString(hasher.Sum(nil))
	return res
}
