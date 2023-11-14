package main

import (
	"sort"
	"strings"
)

/*
***Поиск анаграмм по словарю***

Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/

func findAnagrams(words []string) map[string][]string {
	// Создаем мапу для хранения множеств анаграмм
	anagramGroups := make(map[string][]string)

	// Пройдемся по всем словам в словаре
	for _, word := range words {
		// Приводим слово к нижнему регистру
		word = strings.ToLower(word)

		// Преобразуем слово в сортированный слайс байт
		wordBytes := []byte(word)
		sort.Slice(wordBytes, func(i, j int) bool { return wordBytes[i] < wordBytes[j] })

		// Преобразуем слайс байт обратно в строку
		sortedWord := string(wordBytes)

		// Добавляем слово в соответствующее множество анаграмм
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	// Удаляем множества из одного элемента
	for key, value := range anagramGroups {
		if len(value) == 1 {
			delete(anagramGroups, key)
		}
	}

	return anagramGroups
}

func main() {
	// Пример использования функции
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramGroups := findAnagrams(words)

	// Выводим результат
	for key, value := range anagramGroups {
		println(key, ":", strings.Join(value, ", "))
	}
}
