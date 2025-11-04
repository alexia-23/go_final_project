package db

import (
	"fmt"
	"os"

	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/jmoiron/sqlx"
)

func UpdateTask(task domain.Task) (string, error) {
	dataSource := os.Getenv("DATA_SOURCE")
	db, err := sqlx.Open("sqlite3", dataSource)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к БД: %w", err)
	}
	defer db.Close()

	query := `
		UPDATE scheduler
		SET
			date    = :date,
			title   = :title,
			comment = :comment,
			repeat  = :repeat
		WHERE id = :id
	`

	result, err := db.NamedExec(query, task)
	if err != nil {
		return "", fmt.Errorf("ошибка при обновлении задачи: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("не удалось определить количество изменённых строк: %w", err)
	}
	if rows == 0 {
		return "", fmt.Errorf("задача с id=%d не найдена", task.ID)
	}

	return fmt.Sprintf("%d", task.ID), nil
}
