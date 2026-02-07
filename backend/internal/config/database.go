package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB adalah koneksi database global
var DB *pgxpool.Pool

// ConnectDB membuat koneksi ke database PostgreSQL
func ConnectDB() {
	// Ambil DSN dari environment variable, atau gunakan default
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:rendhi123@localhost:5432/repository_db"
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// Test koneksi
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Database tidak dapat dijangkau:", err)
	}

	DB = pool
	log.Println("✅ Database connected successfully")
}

// CloseDB menutup koneksi database
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
