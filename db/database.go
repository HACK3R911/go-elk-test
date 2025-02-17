package db

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog"
)

type Database struct {
	Conn   *sql.DB
	Logger zerolog.Logger
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Logger   zerolog.Logger
}

// DB connection
func Init(cfg Config) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	db.Logger = cfg.Logger
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}
