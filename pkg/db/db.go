package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

// SQL-команды для создания таблицы scheduler и индекса по колонке date
const schema = `CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR(32) NOT NULL DEFAULT '',
	comment TEXT NOT NULL DEFAULT '',
	repeat VARCHAR(128) NOT NULL DEFAULT ''
);
CREATE INDEX date_id ON scheduler (date);`

var db *sql.DB

func Init(infDBfile string) error {
	var err error

	if infDBfile == "" {
		return fmt.Errorf("получена пустая строка infDBfile в функции Init")
	}

	//	0.	Получение имени базы данных и пути к ней
	nameDBfile := strings.TrimLeft(infDBfile, "../pkg/db")                             //	"scheduler.db"
	dirDBfile := strings.TrimLeft(strings.TrimRight(infDBfile, "scheduler.db"), "../") //	"pkg/db/"

	//	1.	Проверка полученной строки
	if nameDBfile == "" { //	должно быть "scheduler.db"
		return fmt.Errorf("название таблицы не указано")
	}

	//	2.	Входим поддиректорию pkg/db
	err = os.Chdir(dirDBfile)
	if err != nil {
		return fmt.Errorf("ошибка входа в директорию базы данных")
	}

	//	3.	Открытие таблицы
	db, err = sql.Open("sqlite", nameDBfile)
	if err != nil {
		return fmt.Errorf("ошибка: таблица не открылась: %v", err)
	}
	defer db.Close()

	//	4.	Проверка существования файла БД
	_, err = os.Stat(nameDBfile)
	if err != nil {
		if os.IsNotExist(err) {
			//	5.	Создание таблицы, если она отсутствует в директории
			_, err = db.Exec(schema)
			if err != nil {
				return fmt.Errorf("ошибка создания таблицы: %v", err)
			}
		} else {
			return fmt.Errorf("ошибка при проверке существования файла базы данных: %v", err)
		}
	}

	//	6.	Выходим обратно в директорию проекта
	err = os.Chdir("../..")
	if err != nil {
		return fmt.Errorf("ошибка выхода в директорию проекта")
	}

	return err
}
