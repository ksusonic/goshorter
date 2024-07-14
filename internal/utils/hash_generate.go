package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

const hashLen = 6

func GenerateHash(url string) string {
	hash := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:hashLen])
}
