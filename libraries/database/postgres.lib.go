package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dongquinn/tech_news_back_go/assets"
	"github.com/dongquinn/tech_news_back_go/configs"

	_ "github.com/lib/pq"
)

func InitPostgres() (*sql.DB, error) {
	config := configs.DatabaseConfig

	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)

	postgres, postgresErr := sql.Open("postgres", dbUrl)

	if postgresErr != nil {
		log.Printf("[POSTGRES] Init Postgres Err: %v", postgresErr) 
		return nil, postgresErr
	}

	pingErr := postgres.Ping()

	if pingErr != nil {
		log.Printf("[POSTGRES] Ping Error: %v", pingErr)
		return nil, pingErr
	}
	return postgres, nil
}

func StartDatabaseServer() error {
	pg, pgErr := InitPostgres()

	if pgErr != nil {
		return pgErr
	}
	
	CreateTables(pg, assets.QueriesTransaction)

	return nil
}

func Query(postgres *sql.DB, querystring string, args ...string) (*sql.Rows, error){
	var arguments []interface{}

    for _, arg := range args {
        arguments = append(arguments, arg)
    }	

	rows, queryErr := postgres.Query(querystring, arguments...)

	if queryErr != nil {
		log.Printf("[POSTGRES] Query Multiple Rows Error: %v", queryErr)
		return nil, queryErr
	}

	return rows, nil
}

func QueryOne(postgres *sql.DB, querystring string, args ...string) (*sql.Row){
	var arguments []interface{}

    for _, arg := range args {
        arguments = append(arguments, arg)
    }	

	rows := postgres.QueryRow(querystring, arguments...)
	// if queryErr != nil {
	// 	log.Printf("[POSTGRES] Query Multiple Rows Error: %v", queryErr)
	// 	return nil, queryErr
	// }

	return rows
}

func Insert(postgres *sql.DB, querystring string,  args ...string) (int64, error) {
	var arguments []interface{}

    for _, arg := range args {
        arguments = append(arguments, arg)
    }

	 insertResult, insertErr := postgres.Exec(querystring, arguments...)

	if insertErr != nil {
		log.Printf("[POSTGRES] Insert Data Error: %v", insertErr)
		return -9999,insertErr
	}

	insertId, insertIdErr :=  insertResult.LastInsertId()

	if insertIdErr != nil {
		log.Printf("[POSTGRES] Get Insert Id Error: %v", insertIdErr)

		return insertId, insertIdErr
	}

	return insertId, nil
}

func CreateTables(postgres *sql.DB, querystrings []string) ( error) {
	ctx := context.Background()

	tx, txErr := postgres.BeginTx(ctx, nil)

	if txErr != nil {
		log.Printf("[POSTGRES] Create Tables Transaction Begin Error: %v", txErr)
		return txErr
	}

	for _, query := range(querystrings) {
		_, execErr := tx.ExecContext(ctx, query)

		if execErr != nil {
			tx.Rollback()
			log.Printf("[POSTGRES] Create Tables Exec Contexts Error: %v", execErr)
			return execErr
		}
	}

	commitErr := tx.Commit()

	if commitErr != nil {
		log.Printf("[POSTGRES] Create Tables Commit Error: %v", commitErr)
		return commitErr
	}

	return nil
}
