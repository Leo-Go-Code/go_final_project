package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/m/pkg/db"
)

func addTask(w http.ResponseWriter, r *http.Request) {
	//	0.	Объявляем db.Task для запонения и вывода
	var task db.Task

	//	1. Десериализуем JSON из запроса в переменную task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": fmt.Sprintf("ошибка декодирования запроса: %w", err)})
		return
	}

	//	2. Проверка task.Title != ""
	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "поле 'Задача' обязательно должно быть заполнено"})
		return
	}

	//	3.  Проверка task на корректность (можно спустить только нужные строки, а не весь task *db.Task)
	err = checkDateTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": fmt.Sprintf("поле 'Дата' заполнено НЕ корректно: %w", err)})
		return
	}

	//	4. Добавляем задачу в БД
	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": fmt.Sprintf("ошибка при добавлении записи в базу данных: %w", err)})
		return
	}

	//	5. Проверка id
	if id == 0 {
		writeJSON(w, map[string]string{"error": "получено НЕ корректное значение id"})
		return
	}

	//	6. 	Возвращаем id добавленной в БД задачи в виде JSON
	resp := map[string]interface{}{
		"id": id,
	}
	writeJSON(w, resp)
}

func checkDateTask(task *db.Task) error {
	//	0. Для пустого task.Date присвоим текущее время
	nowTime := time.Now()
	nowStr := nowTime.Format(FormateDate)
	if task.Date == "" {
		task.Date = nowStr
		return nil
	}

	//	1. Проверяем данные
	err = checkDstartNow(nowStr)
	if err != nil {
		err = fmt.Errorf("ошибка при проверке корректности строки now: %w", err)
		return err
	}
	err = checkDstartNow(task.Date)
	if err != nil {
		err = fmt.Errorf("ошибка при проверке корректности строки dstart: %w", err)
		return err
	}
	if task.Repeat != "" {
		err = checkRepeat(task.Repeat)
		if err != nil {
			err = fmt.Errorf("ошибка при проверке корректности строки repeat: %w", err)
			return err
		}
	}

	//	2.	Если дата меньше сегодняшнего дня и повтор НЕ задан
	if task.Date < nowStr {
		if task.Repeat == "" {
			task.Date = nowStr
		} else {
			//	3.	Если дата меньше сегодняшнего дня и повтор ЗАДАН
			next, err := taskDate(nowStr, task.Date, task.Repeat)
			if err != nil {
				task.Date = nowStr
			} else {
				task.Date = next
			}
		}
	}
	return nil
}
