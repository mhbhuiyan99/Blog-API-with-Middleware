package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=postgres password=101010 dbname=blogAPI sslmode=disable"
}

func NewDBConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()

	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
