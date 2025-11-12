package api

import (
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTask(w, r) //	см. pkg/api/addTask.go
	case http.MethodGet:
		getTask(w, r) //	см. pkg/api/getTask.go
	case http.MethodPut:
		updateTask(w, r) //	см. pkg/api/updateTask.go
	case http.MethodDelete:
		deleteTask(w, r) //	см. pkg/api/deleteTask.go
	default:
		http.Error(w, "ошибка: сервер не обрабатывает данный тип запроса", http.StatusInternalServerError)
	}
}
