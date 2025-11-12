package db

import (
	"database/sql"
	"fmt"
)

func UpdateTask(newTask *Task) error {

	//	1.	Делаем запрос в БД
	res, err := db.Exec(`UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id;`,
		sql.Named("date", newTask.Date),
		sql.Named("title", newTask.Title),
		sql.Named("comment", newTask.Comment),
		sql.Named("repeat", newTask.Repeat),
		sql.Named("id", newTask.ID))
	if err != nil {
		return fmt.Errorf("ошибка запроса к базе данных: %w", err)
	}

	//	2.	Определяем кол-во изменёных записей
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("НЕ корректное кол-во изменённых задач в базе данных: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("НЕ корректный id для обновления данных задачи")
	}

	return err
}
