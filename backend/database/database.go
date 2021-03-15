package database

import (
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DatabaseInfo string `json:"databaseInfo"`
}

func GetDBConn() (*sql.DB, error) {
	file, err := os.Open("config")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return sql.Open("mysql", config.DatabaseInfo)
}
