package news

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dongquinn/tech_news_back_go/libraries/database"
	queries "github.com/dongquinn/tech_news_back_go/queries/news"
	"github.com/dongquinn/tech_news_back_go/types"
)

func GetTodayMlNewsList(receivedToday string,page string, size string) ([]types.MachineLEarningNewsResponse, error) {
	todayDate, parseErr := time.Parse("2006-01-02 15:04:05.000", receivedToday)

	if parseErr != nil {
		log.Printf("Parse Error: %v", parseErr)
		return []types.MachineLEarningNewsResponse{}, parseErr
	}

	today := todayDate.Format("2006-01-02 15:04:05")

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return []types.MachineLEarningNewsResponse{}, dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[ML] Convert String To INT: %v", convErr)
		return []types.MachineLEarningNewsResponse{}, convErr
	}

	queryResult, queryErr := database.Query(dbCon, queries.GetTodayMlByDate, today, fmt.Sprintf("%d", pageInt - 1), size)

	if queryErr != nil {
		return []types.MachineLEarningNewsResponse{}, queryErr
	}

	defer dbCon.Close()

	var hackerNewsList []types.MachineLEarningNewsResponse

	for queryResult.Next() {
		row := types.MachineLEarningNewsResponse{}

		scanErr := queryResult.Scan(
			&row.Uuid,
			&row.Category,
			&row.Title,
			&row.Link,
			&row.Founded)

		if scanErr != nil {
			log.Printf("[ML] Get Today Hacker News: %v", scanErr)
			return []types.MachineLEarningNewsResponse{}, scanErr
		}

		hackerNewsList = append(hackerNewsList, row)
	}

	return hackerNewsList, nil
}