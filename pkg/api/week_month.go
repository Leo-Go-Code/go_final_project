package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func week(now, date, repeat string) (string, error) {
	//	1.	Преобразуем now в формат time.Time
	nowTime, err := time.Parse(FORMAT_DATE, now)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга исходной даты now")
	}

	//	2.	Преобразуем date в формат time.Time
	dateTime, err := time.Parse(FORMAT_DATE, date)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга исходной даты now")
	}

	//	3.	Определяем дату для отсчёта
	workDate := nowTime.AddDate(0, 0, 1)
	dateTime = dateTime.AddDate(0, 0, 1)
	if dateTime.After(nowTime) {
		workDate = dateTime
	}

	//	4.	Переменные для определения дней недели в правиле повторения
	weekDaysSlice := []int{}
	weekDaysMap := map[string]int{
		//	при 1-7 мой код работал даже с НЕкорректными repeat, но из-за теста добавил цифры 0, 8, 9 и проверку ниже
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	//	5.	Вырезаем подстроку %2C из строки repeat
	repeat = strings.ReplaceAll(repeat, "%2C", "")

	//	6.	Определение дней недели в правиле повторения
	for _, symbol := range repeat {
		symbolStr := string(symbol)
		if num, ok := weekDaysMap[symbolStr]; ok {
			weekDaysSlice = append(weekDaysSlice, num)
		}
	}
	if len(weekDaysSlice) == 0 {
		return "", fmt.Errorf("некорректное значение repeat: НЕ найдены числовые значения правила повторения для w")
	}
	for _, n := range weekDaysSlice {
		if n < 1 || n > 7 {
			return "", fmt.Errorf("некорректное значение repeat: найдены числовые значения меньше 1 или больше 7 для правила повторения w")
		}
	}

	//	7.	Упорядочить weekDaysSlice
	for j := 0; j < len(weekDaysSlice)-1; j++ {
		for i := 0; i < len(weekDaysSlice)-1; i++ {
			if weekDaysSlice[i] > weekDaysSlice[i+1] {
				weekDaysSlice[i], weekDaysSlice[i+1] = weekDaysSlice[i+1], weekDaysSlice[i]
			}
		}
	}

	//	8.	Дата для расчёта вывода: инициируем переменную, получаем значение из функции AddDate и проверяем его

	for {
		workDate = workDate.AddDate(0, 0, 1)

		weekDay := int(workDate.Weekday())
		if weekDay == 0 {
			weekDay = 7
		}

		for _, day := range weekDaysSlice {
			if weekDay == day {
				return workDate.Format(FORMAT_DATE), err
			}
		}
	}
}

