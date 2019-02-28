package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DBInit() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/godb")
	if err != nil {
		log.Panic("failed to connect to database", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
