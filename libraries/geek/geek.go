package geek

import (
	"context"
	"log"
	"time"

	"github.com/dongquinn/tech_news_back_go/prisma/db"
	"github.com/dongquinn/tech_news_back_go/types"
)

func GetDailyGeekNews(client *db.PrismaClient, today string) []types.GeekNewsResponse {
	ctx := context.Background()


	todayStart, timeErr := time.Parse(time.RFC3339,today)

	if timeErr != nil {
		log.Fatalln(timeErr)
	}

	year, month, day := todayStart.Year(),todayStart.Month(), todayStart.Day()
	log.Printf("Year: %d, Month: %d, Day: %d", year, month, day)

	result, queryErr := client.Hada.FindMany(
		db.Hada.Founded.Gte(time.Date(year, month, day, 0,0,0,0,time.UTC)),
		db.Hada.Founded.Lte(time.Date(year, month, day, 23,59,59,59,time.UTC)),
	).Exec(ctx)

	if queryErr != nil {
		log.Fatalln(queryErr)
	}
	
	defer func() {
   	 if err := client.Prisma.Disconnect(); err != nil {
      panic(err)
    }
  }()
  
	returnArray := make([]types.GeekNewsResponse, len(result))
	log.Println(result)

	for _, data := range result {
		dd := types.GeekNewsResponse{Uuid: data.UUID, Post: data.Post, OriginalLink: data.Link, DescLink: data.DescLink, Founded: data.Founded}
		returnArray = append(returnArray, dd)
	}

	return returnArray
}