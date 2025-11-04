package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type Database struct {
	DB *sql.DB
}

// New — конструктор, создаёт подключение и применяет миграции.
func New() (*Database, error) {
	dataSource := os.Getenv("DATA_SOURCE")
	if dataSource == "" {
		return nil, fmt.Errorf("переменная окружения DATA_SOURCE не задана")
	}

	db, err := sql.Open("sqlite", dataSource)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	script, err := os.ReadFile("sql/v1.sql")
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("не удалось прочитать файл миграции: %w", err)
	}

	if _, err := db.Exec(string(script)); err != nil {
		db.Close()
		return nil, fmt.Errorf("ошибка при выполнении миграции: %w", err)
	}

	return &Database{DB: db}, nil
}

// Close — метод для корректного закрытия соединения
func (d *Database) Close() {
	if d.DB != nil {
		_ = d.DB.Close()
	}
}
