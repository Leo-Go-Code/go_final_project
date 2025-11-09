package api

import (
	"fmt"
	"strconv"
	"time"
)

// Проверка корректности строк now и dstart
func CheckDstartNow(str string) error {
	//	0.	Инициируем переменную для проверки
	numbersBytes := []byte("0123456789")

	//	1. Проверка количества символов
	if len(str) != 8 {
		return fmt.Errorf("строка str задана некорректно: количество символов НЕ равно 8")
	}

	//	2. Проверка самих символов строки (должны быть только цифры)
	for i := 0; i < len(str); i++ {
		symbol := false
		for j := 0; j < len(numbersBytes); j++ {
			if str[i] == numbersBytes[j] {
				symbol = true
				break
			}
		}
		if !symbol {
			return fmt.Errorf("строка str задана некорректно: строка содержит НЕ только цифры")
		}
	}

	//	3. Проверка корректности значения дня
	dayStr := str[6:8]
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil || dayInt < 1 || dayInt > 31 {
		return fmt.Errorf("строка str задана некорректно: недопустимое кол-во дней в месяце")
	}

	//	4. Проверка корректности значения месяца
	monthStr := str[4:6]
	monthInt, err := strconv.Atoi(monthStr)
	if err != nil || monthInt < 1 || monthInt > 12 {
		return fmt.Errorf("строка str задана некорректно: недопустимое кол-во месяцев в году")
	}

	//	5. Проверка месяцев по 30 дней
	if (monthInt == 4 || monthInt == 6 || monthInt == 9 || monthInt == 11) && dayInt > 30 {
		return fmt.Errorf("строка str задана некорректно: недопустимое кол-во дней в 30ти-дневных месяцах")

	}

	// 6. Проверка февраля (на весокосный год тоже проверь)
	if monthInt == 2 && dayInt > 29 {
		return fmt.Errorf("строка str задана некорректно: недопустимое кол-во дней в феврале")
	}

	//	7. Дополнительная проверка через Parse
	_, err = time.Parse(formatDate, str)
	if err != nil {
		return fmt.Errorf("строка str задана некорректно: она не парсится под формат времени 20060102")
	}

	return err
}

// проверка корректности строки repeat
func CheckRepeat(repeat string) error {
	//	0.	Инициируем переменные для проверки
	lettersBytes := []byte("ydwm")
	numbersBytes := []byte("0123456789")

	//	1. Проверка наличия значения
	if len(repeat) == 0 {
		return fmt.Errorf(`строка repeat задана некорректно: длина строки НЕ может быть равна 0`)
	}

	//	2. Проверка значения y
	if len(repeat) == 1 {
		if repeat != "y" {
			return fmt.Errorf(`строка repeat задана некорректно: строка имеет только 1 символ и это НЕ "y"`)
		}
	}

	//	3. Проверка корректности буквенного обозначения
	q := 0
	for _, lettersByte := range lettersBytes {
		if repeat[0] == lettersByte {
			q++
		}
	}
	if q == 0 {
		return fmt.Errorf(`строка repeat задана некорректно: 1ый символ строки должен быть один из 4ёх вариантов: "y", "d", "w", "m"`)
	}

	//	4. Проверка наличия цифр
	if len(repeat) > 1 {
		isNumber := 0
		for i := 0; i < len(repeat); i++ {
			for j := 0; j < len(numbersBytes); j++ {
				if repeat[i] == numbersBytes[j] {
					isNumber++
				}
			}
		}
		if isNumber < 1 {
			return fmt.Errorf(`строка repeat задана некорректно: для вариантов "d", "w", "m" должно быть задано хотя бы одно число`)
		}
	}

	//	5. Проверка количества символов строки
	if len(repeat) > 20 {
		return fmt.Errorf("строка repeat задана некорректно: строка НЕ может иметь длину больше 20 символов")
	}

	//	6. Проверка наличия пробела между буквами и цифрами для "d", "w", "m"
	if len(repeat) > 1 && (repeat[1] != '+' && repeat[1] != ' ') {
		return fmt.Errorf(`строка repeat задана некорректно: после символов "d", "w", "m" должен ставиться пробел`)
	}
	return err
}
