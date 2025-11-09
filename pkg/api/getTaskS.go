package api

import (
	"net/http"
	"strings"

	"example.com/m/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTaskS(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	if strings.Contains(url, "search=") {
		tasks, err := db.GetDataSearch(w, r, 20) //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)
		writeJSON(w, TasksResp{Tasks: tasks}, err)
	} else {
		tasks, err := db.GetDataS(20) //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)
		writeJSON(w, TasksResp{Tasks: tasks}, err)
	}
}
