package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"example.com/m/pkg/db"
)

func doneTask(w http.ResponseWriter, r *http.Request) {
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
			writeJSON(w, empty, err)
			return
		}
	}

	//	2.	Получаем данные задачи
	task, err := db.GetData(task.ID)
	if err != nil {
		writeJSON(w, empty, err)
		return
	}

	//	3.	Обновляем или удаляем дату задачу, если она не треует повторения
	if task.Repeat == "" {
		err = db.DeleteData(task.ID)
		writeJSON(w, empty, err)
	} else {
		now := time.Now().Format(formatDate)
		task.Date, err = taskDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, empty, err)
			return
		}
		err = db.DoneData(task.ID, task.Date)
		writeJSON(w, empty, err)
	}
}
