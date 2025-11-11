package api

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/m/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTaskS(w http.ResponseWriter, r *http.Request) {
	//	0.	Извлекаем строку из запроса и смотрим содержимое
	url := r.RequestURI
	if strings.Contains(url, "search=") {

		//	1.	Получаем строку данных из поискового запроса
		URL := r.RequestURI
		index := strings.Index(URL, "search=")
		search := URL[index+7:]
		if search == "" {
			tasks, err := db.GetDataS(20) //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				writeJSON(w, TasksResp{Tasks: tasks}, fmt.Errorf("в БД не найдены задачи: %v", err))
				return
			}
			writeJSON(w, TasksResp{Tasks: tasks}, err)
			return
		}

		//	2.	Выполняем запрос
		tasks, err := db.GetDataSearch(search, 20) //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			writeJSON(w, TasksResp{Tasks: tasks}, fmt.Errorf("в БД не найдены задачи: %v", err))
			return
		}
		writeJSON(w, TasksResp{Tasks: tasks}, err)
	} else {
		tasks, err := db.GetDataS(20) //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			writeJSON(w, TasksResp{Tasks: tasks}, fmt.Errorf("в БД не найдены задачи: %v", err))
			return
		}
		writeJSON(w, TasksResp{Tasks: tasks}, err)
	}
}
