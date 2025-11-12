package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/m/pkg/db"
)

func doneTask(w http.ResponseWriter, r *http.Request) {
	//	0.	Проверка метода Post
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJSON(w, map[string]string{"error": "не верно определён метод запроса для func doneTask()"})
		return
	}

	//	1.	Переменные для формирования ответа
	var task db.Task

	//	2.	Определяем ID задачи
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

	//	3.	Получаем данные задачи
	task, err := db.GetTask(&task.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	//	4.	Обновляем или удаляем дату задачу, если она не треует повторения
	if task.Repeat == "" {
		err = db.DeleteTask(&task.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			writeJSON(w, map[string]string{"error": "ошибка удаления задачи из БД"})
		}
		writeJSON(w, nil)
	} else {
		now := time.Now().Format(FormateDate)
		task.Date, err = taskDate(now, task.Date, task.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, map[string]string{"error": fmt.Sprintf("ошибка определения новой даты: %w", err)})
			return
		}
		err = db.UpdateDate(&task.ID, &task.Date)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJSON(w, map[string]string{"error": fmt.Sprintf("в БД не найдена задача с указанным ID: %w", err)})
			return
		}
		writeJSON(w, err)
	}
}
