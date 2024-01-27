package prisma

import "github.com/dongquinn/tech_news_back_go/prisma/db"

func PrismaClient() (*db.PrismaClient, error) {
	client := db.NewClient()

		if err := client.Prisma.Connect(); err != nil {
    return nil, err
  }

  return client, nil
}