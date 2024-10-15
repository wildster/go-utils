package goutils

import (
	"crypto/rand"
	"encoding/hex"
	"path/filepath"
)

func GenerateRandomKey(fileName string) (string, error) {
	ext := filepath.Ext(fileName)

	randomBytes := make([]byte, 32)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	randomString := hex.EncodeToString(randomBytes)

	randomFileName := randomString + ext
	return randomFileName, nil
}
