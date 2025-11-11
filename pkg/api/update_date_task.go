package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/m/pkg/db"
)

func updateDateTask(w http.ResponseWriter, r *http.Request) {
	//	0.	Переменные для формирования ответа
	var task db.Task
	var empty struct{}

	//	1.	Определяем ID задачи
	url := r.RequestURI
	if strings.Contains(url, "id=") {
		index := strings.Index(url, "id=")
		task.ID = url[index+3:]
	} else {
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJSON(w, empty, err)
			return
		}
	}

	//	2.	Получаем данные задачи
	task, err := db.GetData(task.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, empty, err)
		return
	}

	//	3.	Обновляем или удаляем дату задачу, если она не треует повторения
	if task.Repeat == "" {
		err = db.DeleteData(task.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			writeJSON(w, empty, err)
		}
		writeJSON(w, empty, err)
	} else {
		now := time.Now().Format(FORMAT_DATE)
		task.Date, err = taskDate(now, task.Date, task.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, empty, err)
			return
		}
		err = db.UpdateDate(task.ID, task.Date)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJSON(w, empty, fmt.Errorf("в БД не найдена задача с указанным ID: %v", err))
			return
		}
		writeJSON(w, empty, err)
	}
}
