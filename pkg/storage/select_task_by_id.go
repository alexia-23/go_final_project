package storage

import (
	"database/sql"
	"fmt"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (d *Storage) SelectTaskById(id int) (domain.Task, error) {
	var task domain.Task

	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE id = ?
	`

	row := d.db.QueryRow(query, id)

	err := row.Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, fmt.Errorf("задача с id=%d не найдена", id)
		}
		return task, fmt.Errorf("ошибка при выборке задачи: %w", err)
	}

	return task, nil
}
