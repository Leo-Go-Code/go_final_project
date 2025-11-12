package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func checkRows(tasks []*Task, rows *sql.Rows, err error) ([]*Task, error) {
	//	1.	Проверка rows
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к базе данных: %w", err)
	}
	if rows == nil {
		return nil, fmt.Errorf("ошибка запроса к базе данных: БД вернула пустую переменную rows")
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения данных, ошибка в строке rows.Err(): %w", err)
	}
	defer rows.Close()

	//	2.	Сканирование и добавление полученных данных методами Next() и Scan() в ответ
	for rows.Next() {
		var task Task

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения данных из строки row: %w", err)
		}

		if err = rows.Err(); err != nil {
			return nil, fmt.Errorf("ошибка чтения данных, ошибка в строке rows.Err(): %w", err)
		}

		if task.ID == "" || task.Date == "" || task.Title == "" {
			return nil, fmt.Errorf("из строки row были получены НЕ корректные данные (task.ID == '' || &task.Date == '' || &task.Title == ''): %w", err)
		}
		tasks = append(tasks, &task)
	}

	//	3.	Отправляем полученный список задач
	return tasks, err
}
