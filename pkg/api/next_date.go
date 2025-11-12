package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var result = ""
var err error

// задаём формат времени YYYYMMDD
const FormateDate = "20060102"

func nextDate(w http.ResponseWriter, r *http.Request) {
	//	0.	Проверка метода Get
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJSON(w, map[string]string{"error": "ошибка: не верно определён метод запроса для func nextDate()"})
		return
	}

	//	1.	Обрезаем URL
	infAPI := r.URL.RawQuery
	if infAPI == "" {
		http.Error(w, "URL Get-запрос не содержит информации для выполнения работы API", http.StatusBadRequest)
		return
	}

	//	2.	Получаем данные из URL
	//	можно было попробовать через FormValue()
	nowRightIndex := strings.Index(infAPI, "&date=")
	dateRightIndex := strings.Index(infAPI, "&repeat=")
	now := infAPI[len("now="):nowRightIndex]
	date := infAPI[nowRightIndex+len("&date=") : dateRightIndex]
	repeat := infAPI[dateRightIndex+len("&repeat="):]

	//	3.	Вызов функции taskDate для получения новой даты:
	//	проверка данных, определение добавляемого кол-ва дней, определение новой даты
	result, err := taskDate(now, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": fmt.Sprintf("pkg/api/nextdate.go вернул ошибку: %w", err)})
		return
	}

	//	4.	Возвращаем результат
	w.Header().Set("Content-Type", "text/text")
	w.Write([]byte(result))
}

func taskDate(now, date, repeat string) (string, error) {
	//	0.	Проверка корректности строк now, dstart, repeat
	err = checkDstartNow(now)
	if err != nil {
		return "", fmt.Errorf("ошибка при проверке корректности строки now: %w", err)
	}
	err = checkDstartNow(date)
	if err != nil {
		return "", fmt.Errorf("ошибка при проверке корректности строки date: %w", err)
	}
	err = checkRepeat(repeat)
	if err != nil {
		return "", fmt.Errorf("ошибка при проверке корректности строки repeat: %w", err)
	}

	//	1.	Определяем количество прибавляемых дней, месяцев и лет к исходной дате для получения новой даты
	//	1.0	Создаём переменные для добавления дней, месяцев, лет к текущей дате
	days := 0
	year := 0
	lettersBytes := []byte("ydwm")

	//	1.1	Считаем года - YEAR
	if repeat[0] == lettersBytes[0] {
		year++
	}

	//	1.2	Считаем дни - DAYS
	if repeat[0] == lettersBytes[1] {
		days, err = strconv.Atoi(repeat[len("d "):])
		if err != nil {
			return "", fmt.Errorf("ошибка преобразования части строки repeat в число")
		}
		if days > 400 {
			return "", fmt.Errorf("строка repeat задана некорректно: количество дней не должно превышать 400")
		}
	}

	//	1.3 Если дни и годы остаются пустыми, запускаем алгоритм вычисления даты для правил по неделям и месяцам
	if year == 0 && days == 0 {
		result := ""
		if repeat[0] == lettersBytes[2] { //	WEEK
			result, err = week(now, date, repeat)
			if err != nil {
				return "", err
			}
		}
		if repeat[0] == lettersBytes[3] { //	MONTH
			result, err = month(now, date, repeat)
			if err != nil {
				return "", err
			}
		}
		return result, err
	} else {

		//	1.4	Рассчитываем следующую дату для задачи, делаем проверку result > now и выводим результат для правил "d" и "y"
		result, err = resultDaysYear(now, date, days, year)
		if err != nil {
			return "", fmt.Errorf("ошибка работы функции resultDaysYear: %w", err)
		}
		return result, err
	}
}

func resultDaysYear(now, dstart string, days, year int) (string, error) {
	//	0.	Преобразуем now, dstart в формат time.Time
	nowTime, err := time.Parse(FormateDate, now)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга исходной даты now")
	}
	dateStart, err := time.Parse(FormateDate, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга исходной даты dstart")
	}

	//	1.	Дата для расчёта вывода: инициируем переменную, получаем значение из функции и проверяем его
	workDate := dateStart

	for {
		workDate = workDate.AddDate(year, 0, days)
		if workDate.After(nowTime) {
			break
		}
	}

	return workDate.Format(FormateDate), err
}
