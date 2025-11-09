package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func GetData(id string) (Task, error) {
	//	0.	Переменные структуры Task для записи и возвращения информации из ДБ
	var task Task
	var empty Task

	//	1.	Открываем файл базы данных
	db, err := sql.Open("sqlite", "pkg/db/scheduler.db")
	if err != nil {
		return empty, fmt.Errorf("ошибка открытия базы данных: %v", err)
	}
	defer db.Close()

	//	2.	Получаем данные по id и проверяем их
	err = db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?;", id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return empty, fmt.Errorf(`ошибка выполнения GET запроса в БД: %v`, err)
	}
	if task == empty {
		return empty, fmt.Errorf("из БД получен пустой *Task")
	}
	if task.ID == "" || task.Date == "" || task.Title == "" {
		return empty, fmt.Errorf("из строки row были получены НЕ корректные данные (task.ID == '' || &task.Date == '' || &task.Title == '')")
	}

	//	3.	Возвращаем результат
	return task, err
}
