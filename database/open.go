package database

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var dbContext *sqlx.DB

func OpenDb(connectionString string) error {
	db, connectionErr := sqlx.Open("pgx", connectionString)
	if connectionErr != nil {
		log.Fatal(connectionErr)
	}
	dbContext = db
	return nil
}
