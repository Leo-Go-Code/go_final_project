package db

import (
	"fmt"
)

func UpdateDate(id, date string) error {

	//	1.	Делаем запрос в БД
	query := `UPDATE scheduler SET date = $1 WHERE id = $2;`
	res, err := DB.Exec(query, date, id)
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
