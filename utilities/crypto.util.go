package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/dongquinn/tech_news_back_go/configs"
	"golang.org/x/crypto/bcrypt"
)


const (
	MinCost int = 4
	MaxCost int = 31
	DefaultCost int = 10
)

var globalConfig = configs.GlobalConfig

func EncodeBase64(str string) string {
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))

	return encodedStr
}

func DecodeBase64(str string) ([]byte, error) {
	decodedStr, decodErr := base64.StdEncoding.DecodeString(str)

	if decodErr != nil {
		log.Printf("[CRYPT] Decode String With Base64 Error: %v", decodErr)

		return nil, decodErr
	}

	return decodedStr, nil
}

func EncryptPassword(password string) ([]byte, error){
	passwordByte := []byte(password)

	encodedPassword, genErr := bcrypt.GenerateFromPassword(passwordByte, DefaultCost)

	if genErr != nil {
		log.Printf("[CRYPT] Encrypt Password Err: %v", genErr)
		return nil, genErr
	}

	return encodedPassword, nil
}

// 패스워드 비교. DB 의 패스워드는 암호화 후 BASE64로 인코딩되므로, 디코딩 후 넘겨줘야 함
func ComparePassword(givenPassword string, dbPassword string) bool {
	matchErr := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(dbPassword))

	if matchErr != nil {
		log.Printf("[CRYPT] Compare Passwords Does Not Match: %v", matchErr)

		return false
	}

	return true
}

// AES256-GCM 암호화
func AES256GSMEncrypt(plaintext []byte) ([]byte, []byte, error) {
  if len(globalConfig.SecretKey) != 32 {
    return nil, nil, fmt.Errorf("[CRYPT] secret key is not for AES-256: total %d bits", 8*len(globalConfig.SecretKey))
  }

  // prepare AES-256-GSM cipher
  block, err := aes.NewCipher([]byte(globalConfig.SecretKey))
  if err != nil {
	log.Printf("[CRYPT] Create New Cipher Error: %v", err)
    return nil, nil, err
  }

  aesgcm, err := cipher.NewGCM(block)
  if err != nil {
	log.Printf("[CRYPT] Create New GCM Instance Error: %v", err)
    return nil, nil, err
  }

  // make random nonce
  nonce := make([]byte, 12)

  if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	log.Printf("[CRYPT] Read Nonce Error: %v", err)
    return nil, nil, err
  }

  // encrypt plaintext
  ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

  return ciphertext, nonce, nil
}

// AES256-GCM 복호화
func AES256GSMDecrypt(encryptedString string, nonce string) ([]byte, error) {
  if len(globalConfig.SecretKey) != 32 {
    return nil, fmt.Errorf("[CRYPT] secret key is not for AES-256: total %d bits", 8*len(globalConfig.SecretKey))
  }

  // prepare AES-256-GSM cipher
  block, err := aes.NewCipher([]byte(globalConfig.SecretKey))
  if err != nil {
	log.Printf("[CRYPT] Create New Cipher Error: %v", err)
    return nil, err
  }

  aesgcm, err := cipher.NewGCM(block)

  if err != nil {
	log.Printf("[CRYPT] Create New GCM Instance Error: %v", err)

	return nil, err
  }

  nonceSize := aesgcm.NonceSize()

  ciphertext := nonce + encryptedString

  nonce, pureCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
  
  // decrypt ciphertext
  plaintext, err := aesgcm.Open(nil, []byte(nonce), []byte(pureCiphertext), nil)

  if err != nil {
    return nil, err
  }

  return plaintext, nil
}