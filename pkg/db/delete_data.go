package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func DeleteTask(id *string) error {

	//	1.	Запрос на удаление в БД
	res, err := db.Exec(`DELETE FROM scheduler WHERE id = :id;`, sql.Named(`id`, id))
	if err != nil {
		return fmt.Errorf(`ошибка выполнения запроса на удаление задачи из базы данных: %w`, err)
	}

	//	2.	Проверка удаления задачи из БД по id
	deleted, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf(`ошибка проверки удаления строк из БД по соответствующему запросу: %w`, err)
	}
	if deleted == 0 {
		return fmt.Errorf(`при выполнении запроса на удаление задачи из базы данных не было удалено ни одной строки: %w`, err)
	}

	return err
}
