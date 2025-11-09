package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddData(task *Task) (int64, error) {
	var id int64

	//	0.	Открываем файл базы данных
	db, err := sql.Open("sqlite", "pkg/db/scheduler.db")
	if err != nil {
		return 0, fmt.Errorf("ошибка открытия базы данных: %v", err)
	}
	defer db.Close()

	//	1.	Выполняем запрос в БД: добавляем задачу
	query := "INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4);"
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, fmt.Errorf("ошибка запроса к базе данных: %v", err)
	}

	//	2.	Получаем id последней добавленной задачи (проверка err в AddTask)
	id, err = res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("ошибка получения id из базы данных: %v", err)
	}

	return id, err
}
