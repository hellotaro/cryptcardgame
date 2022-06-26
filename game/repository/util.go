package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func GetDBCon(dbname string) *sql.DB {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		panic(err)
	}
	return db
}

func InitTable(db *sql.DB, tableName string, columns string) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS "` + tableName + `" ` + columns,
	)
	if err != nil {
		panic(err)
	}
}
