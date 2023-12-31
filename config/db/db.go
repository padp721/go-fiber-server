package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var Session *sql.DB

type credential struct {
	DB_CLIENT   string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Connect() (*sql.DB, error) {
	var err error

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error loading DB_PORT from environment variable!")
	}

	dbCredential := credential{
		DB_CLIENT:   os.Getenv("DB_CLIENT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     DB_PORT,
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	Session, err = sql.Open(
		dbCredential.DB_CLIENT, fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbCredential.DB_HOST,
			dbCredential.DB_PORT,
			dbCredential.DB_USER,
			dbCredential.DB_PASSWORD,
			dbCredential.DB_NAME,
		),
	)

	if err != nil {
		return Session, err
	}
	if err := Session.Ping(); err != nil {
		return Session, err
	}
	return Session, nil
}
