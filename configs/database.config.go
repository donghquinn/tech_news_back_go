package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfigStruct struct {
	Port string
	Host string
	User string
	Password string
	Database string
}

var DatabaseConfig DatabaseConfigStruct

func SetDatabaseConfig() {
	envErr := godotenv.Load(".postgres.env")

	if envErr != nil {
		log.Printf("[MAIN] Load Database ENV Err: %v", envErr)
	}

	DatabaseConfig.Database = os.Getenv("POSTGRES_DB")
	DatabaseConfig.Port = os.Getenv("POSTGRES_PORT")
	DatabaseConfig.Host = os.Getenv("POSTGRES_HOST")
	DatabaseConfig.User = os.Getenv("POSTGRES_USER")
	DatabaseConfig.Password = os.Getenv("POSTGRES_PASSWORD")
}