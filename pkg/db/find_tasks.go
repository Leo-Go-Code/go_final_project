package db

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

func FindDate(searchDate time.Time, limit int) ([]*Task, error) {
	//	0.	Создаём []*Task для записи ближайших задач
	tasks := []*Task{}

	//	1.	Подготавливаем строку для запроса
	searchStr := searchDate.Format("20060102")

	//	2.	Делаем запрос в БД
	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date = :date LIMIT :limit;", sql.Named("date", searchStr), sql.Named("limit", limit))

	//	3.	Проверяем и извлекаем данные из rows
	tasks, err = checkRows(tasks, rows, err)

	//	4.	Отправляем полученный список задач
	return tasks, err
}

func FindStr(search string, limit int) ([]*Task, error) {
	//	0.	Создаём []*Task для записи ближайших задач
	tasks := []*Task{}

	//	1.	Переводим строку в правильный формат
	search = "%" + search + "%"

	//	2.	Делаем запрос в БД
	rows, err := db.Query(`
		SELECT id, date, title, comment, repeat FROM scheduler WHERE (title LIKE :search OR comment LIKE :search) 
		ORDER BY date LIMIT :limit;`,
		sql.Named("search", search),
		sql.Named("limit", limit))

	//	3.	Проверяем и извлекаем данные из rows
	tasks, err = checkRows(tasks, rows, err)

	//	4.	Отправляем полученный список задач
	return tasks, err
}
