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

func GetTodayGeekNewsList(page string, size string) ([]types.GeekNewsResponse, error) {
	todayDate := time.Now()

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return []types.GeekNewsResponse{}, dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[Geek] Convert String To INT: %v", convErr)
		return []types.GeekNewsResponse{}, convErr
	}

	queryResult, queryErr := database.Query(dbCon, queries.GetGeekTodayNewsByDate, fmt.Sprintf("%s", todayDate), fmt.Sprintf("%d", pageInt - 1), size)

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