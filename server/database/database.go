package database

import (
	"context"
	"database/sql"
	"fmt"
	"hippo/common/errormsg"
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
func ExecuteQuery(ctx context.Context, query string) (sql.Result, error) {
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ExecuteSql, err))
		return nil, err
	}

	return result, nil
}

// Function to execute a SELECT query
func Search(ctx context.Context, query sq.SelectBuilder) (*sql.Rows, error) {
	sql, args, err := toSql(ctx, query)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ConvertSql, err))
		return nil, err
	}

	logging.Info.Printf("Executing SQL: (%s) with params %v", sql, args)

	result, err := db.Query(sql, args...)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ExecuteSql, err))
		return nil, err
	}

	return result, nil
}

// Function to execute an INSERT query
func Insert(ctx context.Context, query sq.InsertBuilder) (sql.Result, error) {
	sql, args, err := toSql(ctx, query)
	if err != nil {
		return nil, err
	}

	result, err := exec(ctx, sql, args)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func InsertTx(
	ctx context.Context,
	tx *sql.Tx,
	query sq.InsertBuilder) (sql.Result, error) {

	sql, args, err := toSql(ctx, query)
	if err != nil {
		logging.Error.Panic(errormsg.FormatError(ctx, errormsg.ConvertSql, err))
		return nil, err
	}

	result, err := tx.ExecContext(ctx, sql, args)
	if err != nil {
		logging.Error.Panic(errormsg.FormatError(ctx, errormsg.ExecuteSql, err))
		return nil, err
	}

	return result, nil
}

// Function to execute a DELETE query
func Delete(ctx context.Context, query sq.DeleteBuilder) (int64, error) {
	sql, args, err := toSql(ctx, query)
	if err != nil {
		return 0, nil
	}

	result, err := exec(ctx, sql, args)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ExecuteSql, err))
		return 0, nil
	}

	return result.RowsAffected()
}

func Transaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, "Error occurred while beginning transaction", err))
		return err
	}
	defer tx.Rollback()

	err = fn(tx)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, "Error occurred during transaction", err))
		return err
	}

	err = tx.Commit()
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, "Error occured after committing transaction", err))
		return err
	}

	return nil
}

func toSql(ctx context.Context, builder Builder) (string, []interface{}, error) {
	result, args, err := builder.ToSql()
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ConvertSql, err))
		return "", nil, err
	}

	return result, args, nil
}

func exec(ctx context.Context, sql string, args []interface{}) (sql.Result, error) {
	logging.Info.Printf("Executing SQL: %s\n%v", sql, args)

	result, err := db.ExecContext(ctx, sql, args...)
	if err != nil {
		logging.Error.Print(errormsg.FormatError(ctx, errormsg.ExecuteSql, err))
		return nil, err
	}

	return result, nil
}

func InitDatabase(ctx context.Context, config *config.Config) {
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
