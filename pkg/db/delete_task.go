package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func DeleteTaskById(id int) (string, error) {
	dataSource := os.Getenv("DATA_SOURCE")
	db, err := sqlx.Open("sqlite3", dataSource)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к БД: %w", err)
	}
	defer db.Close()

	query := `DELETE FROM scheduler WHERE id = :id`

	result, err := db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return "", fmt.Errorf("ошибка при удалении задачи: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("не удалось определить количество удалённых строк: %w", err)
	}
	if rows == 0 {
		return "", fmt.Errorf("задача с id=%d не найдена", id)
	}

	return fmt.Sprintf("%d", id), nil
}
