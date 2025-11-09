package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	_ "modernc.org/sqlite"
)

func GetDataSearch(w http.ResponseWriter, r *http.Request, limit int) ([]*Task, error) {
	//	0.	Создаём []*Task для записи ближайших задач
	tasks := []*Task{}
	empty := []*Task{}

	//	1.	Открываем файл базы данных
	db, err := sql.Open("sqlite", "pkg/db/scheduler.db")
	if err != nil {
		return empty, fmt.Errorf("ошибка открытия базы данных: %v", err)
	}
	defer db.Close()

	//	2. Получаем строку данных из поискового запроса
	URL := r.RequestURI
	index := strings.Index(URL, "search=")
	searchURL := URL[index+7:]
	if searchURL == "" {
		return GetDataS(limit)
	}

	search, err := url.QueryUnescape(searchURL)
	if err != nil {
		return empty, fmt.Errorf("ошибка конвертации строки из searchURL в search: %v", err)
	}

	//	3.	Проверяем строку на наличие цифр
	var numbers int
	for _, r := range search {
		if r >= '0' && r <= '9' {
			numbers++
		}
	}

	//	4.	Обрабатываем запрос
	if numbers == 8 { //	обрабатываем запрос с датой

		//	4.2.1.	Переводим строку в правильный формат
		datePieces := strings.Split(search, ".")
		if len(datePieces) != 3 {
			return empty, fmt.Errorf("в поисковой строке НЕ корректно задана дата")
		}
		date := datePieces[2] + datePieces[1] + datePieces[0]

		//	4.1.2.	Делаем запрос в БД
		rows, err := db.Query("SELECT * FROM scheduler WHERE date = :date LIMIT :limit;", sql.Named("date", date), sql.Named("limit", limit))
		if err != nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: %v", err)
		}
		defer rows.Close()

		//	4.1.3.	Сканирование и добавление полученных данных методами Next() и Scan() в ответ (+проверки)
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

		//	4.1.4.	Проверка заполненности списка возвращаемых задач
		if len(tasks) < 1 {
			return empty, err
		}

		//	4.1.5.	Отправляем полученный список задач
		return tasks, err

	} else { //	обрабатываем запрос со строкой
		//	4.2.0.	Подготовим строку для запроса: первая буква заглавная, а остальные строчные
		search1 := strings.ToLower(search)
		runeSearch := []rune(search1)
		runeSearch[0] = unicode.ToUpper(runeSearch[0])
		search1 = string(runeSearch)

		//	4.2.1.	Переводим строку в правильный формат
		search = "%" + search + "%"
		searchUpper := "%" + strings.ToUpper(search) + "%"
		searchLower := "%" + strings.ToLower(search) + "%"
		searchUpper1 := "%" + search1 + "%"

		//	4.2.2.	Делаем запрос в БД
		rows, err := db.Query(`
			SELECT * FROM scheduler WHERE (title LIKE :search OR comment LIKE :search) 
			OR (title LIKE :searchUpper OR comment LIKE :searchUpper) 
			OR (title LIKE :searchLower OR comment LIKE :searchLower) 
			OR (title LIKE :searchUpper1 OR comment LIKE :searchUpper1) 
			ORDER BY date LIMIT :limit;`,
			sql.Named("search", search),
			sql.Named("searchUpper", searchUpper),
			sql.Named("searchUpper1", searchUpper1),
			sql.Named("searchLower", searchLower),
			sql.Named("limit", limit))

		if err != nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: %v", err)
		}
		if rows == nil {
			return empty, fmt.Errorf("ошибка запроса к базе данных: БД вернула пустую переменную rows")
		}

		defer rows.Close()

		//	4.2.3.	Сканирование и добавление полученных данных методами Next() и Scan() в ответ (+проверки)
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

		//	4.2.4.	Проверка заполненности списка возвращаемых задач
		if len(tasks) < 1 {
			return empty, err
		}

		//	4.2.5.	Отправляем полученный список задач
		return tasks, err
	}
}
