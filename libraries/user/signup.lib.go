package user

import (
	"log"

	"github.com/dongquinn/tech_news_back_go/libraries/database"
	queries "github.com/dongquinn/tech_news_back_go/queries/users"
	"github.com/dongquinn/tech_news_back_go/utilities"
)

func SignupUser(email string, name string, password string) error {
	encodedEmail, encodedName, encryptedPassword, encryptErr := EncryptUserData(email, name, password)

	if encryptErr != nil {
		log.Printf("[SIGNUP] Encrypt Password: %v", encryptErr)
		return nil
	}

	dbCon, dbErr := database.InitPostgres()

	if dbErr != nil {
		log.Printf("[SIGNUP] Signup Error: %v", dbErr)
		return dbErr
	}

	_, insertErr := database.Insert(dbCon, queries.InsertNewUser, encodedEmail, encodedName, encryptedPassword)

	if insertErr != nil {
		log.Printf("[SIGNUP] Insert New User Error: %v", insertErr)
		return insertErr
	}
	
	return nil
}

func EncryptUserData(email string, name string, password string) (string, string, string, error) {
	encryptedPassword, encryptErr := utilities.EncryptPassword(password)

	if encryptErr != nil {

		return "", "", "", encryptErr
	}
	encrypteddEmail := utilities.EncodeBase64(email)
	encryptedName := utilities.EncodeBase64(name)

	return encrypteddEmail, encryptedName, string(encryptedPassword), nil
}