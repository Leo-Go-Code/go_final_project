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
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, empty, err)
		return
	}

	//	2. Проверка task.ID == ""
	if task.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, empty, fmt.Errorf("поле 'ID' обязательно должно быть заполнено"))
		return
	}

	//	3. Проверка task.Title == ""
	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, empty, fmt.Errorf("поле 'Задача' обязательно должно быть заполнено"))
		return
	}

	//	4.  Проверка task на корректность (можно спустить только нужные строки, а не весь task *db.Task)
	err = checkDateTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, empty, fmt.Errorf("поле 'Дата' заполнено НЕ корректно: %v", err))
		return
	}

	//	5.	Выполняем запрос к базе данных
	err = db.UpdateData(task)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, empty, fmt.Errorf("в БД не найдена задача с указанным ID: %v", err))
		return
	}
	writeJSON(w, task, err)
}
