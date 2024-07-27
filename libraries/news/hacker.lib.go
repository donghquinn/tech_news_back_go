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

func GetTodayHackerNewsList(receivedToday string,page string, size string) ([]types.HackerNewsResponse, error) {
	todayDate, parseErr := time.Parse("2006-01-02", receivedToday)

	if parseErr != nil {
		return []types.HackerNewsResponse{}, parseErr
	}
	
	today := fmt.Sprintf("%d-%d-%d",
		todayDate.Year(),
		todayDate.Month(),
		todayDate.Day())

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return []types.HackerNewsResponse{}, dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[Hacker] Convert String To INT: %v", convErr)
		return []types.HackerNewsResponse{}, convErr
	}

	queryResult, queryErr := database.Query(dbCon, queries.GetTodayHackerByDate, today, fmt.Sprintf("%d", pageInt - 1), size)

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