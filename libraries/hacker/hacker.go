package hacker

import (
	"context"
	"log"
	"time"

	"github.com/dongquinn/tech_news_back_go/prisma/db"
	"github.com/dongquinn/tech_news_back_go/types"
)

func GetHackerNews(client *db.PrismaClient, today string) []types.HackerNewsResponse {
	context := context.Background()

		// Layout for parsing
	layout := "2006-01-02 15:04:05.000"
	dateTime, timeErr := time.Parse(layout, today)

	if timeErr != nil {
		log.Fatalln(timeErr)
	}
	year, month, day := dateTime.Year(), dateTime.Month(), dateTime.Day()

	result, queryErr  := client.Hackers.FindMany(
		db.Hackers.Founded.Gte(time.Date(year, month, day, 0,0,0,0,time.UTC)),
		db.Hackers.Founded.Lte(time.Date(year, month, day, 23, 59, 59, 59, time.UTC)),
	).Exec(context)

	defer func() {
   	 if err := client.Prisma.Disconnect(); err != nil {
      panic(err)
    }
  }()

  if queryErr != nil {
	log.Fatalln(queryErr)
  }

//   log.Println(result)
  returnData := make([]types.HackerNewsResponse, 0)

  for _, data := range result {
	d := types.HackerNewsResponse {
		Uuid: data.UUID,
		Post: data.Post,
		Link: data.Link,
		Founded: data.Founded,
	}

	returnData = append(returnData, d)
  }

  return returnData
}