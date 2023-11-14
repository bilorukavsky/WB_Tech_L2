package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*
***Утилита wget***
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

func main() {
	// Проверяем, передан ли URL в качестве аргумента командной строки
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run task.go <url>")
		return
	}

	// Получаем URL из аргумента командной строки
	url := os.Args[1]
	err := downloadSite(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// downloadSite загружает сайт, начиная с указанного URL
func downloadSite(url string) error {
	// Получаем HTTP-ответ от сервера
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем, что ответ успешен (статус 200 OK)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: %s", resp.Status)
	}

	// Анализируем HTML для поиска ссылок
	links, err := parseLinks(resp.Body)
	if err != nil {
		return err
	}

	// Скачиваем каждый файл по найденным ссылкам
	for _, link := range links {
		// Если ссылка начинается с "/", добавляем к ней базовый URL
		if strings.HasPrefix(link, "/") {
			link = url + link
		}

		// Скачиваем файл
		err := downloadFile(link)
		if err != nil {
			fmt.Println("Error downloading", link, ":", err)
		}
	}

	return nil
}

// parseLinks анализирует HTML и возвращает список найденных ссылок
func parseLinks(r io.Reader) ([]string, error) {
	tokenizer := html.NewTokenizer(r)
	links := make([]string, 0)

	// Парсим HTML
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		// Если достигнут конец HTML
		case html.ErrorToken:
			return links, nil
		// Если встречен тег <a> (ссылка)
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				// Ищем атрибут href и добавляем его значение в список
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

// downloadFile скачивает файл по указанной ссылке
func downloadFile(url string) error {
	// Получаем HTTP-ответ от сервера
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем, что ответ успешен (статус 200 OK)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: %s", resp.Status)
	}

	// Извлекаем имя файла из URL
	fileName := getFileName(url)

	// Создаем файл на диске
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем содержимое файла из HTTP-ответа в созданный файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	// Выводим сообщение об успешном скачивании
	fmt.Println("Downloaded:", fileName)
	return nil
}

// getFileName извлекает имя файла из URL
func getFileName(url string) string {
	// Извлекаем последнюю часть URL в качестве имени файла
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
