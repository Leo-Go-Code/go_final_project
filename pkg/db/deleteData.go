package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func DeleteData(id string) error {
	//	0.	Открываем файл базы данных
	db, err := sql.Open("sqlite", "pkg/db/scheduler.db")
	if err != nil {
		return fmt.Errorf("ошибка открытия базы данных: %v", err)
	}
	defer db.Close()

	//	1.	Удаляем задачу из БД по id
	_, err = db.Exec("DELETE FROM scheduler WHERE id = ?;", id)
	if err != nil {
		return fmt.Errorf("ошибка запроса на удаление задачи из базы данных: %v", err)
	}

	return err
}
