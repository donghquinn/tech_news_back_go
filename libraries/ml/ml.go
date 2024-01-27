package ml

import (
	"context"
	"log"
	"time"

	"github.com/dongquinn/tech_news_back_go/prisma/db"
	"github.com/dongquinn/tech_news_back_go/types"
)

func GetMlNews(client *db.PrismaClient, today string) []types.MachineLEarningNewsResponse {
	context := context.Background()

	layout := "2006-01-02 15:04:05.000"
	dateTime, dateErr := time.Parse(layout, today)

	if dateErr != nil {
		log.Fatalln(dateErr)
	}

	year, month, day := dateTime.Year(), dateTime.Month(), dateTime.Day()

	startDate := time.Date(year, month, day, 0 ,0, 0, 0, time.UTC)
	endDate := time.Date(year, month, day, 23, 59, 59, 59, time.UTC)

	result, queryErr  := client.MachineNews.FindMany(
		db.MachineNews.Founded.Gte(startDate),
		db.MachineNews.Founded.Lte(endDate),
	).Exec(context)

	defer func() {
   	 if err := client.Prisma.Disconnect(); err != nil {
      panic(err)
    }
  }()

  if queryErr != nil {
	log.Fatalln(queryErr)
  }

  returnData := make([]types.MachineLEarningNewsResponse, 0)

  for _, data := range result {
	d := types.MachineLEarningNewsResponse {
		Uuid: data.UUID,
		Title: data.Title,
		Category: data.Category,
		Link: data.Link,
		Founded: data.Founded,
	}

	returnData = append(returnData, d)
  }

  return returnData
}