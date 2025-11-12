package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"example.com/m/pkg/db"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {
	//	0.	Переменные для формирования ответа
	var task db.Task

	//	1.	Определяем ID задачи
	url := r.RequestURI
	if strings.Contains(url, "id=") {
		index := strings.Index(url, "id=")
		task.ID = url[index+3:]
	} else {
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJSON(w, map[string]string{"error": fmt.Sprintf("ошибка декодирования запроса: %w", err)})
			return
		}
	}

	//	2.	Проверка ID
	if task.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": "в запросе НЕ был получен ID"})
		return
	}
	_, err := strconv.Atoi(task.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": "получено не корректное значение ID"})
		return
	}

	//	3.	Удаляем из БД задачу по ID
	err = db.DeleteTask(&task.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, map[string]string{"error": fmt.Sprintf("в БД не найдена задача с указанным ID: %w", err)})
		return
	}
	writeJSON(w, nil)
}
