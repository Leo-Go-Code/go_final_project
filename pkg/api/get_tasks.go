package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"example.com/m/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

const limit = 20 //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)

func getTasks(w http.ResponseWriter, r *http.Request) {
	//	0.	Проверяем метод Get
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJSON(w, map[string]string{"error": "не верно определён метод запроса для func doneTask()"})
		return
	}

	//	1.	Вводим переменную для сбора и вывода результатов поиска
	tasks := []*db.Task{}
	var err error

	//	2.	Извлекаем строку из запроса и проверяем содержимое
	url := r.RequestURI
	if strings.Contains(url, "search=") {

		//	2.1.	Получаем строку данных из поискового запроса
		search := r.URL.Query().Get("search")
		if search == "" {
			tasks, err = db.GetTasks(limit)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				writeJSON(w, map[string]string{"error": fmt.Sprintf("в БД не найдены задачи: %w", err)})
				return
			}
			writeJSON(w, TasksResp{Tasks: tasks})
			return
		}

		//	2.2.	Проверяем строку на формат даты
		searchDate, errDate := time.Parse("02.01.2006", search)

		//	2.2.1.	Поиск по дате
		if errDate == nil {
			tasks, err = db.FindDate(searchDate, limit)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				writeJSON(w, map[string]string{"error": fmt.Sprintf("ошибка поиска задач по дате: %w", err)})
				return
			}

			//	2.2.2.	Поиск по строке
		} else {
			tasks, err = db.FindStr(search, limit)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				writeJSON(w, map[string]string{"error": fmt.Sprintf("ошибка поиска задач по строке: %w", err)})
				return
			}
		}

		writeJSON(w, TasksResp{Tasks: tasks})

		//	3.	Поиск записей без запроса из поисковой строки
	} else {
		tasks, err = db.GetTasks(limit) //	максимальное кол-во возвращаемых записей 10..50 (на выбор разработчика)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			writeJSON(w, fmt.Errorf("в БД не найдены задачи: %w", err))
			return
		}
		writeJSON(w, TasksResp{Tasks: tasks})
	}
}
