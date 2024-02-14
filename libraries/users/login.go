package users

import (
	"context"
	"errors"
	"log"

	"github.com/dongquinn/tech_news_back_go/libraries/crypt"
	"github.com/dongquinn/tech_news_back_go/libraries/prisma"
	"github.com/dongquinn/tech_news_back_go/libraries/redis"
	"github.com/dongquinn/tech_news_back_go/prisma/db"
)

func LoginLibrary(email string, password string) (string, error) {
	context := context.Background()

	encodedPassword, encodErr := crypt.EncryptPassword(password)

	if encodErr != nil {
		return "", nil
	}

	client := redis.RedisClient()

	prisma, prismaErr := prisma.PrismaClient()

	if prismaErr != nil {
		return "", prismaErr
	}

	queryResult, queryErr := prisma.Client.FindFirst(
		db.Client.Email.Equals(email),
		db.Client.Password.Equals(encodedPassword),
	).Exec(context)

	if queryErr != nil {
		return "", queryErr
	}

	uuid := queryResult.UUID
	token := queryResult.PasswordToken
	encryptedPassword := queryResult.Password

	isSame := comparePassword(password, encryptedPassword, token)

	if !isSame {
		return "", errors.New("Given Password is Not Match")
	}

	setErr := redis.SetItem(client, email, uuid)

	if setErr != nil {
		return "", setErr
	}

	return uuid, nil
}

func comparePassword(rawPassword string, encodedPassword string, token string) bool {
	decodedPassword, decryptErr := crypt.DecryptPassword(encodedPassword, token)

	if decryptErr != nil {
		return false
	}

	log.Printf("Raw: %s\nDecoded: %s\nEncoded: %s\nToken: %s\n", rawPassword, decodedPassword, encodedPassword, token)

	if rawPassword != decodedPassword {
		return false
	}

	return true
}