func month(now, date, repeat string) (string, error) {
	//	1.	Преобразуем now в формат time.Time
	nowTime, err := time.Parse(FORMAT_DATE, now)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга исходной даты now")
	}

	//	2.	Преобразуем date в формат time.Time
	dateTime, err := time.Parse(FORMAT_DATE, date)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга исходной даты now")
	}

	//	3.	Определяем дату для отсчёта
	workDate := nowTime.AddDate(0, 0, 1)
	dateTime = dateTime.AddDate(0, 0, 1)
	if dateTime.After(nowTime) {
		workDate = dateTime
	}

	//	4.	Определяем день и месяц workDate для отсчёта в формате int
	var dayInt, monthInt int
	dayInt, err = strconv.Atoi(workDate.Format(FORMAT_DATE)[6:8])
	if err != nil {
		return "", err
	}
	monthInt, err = strconv.Atoi(workDate.Format(FORMAT_DATE)[4:6])
	if err != nil {
		return "", err
	}

	//	5.	Получаем из repeat строки дней и месяцев
	monthStr := repeat[2:]
	monthNumbersStr := ""
	monthDatesStr := ""
	if strings.Contains(monthStr, " ") {
		index := strings.Index(monthStr, " ")
		monthNumbersStr = monthStr[index+1:]
		monthDatesStr = monthStr[:index]
	} else if strings.Contains(monthStr, "+") {
		index := strings.Index(monthStr, "+")
		monthNumbersStr = monthStr[index+1:]
		monthDatesStr = monthStr[:index]
	} else {
		monthDatesStr = monthStr
	}

	//	6.	Определение месяцев в правиле повторения repeat
	monthNumbersInt := []int{}
	if monthNumbersStr != "" {
		monthNumbersStrs := []string{}
		if strings.Contains(monthNumbersStr, "%2C") {
			monthNumbersStrs = strings.Split(monthNumbersStr, "%2C")
		} else {
			monthNumbersStrs = strings.Split(monthNumbersStr, ",")
		}
		// fmt.Println(monthNumbersStrs)
		for _, str := range monthNumbersStrs {
			numb, err := strconv.Atoi(str)
			if err != nil {
				return "", err
			}
			monthNumbersInt = append(monthNumbersInt, numb)
		}
	}

	//	7.	Определение дней в правиле повторения repeat
	monthDatesInt := []int{}
	monthDatesStrs := []string{}
	if strings.Contains(monthDatesStr, "%2C") {
		monthDatesStrs = strings.Split(monthDatesStr, "%2C")
	} else {
		monthDatesStrs = strings.Split(monthDatesStr, ",")
	}
	for _, str := range monthDatesStrs {
		numb, err := strconv.Atoi(str)
		if err != nil {
			return "", err
		}
		monthDatesInt = append(monthDatesInt, numb)
	}
	if len(monthDatesInt) == 0 {
		return "", fmt.Errorf("некорректное значение repeat: не найдены числовые значения правила повторения по датам")
	}

	//	8.	Упорядочить monthDatesInt
	if len(monthDatesInt) != 1 {
		for j := 1; j < len(monthDatesInt); j++ {
			for i := 1; i < len(monthDatesInt); i++ {
				if monthDatesInt[i-1] > monthDatesInt[i] {
					monthDatesInt[i-1], monthDatesInt[i] = monthDatesInt[i], monthDatesInt[i-1]
				}
			}
		}

		negInt := []int{}
		for i := 0; i < len(monthDatesInt); i++ {
			if monthDatesInt[i] < 0 {
				negInt = append(negInt, monthDatesInt[i])
			}
		}

		if len(negInt) > 0 {
			monthDatesInt = monthDatesInt[len(negInt):]

			for j := 1; j < len(negInt); j++ {
				for i := 1; i < len(negInt); i++ {
					if negInt[i-1] > negInt[i] {
						negInt[i-1], negInt[i] = negInt[i], negInt[i-1]
					}
				}
			}

			monthDatesInt = append(monthDatesInt, negInt...)
		}
	}

	//	9.	Проверка monthDatesInt
	for i := 0; i < len(monthDatesInt); i++ {
		if monthDatesInt[i] < -2 || monthDatesInt[i] > 31 {
			return "", fmt.Errorf("некорректное значение repeat: число месяца НЕ может быть задано меньше -2 или больше 31")
		}
	}

	//	10.	Упорядочить monthNumbersInt
	for j := 1; j < len(monthNumbersInt); j++ {
		for i := 1; i < len(monthNumbersInt); i++ {
			if monthNumbersInt[i-1] > monthNumbersInt[i] {
				monthNumbersInt[i-1], monthNumbersInt[i] = monthNumbersInt[i], monthNumbersInt[i-1]
			}
		}
	}

	//	11.	Проверка monthNumbersInt
	for i := 0; i < len(monthNumbersInt); i++ {
		if monthNumbersInt[i] < 1 || monthNumbersInt[i] > 12 {
			return "", fmt.Errorf("некорректное значение repeat: число месяцев НЕ может быть задано меньше 1 или больше 12")
		}
	}

	//	12.	Ищем нужную дату, прибавляя дни и месяцы
	d := 0
	for i := 0; i < 5; i++ {
		//	12.1	Добавляем дни и месяцы
		monthDatesIntCopy := monthDatesInt
		d, workDate, err = monthDayAdder(monthDatesIntCopy, monthNumbersInt, monthInt, dayInt, workDate)
		if err != nil {
			return workDate.Format(FORMAT_DATE), err
		}

		//	12.2	Выход из цикла, если получена подходящая дата
		if d != 0 {
			break
		}

		//	12.3	Обновляем значения dayInt, monthInt
		dayInt, err = strconv.Atoi(workDate.Format(FORMAT_DATE)[6:8])
		if err != nil {
			return "", err
		}
		monthInt, err = strconv.Atoi(workDate.Format(FORMAT_DATE)[4:6])
		if err != nil {
			return "", err
		}
	}

	//	13.	ВЫВОД
	return workDate.Format(FORMAT_DATE), err
}

