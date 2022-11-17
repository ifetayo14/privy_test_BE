package database

import (
	"database/sql"
	"fmt"
	"privy_cake_store/config"

	_ "github.com/go-sql-driver/mysql"
)

func StartDB() (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.USERNAME, config.PASSWORD, config.DBNAME)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error on creating connection: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, err
}
