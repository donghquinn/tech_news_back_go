package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"
)

func CryptoPassword(secretKey string, password string) (string, string) {
	key := os.Getenv("SECRET_KEY")
	iv := createIv()

	encoded := Ase256(password, key, iv, aes.BlockSize)

	log.Printf("Key: %s\n IV: %s\n EncodedPassword: %s", key, iv, encoded)

	return encoded, iv
}

func createIv() string {
	var iv [16]byte

    return string(iv[:])
}

func Ase256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)

	return hex.EncodeToString(ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	
	return append(ciphertext, padtext...)
}