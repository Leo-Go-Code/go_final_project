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

func AddTask(task *Task) (int64, error) {
	var id int64

	//	1.	Выполняем запрос в БД: добавляем задачу
	res, err := db.Exec(`INSERT INTO scheduler (date, title, comment, repeat) 
	VALUES (:date, :title, :comment, :repeat);`,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return 0, fmt.Errorf("ошибка запроса к базе данных: %w", err)
	}

	//	2.	Получаем id последней добавленной задачи (проверка err в AddTask)
	id, err = res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("ошибка получения id из базы данных: %w", err)
	}

	return id, err
}
