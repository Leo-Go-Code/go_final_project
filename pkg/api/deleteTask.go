package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/m/pkg/db"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {
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

	//	2.	Проверка ID
	if task.ID == "" {
		writeJSON(w, empty, fmt.Errorf("в запросе НЕ был получен ID"))
		return
	}
	numbers := "0123456789"
	var n int
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(task.ID); j++ {
			if task.ID[j] == numbers[i] {
				n++
			}
		}
	}
	if len(task.ID) != n {
		writeJSON(w, empty, fmt.Errorf("получено не корректное значение ID"))
		return
	}

	//	3.	Удаляем из БД задачу по ID
	err = db.DeleteData(task.ID)
	writeJSON(w, empty, err)
}
