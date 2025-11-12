package api

import (
	"fmt"
	"time"
)

// Проверка корректности строк now и dstart
func checkDstartNow(str string) error {
	//	1. Проверка количества символов
	if len(str) != 8 {
		return fmt.Errorf("строка str задана некорректно: количество символов НЕ равно 8")
	}

	//	2. Проверка самих символов строки (должны быть только цифры)
	for _, letter := range str {
		if letter < '0' || letter > '9' {
			return fmt.Errorf("строка str задана некорректно: строка содержит не только цифры, что недопустимо")
		}
	}

	//	3. Проверка полученной даты через time.Parse
	_, err := time.Parse(FormateDate, str)
	if err != nil {
		return fmt.Errorf("строка str задана некорректно: %w", err)
	}

	return err
}

// проверка корректности строки repeat
func checkRepeat(repeat string) error {
	//	0.	Инициируем переменные для проверки
	lettersBytes := []byte("ydwm")
	numbersBytes := []byte("0123456789")

	//	1. Проверка наличия значения
	if len(repeat) == 0 {
		return fmt.Errorf("строка repeat задана некорректно: длина строки НЕ может быть равна 0")
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
	if len(repeat) > 100 {
		return fmt.Errorf("строка repeat задана некорректно: строка НЕ может иметь длину больше 100 символов")
	}

	//	6. Проверка наличия пробела между буквами и цифрами для "d", "w", "m"
	if len(repeat) > 1 && (repeat[1] != '+' && repeat[1] != ' ') {
		return fmt.Errorf(`строка repeat задана некорректно: после символов "d", "w", "m" должен ставиться пробел`)
	}
	return err
}
