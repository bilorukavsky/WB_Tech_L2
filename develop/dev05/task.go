package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
***Утилита grep***
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

func parseFlags() (string, string, int, int, bool, bool, bool, bool, bool) {
	inputFile := flag.String("input", "", "Исходный файл для поиска")
	pattern := flag.String("pattern", "", "Паттерн для поиска")
	afterLines := flag.Int("A", 0, "Количество строк после совпадения")
	beforeLines := flag.Int("B", 0, "Количество строк до совпадения")
	countOnly := flag.Bool("c", false, "Выводить только количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invertMatch := flag.Bool("v", false, "Исключать совпадения")
	fixedString := flag.Bool("F", false, "Точное совпадение со строкой, не паттерн")
	showLineNumbers := flag.Bool("n", false, "Выводить номера строк")

	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Необходимо указать исходный файл для поиска с помощью -input")
		os.Exit(1)
	}

	return *inputFile, *pattern, *afterLines, *beforeLines, *countOnly, *ignoreCase, *invertMatch, *fixedString, *showLineNumbers
}

func prepareRegex(pattern string, ignoreCase, fixedString bool) *regexp.Regexp {
	if fixedString {
		pattern = regexp.QuoteMeta(pattern)
	}
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	return regexp.MustCompile(pattern)
}

func processFile(file *os.File, re *regexp.Regexp, beforeLines, afterLines int, invertMatch, showLineNumbers bool) []string {
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	foundMatch := false
	matching := false
	matchedLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			if invertMatch {
				matching = false
			} else {
				matching = true
				foundMatch = true
			}
		} else {
			if invertMatch {
				matching = true
			} else {
				matching = false
			}
		}

		if matching {
			matchedLines = append(matchedLines, line)
		} else if foundMatch {
			beforeStart := len(matchedLines) - beforeLines
			if beforeStart < 0 {
				beforeStart = 0
			}
			printMatchingLines(matchedLines[beforeStart:], showLineNumbers)
			matchedLines = []string{}
			for i := 1; i <= afterLines; i++ {
				if scanner.Scan() {
					line = scanner.Text()
					if showLineNumbers {
						fmt.Printf("%d:", lineNumber+i)
					}
					fmt.Println(line)
				} else {
					break
				}
			}
			matchedLines = []string{}
			foundMatch = false
		}

		lineNumber++
	}

	if foundMatch && !invertMatch {
		printMatchingLines(matchedLines, showLineNumbers)
	}

	return matchedLines
}

func printMatchingLines(lines []string, showLineNumbers bool) {
	for i, line := range lines {
		if showLineNumbers {
			fmt.Printf("%d:", i+1)
		}
		fmt.Println(line)
	}
}

func printMatchCount(count int) {
	fmt.Printf("Total lines with match: %d\n", count)
}

func main() {
	inputFile, pattern, afterLines, beforeLines, countOnly, ignoreCase, invertMatch, fixedString, showLineNumbers := parseFlags()

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка открытия исходного файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	re := prepareRegex(pattern, ignoreCase, fixedString)
	matchedLines := processFile(file, re, beforeLines, afterLines, invertMatch, showLineNumbers)

	if countOnly {
		printMatchCount(len(matchedLines))
	} else {
		printMatchingLines(matchedLines, showLineNumbers)
	}
}
