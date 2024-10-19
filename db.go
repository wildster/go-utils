package goutils

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	DB *sql.DB
)

func initDB() {
	dbURL := GetStringEnv("PG_LOCAL_DB_URL", DefaultString("localhost"))

	db, err := sql.Open("pgx", dbURL)

	if err != nil {
		panic(err)
	}

	DB = db
}
