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

func GetTodayMlNewsList(page string, size string) ([]types.MachineLEarningNewsResponse, error) {
	todayDate := time.Now()

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return []types.MachineLEarningNewsResponse{}, dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[ML] Convert String To INT: %v", convErr)
		return []types.MachineLEarningNewsResponse{}, convErr
	}

	queryResult, queryErr := database.Query(dbCon, queries.GetGeekTodayNewsByDate, fmt.Sprintf("%s", todayDate), fmt.Sprintf("%d", pageInt - 1), size)

	if queryErr != nil {
		return []types.MachineLEarningNewsResponse{}, queryErr
	}

	defer dbCon.Close()

	var hackerNewsList []types.MachineLEarningNewsResponse

	for queryResult.Next() {
		row := types.MachineLEarningNewsResponse{}

		scanErr := queryResult.Scan(&row)

		if scanErr != nil {
			log.Printf("[ML] Get Today Hacker News: %v", scanErr)
			return []types.MachineLEarningNewsResponse{}, scanErr
		}

		hackerNewsList = append(hackerNewsList, row)
	}

	return hackerNewsList, nil
}