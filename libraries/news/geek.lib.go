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

func GetTodayGeekNewsList(receivedToday string,page string, size string) ([]types.GeekNewsResponse, error) {
	todayDate, parseErr := time.Parse("2006-01-02", receivedToday)

	if parseErr != nil {
		return []types.GeekNewsResponse{}, parseErr
	}
	
	today := fmt.Sprintf("%d-%d-%d",
		todayDate.Year(),
		todayDate.Month(),
		todayDate.Day())

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return []types.GeekNewsResponse{}, dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[Geek] Convert String To INT: %v", convErr)
		return []types.GeekNewsResponse{}, convErr
	}

	queryResult, queryErr := database.Query(dbCon, queries.GetGeekTodayNewsByDate, today, fmt.Sprintf("%d", pageInt - 1), size)

	if queryErr != nil {
		return []types.GeekNewsResponse{}, queryErr
	}

	defer dbCon.Close()

	var hackerNewsList []types.GeekNewsResponse

	for queryResult.Next() {
		row := types.GeekNewsResponse{}

		scanErr := queryResult.Scan(&row)

		if scanErr != nil {
			log.Printf("[Geek] Get Today Hacker News: %v", scanErr)
			return []types.GeekNewsResponse{}, scanErr
		}

		hackerNewsList = append(hackerNewsList, row)
	}

	return hackerNewsList, nil
}