package user

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dongquinn/tech_news_back_go/libraries/database"
	queries "github.com/dongquinn/tech_news_back_go/queries/users"
	"github.com/dongquinn/tech_news_back_go/types"
)

func GetMyPage(userId string, email string, page string, size string) (
	[]types.HackerNewsResponse,
	[]types.MachineLEarningNewsResponse,
	[]types.GeekNewsResponse,
	error) {
		
	hackerNewsList, mlNewsList, geekNewsList, likedNewsErr := getLikedError(userId, page, size)

	if likedNewsErr != nil {
		return []types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				likedNewsErr
	}

	return hackerNewsList, mlNewsList, geekNewsList, nil
}

func getLikedError(userId string, page string, size string) (
	[]types.HackerNewsResponse,
	[]types.MachineLEarningNewsResponse,
	[]types.GeekNewsResponse,
	error) {
dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		log.Printf("[MYPAGE] Get My Page DB Connection Error: %v", dbErr)
		return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				dbErr
	}

	pageInt, convErr := strconv.Atoi(page)

	if convErr != nil {
		log.Printf("[MYPAGE] Convert String to Int Error: %v", convErr)
		return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				convErr
	}

	geekResult, geekErr := database.Query(dbCon, queries.GetLikedGeekNewsByUserId, userId, fmt.Sprintf("%d", pageInt - 1), size)
	hackerResult, hackerErr := database.Query(dbCon, queries.GetLikedHackerNewsByUserId, userId, fmt.Sprintf("%d", pageInt - 1), size)
	mlResult, mlErr := database.Query(dbCon, queries.GetLikedMlNewsByUserId, userId, fmt.Sprintf("%d", pageInt - 1), size)


	if geekErr != nil {
		log.Printf("[MYPAGE] Query Geek Error: %v", geekErr)
		return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				geekErr
	}

	if hackerErr != nil {
		log.Printf("[MYPAGE] Query Hacker Error: %v", hackerErr)
		return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				hackerErr
	}


	if mlErr != nil {
		log.Printf("[MYPAGE] Query ML Error: %v", mlErr)
		return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				mlErr
	}

	defer dbCon.Close()

	var likedGeekNewsList []types.GeekNewsResponse
	var likedHackerNewsList []types.HackerNewsResponse
	var likedMlNewsList []types.MachineLEarningNewsResponse

	for	geekResult.Next() {
		geekRow := types.GeekNewsResponse{}

		geekScanErr := geekResult.Scan(&geekRow)

		if geekScanErr != nil {
			log.Printf("[MYPAGE] Scan Geek Error: %v", geekScanErr)
			return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				geekScanErr
		}

		likedGeekNewsList = append(likedGeekNewsList, geekRow)
	}

	for hackerResult.Next() {
		hackerRow := types.HackerNewsResponse{}

		hackerScanErr := hackerResult.Scan(&hackerRow)

		if hackerScanErr != nil {
			log.Printf("[MYPAGE] Scan Hacker Error: %v", hackerScanErr)
			return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				hackerScanErr
		}

		likedHackerNewsList = append(likedHackerNewsList, hackerRow)
	}

	for mlResult.Next() {
		mlRow := types.MachineLEarningNewsResponse{}

		mlScanErr := mlResult.Scan(&mlRow)

		if mlScanErr != nil {
			log.Printf("[MYPAGE] Scan ML Error: %v", mlScanErr)
			return 	[]types.HackerNewsResponse{},
				[]types.MachineLEarningNewsResponse{},
				[]types.GeekNewsResponse{},
				mlScanErr
		}

		likedMlNewsList = append(likedMlNewsList, mlRow)
	}

	return likedHackerNewsList, likedMlNewsList, likedGeekNewsList, nil
}