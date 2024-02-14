package users

import (
	"context"

	"github.com/dongquinn/tech_news_back_go/libraries/prisma"
	"github.com/dongquinn/tech_news_back_go/prisma/db"
)

func LoginLibrary(email string, password string) (string, error) {
	context := context.Background()

	prisma, prismaErr := prisma.PrismaClient()

	if prismaErr != nil {
		return "", prismaErr
	}

	queryResult, queryErr := prisma.Client.FindFirst(
		db.Client.Email.Equals(email),
	).Exec(context)

	if queryErr != nil {
		return "", queryErr
	}

	uuid := queryResult.UUID
	
	return uuid, nil
}