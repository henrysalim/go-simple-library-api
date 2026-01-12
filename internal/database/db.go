package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB(cfg string) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg)
	if err != nil {
		return nil, err
	}

	//	Set connection pool settings to avoid timeouts/overloads
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
