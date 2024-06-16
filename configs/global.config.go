package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type GlobalConfigConfig struct {
	AppPort string
	JwtKey string
	AesIv string
	AesKey string
	SecretKey string
}

var GlobalConfig GlobalConfigConfig

func SetGlobalConfig() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf("[MAIN] Load Global ENV Err: %v", envErr)
	}

	GlobalConfig.AppPort = os.Getenv(".env")
	GlobalConfig.AppPort = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	GlobalConfig.AesIv = os.Getenv("AES_IV")
	GlobalConfig.AesKey = os.Getenv("AES_KEY")
	GlobalConfig.JwtKey = os.Getenv("JWT_KEY")
	GlobalConfig.SecretKey = os.Getenv("SECRET_KEY")
}