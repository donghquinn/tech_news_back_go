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

func GetTodayHackerNewsList(page string, size string) ([]types.HackerNewsResponse, error) {
	todayDate := time.Now()

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return []types.HackerNewsResponse{}, dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[Hacker] Convert String To INT: %v", convErr)
		return []types.HackerNewsResponse{}, convErr
	}

	queryResult, queryErr := database.Query(dbCon, queries.GetTodayHackerByDate, fmt.Sprintf("%s", todayDate), fmt.Sprintf("%d", pageInt - 1), size)

	if queryErr != nil {
		return []types.HackerNewsResponse{}, queryErr
	}

	defer dbCon.Close()

	var hackerNewsList []types.HackerNewsResponse

	for queryResult.Next() {
		row := types.HackerNewsResponse{}

		scanErr := queryResult.Scan(&row)

		if scanErr != nil {
			log.Printf("[Hacker] Get Today Hacker News: %v", scanErr)
			return []types.HackerNewsResponse{}, scanErr
		}

		hackerNewsList = append(hackerNewsList, row)
	}

	return hackerNewsList, nil
}