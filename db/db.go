package db

import (
	"database/sql"
	"fmt"

	"loosidAPI/config"
)

var dbConfig *config.Database = &config.Cfg.Database

// DbConnection - database connection variable
var DbConnection *sql.DB

// InitDb - initiate database connection
func InitDb() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.DBHost, dbConfig.DBPort,
		dbConfig.Username, dbConfig.Password, dbConfig.DBName)

	DbConnection, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DbConnection.Ping()
	if err != nil {
		panic(err)
	}

	initGuide()
	initListing()

	fmt.Println("Successfully connected!")
}
