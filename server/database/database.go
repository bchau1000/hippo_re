package database

import (
	"database/sql"
	"fmt"
	"hippo/config"
	"hippo/logging"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ExecuteQuery(query string) sql.Result {
	result, err := db.Exec(query)

	if err != nil {
		logging.Error.Printf("Error in executing SQL: %v", err)
	}

	return result
}

func Init(config *config.Config) {
	var err error
	db, err = sql.Open(
		"mysql",
		fmt.Sprintf(
			"root:%s@tcp(%s:%d)/%s",
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name))
	if err != nil {
		logging.Fatal.Fatalf("Fatal error encountered while opening DB connection %v", err)
	}

	logging.Info.Printf(
		"Successfully opened database connection: %s:%d/%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)
}
