package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

func GetDataSearch(search string, limit int) ([]*Task, error) {
	//	1.	Создаём []*Task для записи ближайших задач
	tasks := []*Task{}
	empty := []*Task{}

	//	2.	Проверяем строку на наличие цифр
	var numbers int
	for _, r := range search {
		if r >= '0' && r <= '9' {
			numbers++
		}
	}

	//	3.	Обрабатываем запрос
	if numbers == 8 { //	обрабатываем запрос с датой

		//	3.1.1.	Переводим строку в правильный формат
		datePieces := strings.Split(search, ".")
		if len(datePieces) != 3 {
			return empty, fmt.Errorf("в поисковой строке НЕ корректно задана дата")
		}
		date := datePieces[2] + datePieces[1] + datePieces[0]

		//	3.1.2.	Делаем запрос в БД
		rows, err := DB.Query("SELECT * FROM scheduler WHERE date = :date LIMIT :limit;", sql.Named("date", date), sql.Named("limit", limit))
		if err != nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: %v", err)
		}
		if rows == nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: БД вернула пустую переменную rows")
		}
		if err = rows.Err(); err != nil {
			return empty, fmt.Errorf("ошибка чтения данных, ошибка в строке rows.Err(): %v", err)
		}
		defer rows.Close()

		//	3.1.3.	Сканирование и добавление полученных данных методами Next() и Scan() в ответ (+проверки)
		for rows.Next() {
			var task Task

			err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
			if err != nil {
				return empty, fmt.Errorf("ошибка чтения данных из строки row: %v", err)
			}

			if task.ID == "" || task.Date == "" || task.Title == "" {
				return empty, fmt.Errorf("из строки row были получены НЕ корректные данные (task.ID == '' || &task.Date == '' || &task.Title == ''): %v", err)
			}
			tasks = append(tasks, &task)
		}

		//	3.1.4.	Проверка заполненности списка возвращаемых задач
		if len(tasks) < 1 {
			return empty, err
		}

		//	3.1.5.	Отправляем полученный список задач
		return tasks, err

	} else { //	обрабатываем запрос со строкой
		//	3.2.1.	Переводим строку в правильный формат
		search = "%" + search + "%"

		//	3.2.2.	Делаем запрос в БД
		rows, err := DB.Query(`
			SELECT * FROM scheduler WHERE (title LIKE :search OR comment LIKE :search) 
			ORDER BY date LIMIT :limit;`,
			sql.Named("search", search),
			sql.Named("limit", limit))

		if err != nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: %v", err)
		}
		if rows == nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: БД вернула пустую переменную rows")
		}

		defer rows.Close()

		//	3.2.3.	Сканирование и добавление полученных данных методами Next() и Scan() в ответ (+проверки)
		for rows.Next() {
			var task Task

			err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
			if err != nil {
				return empty, fmt.Errorf("ошибка чтения данных из строки row: %v", err)
			}

			if task.ID == "" || task.Date == "" || task.Title == "" {
				return empty, fmt.Errorf("из строки row были получены НЕ корректные данные (task.ID == '' || &task.Date == '' || &task.Title == ''): %v", err)
			}
			tasks = append(tasks, &task)
		}

		//	3.2.4.	Отправляем полученный список задач
		return tasks, err
	}
}
