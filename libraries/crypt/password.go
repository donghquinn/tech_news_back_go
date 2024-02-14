package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func EncryptPassword(password string) (string, error) {
	key := os.Getenv("SECRET_KEY")
	//Since the key is in string, we need to convert decode it to bytes
	keyBytes, _ := hex.DecodeString(key)
	plaintext := []byte(password)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	
	return fmt.Sprintf("%x", ciphertext), nil
}

func DecryptPassword(encryptedString string, token string) (string, error) {
	key := os.Getenv("SECRET_KEY")
	keyBytes, _ := hex.DecodeString(key)
	tokenBytes, _ := hex.DecodeString(token)
	
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, enc, tokenBytes, nil)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}
// func CryptoPassword(secretKey string, password string) (string, string) {
// 	key := os.Getenv("SECRET_KEY")
// 	iv := createIv()

// 	encoded := Ase256(password, key, iv, aes.BlockSize)

// 	log.Printf("Key: %s\n IV: %s\n EncodedPassword: %s", key, iv, encoded)

// 	return encoded, iv
// }

// func createIv() string {
// 	var iv [16]byte

//     return string(iv[:])
// }

// func Ase256(plaintext string, key string, iv string, blockSize int) string {
// 	bKey := []byte(key)
// 	bIV := []byte(iv)
// 	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
// 	block, _ := aes.NewCipher(bKey)
// 	ciphertext := make([]byte, len(bPlaintext))
// 	mode := cipher.NewCBCEncrypter(block, bIV)
// 	mode.CryptBlocks(ciphertext, bPlaintext)

// 	return hex.EncodeToString(ciphertext)
// }

// func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
// 	padding := (blockSize - len(ciphertext)%blockSize)
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

// 	return append(ciphertext, padtext...)
// }