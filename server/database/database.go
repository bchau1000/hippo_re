package database

import (
	"database/sql"
	"fmt"
	"hippo/config"
	"hippo/logging"

	sq "github.com/Masterminds/squirrel"

	_ "github.com/go-sql-driver/mysql"
)

type Builder interface {
	ToSql() (string, []interface{}, error)
}

var db *sql.DB

// Function to execute a raw SQL query
func ExecuteQuery(query string) (sql.Result, error) {
	result, err := db.Exec(query)

	if err != nil {
		logging.Error.Printf("Error in executing SQL: %v", err)
		return nil, err
	}

	return result, nil
}

// Function to execute a SELECT query
func Search(query sq.SelectBuilder) (*sql.Rows, error) {
	sql, args, err := toSql(query)

	if err != nil {
		return nil, err
	}

	logging.Info.Printf("Executing SQL: (%s) with params: %v", sql, args)
	result, err := db.Query(sql, args...)

	if err != nil {
		logging.Error.Printf("Error occurred while executing search: %v", err)
		return nil, err
	}

	return result, nil
}

// Function to execute an INSERT query
func Insert(query sq.InsertBuilder) (sql.Result, error) {
	sql, args, err := toSql(query)
	if err != nil {
		return nil, err
	}

	result, err := exec(sql, args)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Function to execute a DELETE query
func Delete(query sq.DeleteBuilder) (int64, error) {
	result, err := query.RunWith(db).Exec()

	if err != nil {
		logging.Error.Printf("Error occurred while executing delete: %v", err)
		return 0, nil
	}

	return result.RowsAffected()
}

func toSql(builder Builder) (string, []interface{}, error) {
	result, args, err := builder.ToSql()

	if err != nil {
		logging.Error.Printf("Error occurred while converting SQL: %v", err)
		return "", nil, err
	}

	return result, args, nil
}

func exec(sql string, args []interface{}) (sql.Result, error) {
	logging.Info.Printf("Executing SQL: %s\n%v", sql, args)
	result, err := db.Exec(sql, args...)

	if err != nil {
		logging.Error.Printf("Error encountered executing SQL: %v", err)
		return nil, err
	}

	return result, nil
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
