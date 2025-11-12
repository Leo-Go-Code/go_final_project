package db

import (
	"database/sql"
	"fmt"
)

func UpdateDate(id, date *string) error {

	//	1.	Делаем запрос в БД
	res, err := db.Exec(`UPDATE scheduler SET date = :date WHERE id = :id;`,
		sql.Named("date", date),
		sql.Named("id", id))
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
