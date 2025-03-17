package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbpool *pgxpool.Pool
	once   sync.Once
)

// InitDB инициализирует подключение к БД
func InitDB() *pgxpool.Pool {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL") // Берём строку подключения из ENV
		if dsn == "" {
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
