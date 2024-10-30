package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	res := hex.EncodeToString(hasher.Sum(nil))
	return res
}

func GetSalt() string {
	tmp := make([]byte, 16)
	rand.Read(tmp)
	salt := hex.EncodeToString(tmp)

	return salt
}
