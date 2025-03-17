package pg_connector

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	dbpool *pgxpool.Pool
	once   sync.Once
)

// InitDB инициализирует подключение к БД
func InitDB() *pgxpool.Pool {
	once.Do(func() {
		_ = godotenv.Load() // Загружаем .env

		// Формируем строку подключения вручную
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

		if dsn == "postgres://::::" {
			log.Fatal("DATABASE_URL не задан")
		}

		var err error
		dbpool, err = pgxpool.New(context.Background(), dsn)
		if err != nil {
			log.Fatalf("Ошибка подключения к БД: %v", err)
		}
	})
	return dbpool
}

// CloseDB закрывает соединение с БД
func CloseDB() {
	if dbpool != nil {
		dbpool.Close()
	}
}
