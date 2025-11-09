package api

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDate)  // Обработчик для "/api/nextdate" (GET) см. pkg/api/nextDate.go
	http.HandleFunc("/api/task", taskHandler)   // Обработчик для "/api/task" (GET, PUT, DELETE, POST) см. pkg/api/taskHandler.go
	http.HandleFunc("/api/task/done", doneTask) // Обработчик для "/api/task" (POST) см. pkg/api/doneTask.go
	// http.HandleFunc("/api/tasks?search=", getTaskS) // Обработчик для "/api/tasks" (GET) см. pkg/api/getTaskS.go
	http.HandleFunc("/api/tasks", getTaskS) // Обработчик для "/api/tasks" (GET) см. pkg/api/getTaskS.go
}
