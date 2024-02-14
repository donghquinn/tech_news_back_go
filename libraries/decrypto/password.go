package decrypto

import "os"

func DecryptoPassword(password string, token string) string {
	key := os.Getenv("SECRET_KEY")

	return key
}