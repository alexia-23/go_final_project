package database

import (
	"fmt"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (d *Database) UpdateTask(task domain.Task) (string, error) {
	query := `
		UPDATE scheduler
		SET
			date    = ?,
			title   = ?,
			comment = ?,
			repeat  = ?
		WHERE id = ?
	`

	result, err := d.DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
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
