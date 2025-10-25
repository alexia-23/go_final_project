package db

import (
	"fmt"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/jmoiron/sqlx"
	"os"
)

func InsertTask(task domain.TaskCreatePayload) (string, error) {
	dataSource := os.Getenv("DATA_SOURCE")

	db, err := sqlx.Open("sqlite3", dataSource)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к БД: %w", err)
	}
	defer db.Close()

	query := `
		INSERT INTO tasks (date, title, comment, repeat)
		VALUES (:date, :title, :comment, :repeat)
	`

	result, err := db.NamedExec(query, task)
	if err != nil {
		return "", fmt.Errorf("ошибка вставки: %w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("ошибка получения ID: %w", err)
	}

	return fmt.Sprintf("%d", lastID), nil
}
