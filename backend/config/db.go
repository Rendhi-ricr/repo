package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := "postgres://postgres:rendhi123@localhost:5432/repository_db"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB = pool
	log.Println("Database connected")
}


