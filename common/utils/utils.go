package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

// GenRandomNumString 生成任意长度的数字字符串
func GenRandomNumString(lens int) string {
	data := make([]byte, lens)
	for i := 0; i < lens; i++ {
		data[i] = byte(rand.Intn(10) + '0')
	}
	return string(data)
}

// EncString sha256 hash
func EncString(s string) string {
	data := sha256.Sum256([]byte(s))
	enc := hex.EncodeToString(data[:])
	return enc[:32]
}
