package cache

import (
	"database/sql"
	"github.com/rs/zerolog"
	"task4/server"
	"task4/storage"
)

var db *sql.DB
var log *zerolog.Logger


func InitDB() {
	log = server.NewLogger()
	var err error
	connstring := "user=test password=qwe host= port=5432 database=cache sslmode=disable"

	db, err = sql.Open("postgres", connstring)
	if err != nil {
		return
	}

	p, err  := storage.NewProvider(connstring, log)
	if err != nil {
		return
	}
	db = p.Conn
}
