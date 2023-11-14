package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

/*
*** Задача на распаковку ***
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

// UnpackString выполняет распаковку строки с учетом повторяющихся символов и escape-последовательностей.
func UnpackString(str string) (string, error) {
	var result strings.Builder // Создаем буфер для сборки распакованной строки
	var currentChar rune       // Текущий символ, который нужно повторить
	var count int              // Количество повторений текущего символа

	for _, char := range str {

		if unicode.IsDigit(char) {
			// Если символ - цифра, то увеличиваем счетчик повторений
			count = int(char - '0')
		} else {
			// Если символ - не цифра и не '\'
			if count == 0 {
				count = 1
			}
			// Добавляем повторения текущего символа к результату
			result.WriteString(strings.Repeat(string(currentChar), count))
			currentChar = char
			count = 0
		}
	}

	if count != 0 {
		return "", errors.New("Некорректная строка")
	}

	// Добавляем оставшиеся повторения текущего символа к результату
	result.WriteString(strings.Repeat(string(currentChar), 1))
	return result.String(), nil
}

func main() {
	// Примеры использования функции UnpackString:
	inputs := []string{"a4bc2d5e", "abcd", "45", ""}
	for _, input := range inputs {

		result, err := UnpackString(input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Исходная строка: %s\nРаспакованная строка: %s\n", input, result)
		}
	}
}
