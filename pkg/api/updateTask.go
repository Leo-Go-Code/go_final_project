package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/m/pkg/db"
)

func updateTask(w http.ResponseWriter, r *http.Request) {
	//	0.	Переменные для формирования ответа
	var task db.Task
	var empty struct{}

	//	1.	Извлекаем URL из запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, empty, err)
		return
	}

	//	2. Проверка task.Title != ""
	if task.Title == "" {
		writeJSON(w, empty, fmt.Errorf("поле 'Задача' обязательно должно быть заполнено"))
		return
	}

	//	3.  Проверка task на корректность (можно спустить только нужные строки, а не весь task *db.Task)
	err = checkDateTask(&task)
	if err != nil {
		writeJSON(w, empty, fmt.Errorf("поле 'Дата' заполнено НЕ корректно: %v", err))
		return
	}

	//	4.	Выполняем запрос к базе данных
	err = db.UpdateData(task)
	writeJSON(w, task, err)
}