func monthDayAdder(monthDatesInt, monthNumbersInt []int, monthInt, dayInt int, workDate time.Time) (int, time.Time, error) {
	//	0.	Вводим переменные для добавления к текущей дате
	months := 0
	d := 0

	//	1.	Определяем сколько месяцев нужно добавить и делаем это, обновляя workDate.Day() по необходимости
	if len(monthNumbersInt) != 0 {
		br := 0
		for i := 0; i < 12; i++ {
			for _, month := range monthNumbersInt {
				if monthInt == month {
					br++
					break
				}
			}

			if br > 0 {
				break
			}

			monthInt++
			if monthInt > 12 {
				monthInt = 1
			}
			months++

			if months == 1 {
				workDate = time.Date(workDate.Year(), workDate.Month(), 1, workDate.Hour(), workDate.Minute(), workDate.Second(), workDate.Nanosecond(), workDate.Location())
				dayInt = 1
			}
		}
		workDate = workDate.AddDate(0, months, 0)
	}

	//	2.	Определяем полученные месяц и год
	workDateStr := workDate.Format(FORMAT_DATE)
	resMonthStr := workDateStr[4:6]
	resMonthInt, err := strconv.Atoi(resMonthStr)
	if err != nil {
		return d, workDate, err
	}
	resYearStr := workDateStr[:4]
	resYearInt, err := strconv.Atoi(resYearStr)
	if err != nil {
		return d, workDate, err
	}

	//	3.	Определяем кол-во дней в полученном месяце полученного года
	maxDay := 0
	if resMonthInt == 4 || resMonthInt == 6 || resMonthInt == 9 || resMonthInt == 11 {
		maxDay = 30
	} else if resMonthInt == 1 || resMonthInt == 3 || resMonthInt == 5 || resMonthInt == 7 || resMonthInt == 8 || resMonthInt == 10 || resMonthInt == 12 {
		maxDay = 31
	} else {
		if (resYearInt%4 == 0 && resYearInt%100 != 0) || resYearInt%4 == 400 {
			maxDay = 29
		} else {
			maxDay = 28
		}
	}

	//	4.	Конвертируем отрицательные даты месяца в положительные
	for i := 0; i < len(monthDatesInt); i++ {
		if monthDatesInt[i] < 0 {
			monthDatesInt[i] = maxDay + monthDatesInt[i] + 1
		}
	}

	//	5.	Определяем сколько дней нужно добавить и делаем это
	d, workDate, err = dayAdder(monthDatesInt, dayInt, maxDay, workDate)
	if err != nil {
		return d, workDate, err
	}

	return d, workDate, err
}

func dayAdder(monthDatesInt []int, dayInt, maxDay int, workDate time.Time) (int, time.Time, error) {
	//	0.	Вводим переменные для добавления к текущей дате
	d := 0
	days := 0

	//	1.	Проверка среза monthDatesInt на наличие значений/дат
	if len(monthDatesInt) == 0 {
		return 0, workDate, fmt.Errorf("monthDatesInt должен содержать даты подходящих дней для повторения задачи")
	}

	//	2.	Если рабочая дата совпадает со списочной, поиск завершён
	for _, day := range monthDatesInt {
		if dayInt == day {
			return 1, workDate, nil
		}
	}

	//	3.	Если рабочая дата меньше одной из списочных, и списочная дата НЕ больше максимального кол-ва дней в месяце, поиск завершён
	if d == 0 {
		for _, day := range monthDatesInt {
			if day > dayInt && day <= maxDay {
				days = day - dayInt
				workDate = workDate.AddDate(0, 0, days)
				return 1, workDate, nil
			}
		}
	}

	//	4.	Делаем переход в следующий месяц через добавление дней к workDate
	if d == 0 {
		days = maxDay - dayInt + 1
		workDate = workDate.AddDate(0, 0, days)
		return 0, workDate, nil
	}

	return d, workDate, err
}
