package database

import (
	config "ConnectApp/src/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() error {
	db_values := config.Db_Config()
	source := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", db_values.Username, db_values.Password, db_values.Protocol, db_values.Ip_address, db_values.Port, db_values.Database)
	database, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	database.SetMaxOpenConns(20)
	DB = database
	return nil
}
