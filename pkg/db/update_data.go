package db

import (
	"fmt"
)

func UpdateData(newTask Task) error {

	//	1.	Делаем запрос в БД
	query := `UPDATE scheduler SET date = $1, title = $2, comment = $3, repeat = $4 WHERE id = $5;`
	res, err := DB.Exec(query, newTask.Date, newTask.Title, newTask.Comment, newTask.Repeat, newTask.ID)
	if err != nil {
		return fmt.Errorf("ошибка запроса к базе данных: %v", err)
	}

	//	2.	Определяем кол-во изменёных записей
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("НЕ корректное кол-во изменённых задач в базе данных: %v", err)
	}
	if count == 0 {
		return fmt.Errorf("НЕ корректный id для обновления данных задачи")
	}

	return err
}
