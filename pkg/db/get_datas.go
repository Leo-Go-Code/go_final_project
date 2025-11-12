package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetTasks(limit int) ([]*Task, error) {
	//	1.	Создаём []*Task для записи ближайших задач
	tasks := []*Task{}

	//	2.	Делаем запрос в БД
	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler LIMIT :limit;", sql.Named("limit", limit))

	//	3.	Проверяем и извлекаем данные из rows
	tasks, err = checkRows(tasks, rows, err)

	//	4.	Отправляем полученный список задач
	return tasks, err
}
