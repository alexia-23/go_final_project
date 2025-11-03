package db

import (
	"database/sql"
	"fmt"
	"github.com/alexia-23/go_final_project/pkg/domain"
	"github.com/jmoiron/sqlx"
	"os"
)

func SelectTaskById(id int) (domain.Task, error) {
	var task domain.Task

	dataSource := os.Getenv("DATA_SOURCE")
	db, err := sqlx.Open("sqlite3", dataSource)
	if err != nil {
		return task, fmt.Errorf("ошибка подключения к БД: %w", err)
	}
	defer db.Close()

	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE id = :id
	`

	rows, err := db.NamedQuery(query, map[string]any{"id": id})
	if err != nil {
		return task, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return task, sql.ErrNoRows
	}

	err = rows.StructScan(&task)
	if err != nil {
		return task, fmt.Errorf("ошибка при сканировании результата: %w", err)
	}

	return task, nil
}
