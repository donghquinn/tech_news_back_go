package user

import (
	"log"

	"github.com/dongquinn/tech_news_back_go/libraries/crypt"
	"github.com/dongquinn/tech_news_back_go/libraries/database"
	queries "github.com/dongquinn/tech_news_back_go/queries/users"
	types "github.com/dongquinn/tech_news_back_go/types/user"
)

func GetLoginUserData(email string) (types.LoginUserQueryResult, error){
	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		return types.LoginUserQueryResult{}, dbErr
	}

	queryResult := database.QueryOne(dbCon, queries.GetUserData, email)
	

	defer dbCon.Close()

	var loginData types.LoginUserQueryResult

	scanErr := queryResult.Scan(
		&loginData.Uuid,
		&loginData.Email,
		&loginData.Password,
		&loginData.UserStatus)

	if scanErr != nil {
		log.Printf("[LOGIN] Scan User Data Error: %v", scanErr)
		return types.LoginUserQueryResult{}, scanErr
	}

	return loginData, nil
}

func GetDecodeUserData(userRequest types.LoginRequestStruct) (string, string, error) {
	decodedEmail, decodeEmailErr := crypt.DecryptString(userRequest.Email)

	if decodeEmailErr != nil {
		log.Printf("[LOGIN] Decode Email Error: %v", decodeEmailErr)
		return "", "", decodeEmailErr
	}
	
	decodedPassword, decodePasswordErr := crypt.DecryptString(userRequest.Password)

	if decodePasswordErr != nil {
		log.Printf("[LOGIN] Decode Email Error: %v", decodePasswordErr)
		return "", "", decodePasswordErr
	}

	return decodedEmail, decodedPassword, nil
}