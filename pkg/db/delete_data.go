package db

import (
	"fmt"

	_ "modernc.org/sqlite"
)

func DeleteData(id string) error {

	_, err := DB.Exec("DELETE FROM scheduler WHERE id = ?;", id)
	if err != nil {
		return fmt.Errorf("ошибка запроса на удаление задачи из базы данных: %v", err)
	}

	return err
}
