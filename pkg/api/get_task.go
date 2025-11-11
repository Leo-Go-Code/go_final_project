package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/m/pkg/db"
)

func getTask(w http.ResponseWriter, r *http.Request) {
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

	//	2.	Получаем из БД структуру задачи по ID
	task, err = db.GetData(task.ID)
	// fmt.Println(task)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, empty, fmt.Errorf("в БД не найдена задача с указанным ID: %v", err))
		return
	}
	writeJSON(w, task, err)
}
