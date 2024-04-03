package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Type     string
}

func ConnectDB(config DBConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Paris",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open(config.Type, connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	return db
}
