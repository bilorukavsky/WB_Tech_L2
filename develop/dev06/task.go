package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
***Утилита cut***

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

func parseFieldsToSelect(fieldsFlag string) map[int]bool {
	fields := make(map[int]bool)
	if fieldsFlag != "" {
		fieldsToSelect := strings.Split(fieldsFlag, ",")
		for _, field := range fieldsToSelect {
			fieldIndex := parseFieldIndex(field)
			fields[fieldIndex] = true
		}
	}
	return fields
}

func parseFieldIndex(field string) int {
	index, err := strconv.Atoi(field)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing field:", err)
		os.Exit(1)
	}
	return index
}

func processInput(input io.Reader, delimiter string, fieldsToSelect map[int]bool, printSeparated bool) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, delimiter) || !printSeparated {
			selectedColumns := selectFields(line, delimiter, fieldsToSelect)
			fmt.Println(strings.Join(selectedColumns, delimiter))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}

func selectFields(line, delimiter string, fieldsToSelect map[int]bool) []string {
	columns := strings.Split(line, delimiter)
	var selectedColumns []string
	for i, col := range columns {
		if len(fieldsToSelect) == 0 || fieldsToSelect[i+1] {
			selectedColumns = append(selectedColumns, col)
		}
	}
	return selectedColumns
}

func main() {
	delimiterFlag := flag.String("d", "\t", "Delimiter to split columns")
	fieldsFlag := flag.String("f", "", "Fields (columns) to select")
	separatedFlag := flag.Bool("s", false, "Print only lines with delimiter")

	flag.Parse()

	fieldsToSelect := parseFieldsToSelect(*fieldsFlag)

	processInput(os.Stdin, *delimiterFlag, fieldsToSelect, *separatedFlag)
}
