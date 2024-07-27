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
	todayDate, parseErr := time.Parse("2006-01-02 15:04:05.000", receivedToday)

	if parseErr != nil {
		log.Printf("Parse Error: %v", parseErr)
		return []types.GeekNewsResponse{}, parseErr
	}

	today := todayDate.Format("2006-01-02 15:04:05")
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

		scanErr := queryResult.Scan(
			&row.Uuid,
			&row.Post,
			&row.DescLink,
			&row.OriginalLink,
			&row.Founded)

		if scanErr != nil {
			log.Printf("[Geek] Get Today Hacker News: %v", scanErr)
			return []types.GeekNewsResponse{}, scanErr
		}

		hackerNewsList = append(hackerNewsList, row)
	}

	return hackerNewsList, nil
}