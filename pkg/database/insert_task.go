package database

import (
	"fmt"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (d *Database) InsertTask(task domain.Task) (string, error) {
	query := `
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`

	result, err := d.DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return "", fmt.Errorf("ошибка вставки: %w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("ошибка получения ID: %w", err)
	}

	return fmt.Sprintf("%d", lastID), nil
}
