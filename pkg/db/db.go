package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// SQL-команды для создания таблицы scheduler и индекса по колонке date
const SCHEMA = `CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR(32) NOT NULL DEFAULT '',
	comment TEXT NOT NULL DEFAULT '',
	repeat VARCHAR(128) NOT NULL DEFAULT ''
);
CREATE INDEX IF NOT EXISTS date_id ON scheduler (date);`

var db *sql.DB

func InitDB(infDBfile string) (*sql.DB, error) {
	var err error

	//	1.	Проверка полученной строки
	if infDBfile == "" { //	должно быть "scheduler.db"
		return nil, fmt.Errorf("название таблицы не указано")
	}

	//	2.	Открытие таблицы
	db, err = sql.Open("sqlite", infDBfile)
	if err != nil {
		return nil, fmt.Errorf("ошибка: таблица не открылась: %w", err)
	}

	//	3.	Создание таблицы, если она отсутствует в директории
	_, err = db.Exec(SCHEMA)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания таблицы: %w", err)
	}

	return db, err
}
