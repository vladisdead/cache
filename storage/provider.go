package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func NewProvider(connstring string, log *zerolog.Logger) (*Provider, error) {
	p := Provider{
		Log: log,
	}
	var err error

	p.Conn, err = sql.Open("postgres", connstring)
	if err != nil {
		return nil, err
	}

	return &p, p.Conn.Ping()
}
