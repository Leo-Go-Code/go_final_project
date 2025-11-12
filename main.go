package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"example.com/m/pkg/api"
	"example.com/m/pkg/db"
)

func main() {
	//	0.	Код для проверки наличия переменных окружения: TODO_PORT, TODO_DBFILE
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	//	1.	Получаем значение переменной окружения: расположение базы данных
	infDBfile := os.Getenv("TODO_DBFILE")
	infDBfile = strings.ReplaceAll(infDBfile, "../", "")

	//	2.	Открываем (создаём) базу данных
	db, err := db.InitDB(infDBfile)
	if err != nil {
		err = fmt.Errorf("ошибка открытия базы данных: %v", err)
		log.Fatal(err)
	}
	defer db.Close()

	//	3.	Получаем значение переменной окружения: порт сервера
	port := os.Getenv("TODO_PORT")

	//	4.	Определяем функции для работы с URL
	api.Init() // Функция-обработчик для "/api/"

	//	5.	Обработка запроса на возврат формы пользователю
	http.Handle("/", http.FileServer(http.Dir("./web")))

	//	6.	Для удобства выводим в консоль сообщения о нашем сервере
	fmt.Println("Сервер запущен на http://localhost:7540")

	//	7.	Запускаем сервер
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
