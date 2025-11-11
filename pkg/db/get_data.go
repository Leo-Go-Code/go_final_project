package db

import (
	"fmt"

	_ "modernc.org/sqlite"
)

func GetData(id string) (Task, error) {
	//	1.	Переменные структуры Task для записи и возвращения информации из ДБ
	var task Task
	var empty Task

	//	2.	Получаем данные по id и проверяем их
	err := DB.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?;", id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return empty, fmt.Errorf(`ошибка выполнения GET запроса в БД: %v`, err)
	}

	//	3.	Возвращаем результат
	return task, err
}
