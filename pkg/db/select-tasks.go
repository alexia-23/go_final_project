package db

import (
	"fmt"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/jmoiron/sqlx"
	"os"
)

func SelectTasks() ([]domain.Task, error) {
	tasks := []domain.Task{}
	dataSource := os.Getenv("DATA_SOURCE")

	db, err := sqlx.Open("sqlite3", dataSource)
	if err != nil {
		return tasks, fmt.Errorf("ошибка подключения к БД: %w", err)
	}
	defer db.Close()

	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY date
	`

	err = db.Select(&tasks, query)
	if err != nil {
		return tasks, fmt.Errorf("ошибка при выборке задач: %w", err)
	}

	return tasks, nil
}
