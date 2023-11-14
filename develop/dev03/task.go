package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
***Утилита sort***
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

// Проверка, отсортированы ли данные
func isSorted(lines []string, compareFunc func(i, j int) bool) bool {
	for i := 1; i < len(lines); i++ {
		if !compareFunc(i-1, i) {
			return false
		}
	}
	return true
}

// Удаление повторяющихся строк
func removeDuplicateLines(lines []string) []string {
	uniqueLines := make(map[string]bool)
	result := []string{}
	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			result = append(result, line)
		}
	}
	return result
}

// Сортировка по числовому значению
func numericCompare(line1, line2 string) bool {
	num1, err1 := strconv.Atoi(line1)
	num2, err2 := strconv.Atoi(line2)

	if err1 == nil && err2 == nil {
		return num1 < num2
	}

	return line1 < line2
}

// Сортировка по числовому значению с учетом суффиксов
func numericSuffixCompare(line1, line2 string) bool {
	num1, suffix1, err1 := parseNumericSuffix(line1)
	num2, suffix2, err2 := parseNumericSuffix(line2)

	if err1 == nil && err2 == nil {
		if num1 < num2 {
			return true
		} else if num1 > num2 {
			return false
		} else {
			return suffix1 < suffix2
		}
	}

	return line1 < line2
}

// Разбор числа с суффиксом (например, 42K)
func parseNumericSuffix(s string) (int, string, error) {
	parts := strings.Split(s, "")
	if len(parts) < 2 {
		return 0, s, fmt.Errorf("Invalid format")
	}
	numPart := parts[0 : len(parts)-1]
	suffix := parts[len(parts)-1]
	numStr := strings.Join(numPart, "")
	num, err := strconv.Atoi(numStr)
	return num, suffix, err
}

// Сортировка по названию месяца
func monthCompare(line1, line2 string) bool {
	months := []string{
		"январь", "февраль", "март", "апрель",
		"май", "июнь", "июль", "август",
		"сентябрь", "октябрь", "ноябрь", "декабрь",
	}
	index1 := -1
	index2 := -1
	for i, month := range months {
		if strings.Contains(line1, month) {
			index1 = i
		}
		if strings.Contains(line2, month) {
			index2 = i
		}
	}
	if index1 != -1 && index2 != -1 {
		return index1 < index2
	}
	return line1 < line2
}

func main() {
	// Определение флагов командной строки
	inputFile := flag.String("input", "", "Исходный файл для сортировки")
	outputFile := flag.String("output", "", "Файл для записи отсортированных данных")
	column := flag.Int("k", 0, "Номер колонки для сортировки (по умолчанию 0)")
	numericSort := flag.Bool("n", false, "Сортировка по числовому значению")
	reverseSort := flag.Bool("r", false, "Сортировка в обратном порядке")
	uniqueLines := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	monthSort := flag.Bool("M", false, "Сортировка по названию месяца")
	ignoreTrailingSpace := flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "Проверить, отсортированы ли данные")
	numericSuffix := flag.Bool("h", false, "Сортировка по числовому значению с учетом суффиксов")

	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Необходимо указать исходный файл для сортировки с помощью -input")
		os.Exit(1)
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Ошибка открытия исходного файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Определение функции сравнения для сортировки
	compareFunc := func(i, j int) bool {
		line1 := lines[i]
		line2 := lines[j]

		if *monthSort {
			// Сортировка по названию месяца
			return monthCompare(line1, line2)
		}

		if *numericSuffix {
			// Сортировка по числовому значению с учетом суффиксов
			return numericSuffixCompare(line1, line2)
		}

		if *numericSort {
			// Сортировка по числовому значению
			return numericCompare(line1, line2)
		}

		if *ignoreTrailingSpace {
			// Игнорирование хвостовых пробелов
			line1 = strings.TrimRight(line1, " ")
			line2 = strings.TrimRight(line2, " ")
		}

		// Сортировка по колонке
		columns1 := strings.Fields(line1)
		columns2 := strings.Fields(line2)

		if *column >= 1 && *column <= len(columns1) && *column <= len(columns2) {
			return columns1[*column-1] < columns2[*column-1]
		}

		return line1 < line2
	}

	if *uniqueLines {
		// Удаление повторяющихся строк
		lines = removeDuplicateLines(lines)
	}

	if *reverseSort {
		// Обратная сортировка
		sort.SliceStable(lines, func(i, j int) bool {
			return compareFunc(j, i)
		})
	} else {
		// Сортировка
		sort.SliceStable(lines, compareFunc)
	}

	if *outputFile != "" {
		// Запись отсортированных данных в файл
		output, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("Ошибка создания файла для записи:", err)
			os.Exit(1)
		}
		defer output.Close()
		for _, line := range lines {
			fmt.Fprintln(output, line)
		}
	} else {
		// Вывод отсортированных данных на стандартный вывод
		for _, line := range lines {
			fmt.Println(line)
		}
	}

	if *checkSorted {
		// Проверка, отсортированы ли данные
		if isSorted(lines, compareFunc) {
			fmt.Println("Данные отсортированы")
		} else {
			fmt.Println("Данные не отсортированы")
		}
	}
}
