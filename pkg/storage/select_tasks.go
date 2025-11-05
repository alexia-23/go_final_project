package storage

import (
	"fmt"

	"github.com/alexia-23/go_final_project/pkg/domain"
)

func (d *Storage) SelectTasks() ([]domain.Task, error) {
	tasks := []domain.Task{}

	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY date
	`

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выборке задач: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при чтении результатов: %w", err)
	}

	return tasks, nil
}
