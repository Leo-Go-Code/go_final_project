package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"example.com/m/pkg/api"
	"example.com/m/pkg/db"
)

func handleHTML(w http.ResponseWriter, r *http.Request) {
	//	0.	Обрезаем URL и определяем имя файла для возвращения
	dwnFile := r.URL.Path[len("/"):]
	if dwnFile == "" {
		dwnFile = "index.html"
	}

	//	1.	Открываем директорию необходимого файла
	root, err := os.OpenRoot("web")
	if err != nil {
		http.Error(w, "внутренняя ошибка открытия директории файла", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	//	2.	Находим файл и считываем из него данные
	file, err := os.ReadFile(filepath.Join("./web/", dwnFile))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("по данному URL файл %s НЕ найден\n", dwnFile)
			return
		}
		log.Fatal(err)
	}

	//	3.	Выдаём ответ
	//	выставляем w.Header()
	w.Header().Set("Content-Type", "text/html")

	//	возвращаем содержимое файла в браузер
	w.Write(file)
}

func handleJS(w http.ResponseWriter, r *http.Request) {
	//	0.	Обрезаем URL и определяем имя файла для возвращения
	dwnFile := r.URL.Path[len("/js/"):]
	if dwnFile == "" {
		http.Error(w, "имя файла НЕ указано", http.StatusBadRequest)
		return
	}

	//	1.	Открываем директорию необходимого файла
	root, err := os.OpenRoot("web/js")
	if err != nil {
		http.Error(w, "внутренняя ошибка открытия директории файла", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	//	2.	Находим файл и считываем из него данные
	file, err := os.ReadFile(filepath.Join("./web/js", dwnFile))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("по данному URL файл %s НЕ найден\n", dwnFile)
			return
		}
		log.Fatal(err)
	}

	//	3.	Выдаём ответ
	//	выставляем w.Header()
	w.Header().Set("Content-Type", "text/javascript")

	//	возвращаем содержимое файла в браузер
	w.Write(file)
}

func handleCSS(w http.ResponseWriter, r *http.Request) {
	//	0.	Обрезаем URL и определяем имя файла для возвращения
	dwnFile := r.URL.Path[len("/css/"):]
	if dwnFile == "" {
		http.Error(w, "имя файла НЕ указано", http.StatusBadRequest)
		return
	}

	//	1.	Открываем директорию необходимого файла
	root, err := os.OpenRoot("web/css")
	if err != nil {
		http.Error(w, "внутренняя ошибка открытия директории файла", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	//	2.	Находим файл и считываем из него данные
	file, err := os.ReadFile(filepath.Join("./web/css", dwnFile))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("по данному URL файл %s НЕ найден\n", dwnFile)
			return
		}
		log.Fatal(err)
	}

	//	3.	Выдаём ответ
	//	выставляем w.Header()
	w.Header().Set("Content-Type", "text/css")

	//	возвращаем содержимое файла в браузер
	w.Write(file)
}

func handleIco(w http.ResponseWriter, r *http.Request) {
	//	0.	Обрезаем URL и определяем имя файла для возвращения
	dwnFile := r.URL.Path[len("/"):]
	if dwnFile == "" {
		http.Error(w, "имя файла НЕ указано", http.StatusBadRequest)
		return
	}

	//	1.	Открываем директорию необходимого файла
	root, err := os.OpenRoot("web")
	if err != nil {
		http.Error(w, "внутренняя ошибка открытия директории файла", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	//	2.	Находим файл и считываем из него данные
	file, err := os.ReadFile(filepath.Join("./web/", dwnFile))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("по данному URL файл %s НЕ найден\n", dwnFile)
			return
		}
		log.Fatal(err)
	}

	//	3.	Выдаём ответ
	//	выставляем w.Header()
	w.Header().Set("Content-Type", "image/ico")

	//	возвращаем содержимое файла в браузер
	w.Write(file)
}

func main() {
	//	код для проверки наличия переменных окружения: TODO_PORT, TODO_DBFILE
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	//	получаем значение переменной окружения: расположение базы данных
	infDBfile := os.Getenv("TODO_DBFILE")

	//	открываем (создаём) базу данных
	err := db.Init(infDBfile)
	if err != nil {
		err = fmt.Errorf("ошибка открытия базы данных: %v", err)
		log.Fatal(err)
	}

	//	получаем значение переменной окружения: порт сервера
	port := os.Getenv("TODO_PORT")

	//	определяем функции для работы с URL
	api.Init() // Функция-обработчик для "/api/"

	//	обработка запросов на возврат файлов
	http.HandleFunc("/js/", handleJS)
	http.HandleFunc("/css/", handleCSS)
	http.HandleFunc("/favicon.ico", handleIco)

	//	обработка запроса на возврат формы пользователю
	http.HandleFunc("/", handleHTML)

	//	для удобства выводим в консоль сообщения о нашем сервере
	fmt.Println("Сервер запущен на http://localhost:7540")
	fmt.Println("Попробуйте пройти по ссылкам:\nhttp://localhost:7540/js/scripts.min.js\nhttp://localhost:7540/css/style.css\nhttp://localhost:7540/favicon.ico\nhttp://localhost:7540/api/nextdate?now=20240126&date=20240126&repeat=y\n")

	//	запускаем сервер
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
