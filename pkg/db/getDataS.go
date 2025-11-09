package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func GetDataS(limit int) ([]*Task, error) {
	//	0.	Создаём []*Task для записи ближайших задач
	tasks := []*Task{}
	empty := []*Task{}

	//	1.	Открываем файл базы данных
	db, err := sql.Open("sqlite", "pkg/db/scheduler.db")
	if err != nil {
		return empty, fmt.Errorf("ошибка открытия базы данных: %v", err)
	}
	defer db.Close()

	//	2.	Делаем запрос в БД
	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler LIMIT ?;", limit)
	if err != nil {
		return empty, fmt.Errorf("ошибка запроса к базе данных: %v", err)
	}
	if rows == nil {
		return empty, fmt.Errorf("ошибка запроса к базе данных: БД вернула пустую переменную rows")
	}
	defer rows.Close()

	//	3.	Сканирование и добавление полученных данных методами Next() и Scan() в ответ (+проверки)
	for rows.Next() {
		var task Task

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return empty, fmt.Errorf("ошибка чтения данных из строки row: %v", err)
		}

		if task.ID == "" || task.Date == "" || task.Title == "" {
			return empty, fmt.Errorf("из строки row были получены НЕ корректные данные (task.ID == '' || &task.Date == '' || &task.Title == ''): %v", err)
		}
		tasks = append(tasks, &task)
	}

	//	4.	Проверка заполненности списка возвращаемых задач
	if len(tasks) < 1 {
		return empty, err
	}

	//	5.	Отправляем полученный список задач
	return tasks, err
}
