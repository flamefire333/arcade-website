package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DatabaseInfo string `json:"databaseInfo"`
}

var Conn *sql.DB

func init() {
	file, err := os.Open("config")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	Conn, err = sql.Open("mysql", config.DatabaseInfo)
	if err != nil {
		log.Fatal(err)
	}
	_, err = Conn.Query("SELECT id, name, avatar, group_id FROM characters")
	if err != nil {
		log.Printf("Character FIRST SELECT failed %+v\n", err)
		return
	}
}
