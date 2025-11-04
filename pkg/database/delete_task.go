package database

import (
	"fmt"
)

func (d *Database) DeleteTaskById(id int) (string, error) {
	query := `DELETE FROM scheduler WHERE id = ?`

	result, err := d.DB.Exec(query, id)
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
