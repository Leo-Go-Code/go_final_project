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
	//	код для проверки наличия переменных окружения: TODO_PORT, TODO_DBFILE
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	//	получаем значение переменной окружения: расположение базы данных
	infDBfile := os.Getenv("TODO_DBFILE")
	infDBfile = strings.ReplaceAll(infDBfile, "../", "")

	//	открываем (создаём) базу данных
	DB, err := db.InitDB(infDBfile)
	if err != nil {
		err = fmt.Errorf("ошибка открытия базы данных: %v", err)
		log.Fatal(err)
	}
	defer DB.Close()

	//	получаем значение переменной окружения: порт сервера
	port := os.Getenv("TODO_PORT")

	//	определяем функции для работы с URL
	api.Init() // Функция-обработчик для "/api/"

	//	обработка запроса на возврат формы пользователю
	http.Handle("/", http.FileServer(http.Dir("./web")))

	//	для удобства выводим в консоль сообщения о нашем сервере
	fmt.Println("Сервер запущен на http://localhost:7540")

	//	запускаем сервер
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
