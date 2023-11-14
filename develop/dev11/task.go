package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
***HTTP-сервер***

Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов

Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month


Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."}
в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

// Event представляет собой структуру для событий в календаре.
type Event struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
	Details string `json:"details"`
}

// Mutex для безопасного доступа к eventsStorage
var eventsMutex sync.Mutex

// eventsStorage для хранения событий
var eventsStorage = make(map[int]Event)

// SerializeEvent преобразует объект Event в формат JSON.
func SerializeEvent(event Event) ([]byte, error) {
	return json.Marshal(event)
}

// SerializeError преобразует сообщение об ошибке в формат JSON.
func SerializeError(errorMessage string) ([]byte, error) {
	errorResponse := map[string]string{"error": errorMessage}
	return json.Marshal(errorResponse)
}

// ParseEventParams извлекает параметры из запроса и возвращает объект Event.
// ParseEventParams извлекает параметры из запроса и возвращает объект Event.
func ParseEventParams(r *http.Request) (Event, error) {
	var event Event

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)
	if err != nil {
		return Event{}, err
	}

	if event.UserID == 0 {
		return Event{}, errors.New("missing user_id parameter")
	}

	return event, nil
}

// createEventHandler обрабатывает запрос на создание нового события.
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	event, err := ParseEventParams(r)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	// Генерация уникального ID для события
	event.ID = len(eventsStorage) + 1

	// Сохранение события в хранилище
	eventsStorage[event.ID] = event

	response, err := SerializeEvent(event)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// updateEventHandler обрабатывает запрос на обновление существующего события.
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	event, err := ParseEventParams(r)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	// Проверка существования события
	if _, exists := eventsStorage[event.ID]; !exists {
		handleError(w, errors.New("event not found"), http.StatusNotFound)
		return
	}

	// Обновление события в хранилище
	eventsStorage[event.ID] = event

	response, err := SerializeEvent(event)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handleError отправляет клиенту JSON с сообщением об ошибке и соответствующим статусом HTTP.
func handleError(w http.ResponseWriter, err error, status int) {
	errorResponse, _ := SerializeError(err.Error())
	http.Error(w, string(errorResponse), status)
}

// LoggingMiddleware логирует информацию о запросе.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)

		// Логирование информации о запросе
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, duration)
	}
}

// deleteEventHandler обрабатывает запрос на удаление события.
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	event, err := ParseEventParams(r)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	// Проверка существования события
	if _, exists := eventsStorage[event.ID]; !exists {
		handleError(w, errors.New("event not found"), http.StatusNotFound)
		return
	}

	// Удаление события из хранилища
	delete(eventsStorage, event.ID)

	w.WriteHeader(http.StatusOK)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()) // Начало текущего дня
	eventForDuration(w, startTime, 24*time.Hour)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	startTime := currentTime.AddDate(0, 0, -int(currentTime.Weekday())) // Начало текущей недели
	eventForDuration(w, startTime, 7*24*time.Hour)                      // Фильтруем события за последние 7 дней
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location()) // Начало текущего месяца
	nextMonth := startTime.AddDate(0, 1, 0)                                                                // Начало следующего месяца
	eventForDuration(w, startTime, nextMonth.Sub(startTime))                                               // Фильтруем события за текущий месяц
}

func eventForDuration(w http.ResponseWriter, startTime time.Time, duration time.Duration) {
	var eventArray []Event

	for _, v := range eventsStorage {
		parsedDate, err := time.Parse("2006-01-02", v.Date) // Формат "год-месяц-день"

		if err != nil {
			log.Printf("Error parsing date: %s", err)
			continue
		}

		// Оставляем только события, которые находятся в запрашиваемом интервале времени
		if parsedDate.After(startTime) && parsedDate.Before(startTime.Add(duration)) {
			eventArray = append(eventArray, v)
		}
	}

	unmarshalledEvent, err := json.Marshal(eventArray)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(unmarshalledEvent)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/create_event", LoggingMiddleware(createEventHandler))
	http.HandleFunc("/update_event", LoggingMiddleware(updateEventHandler))
	http.HandleFunc("/delete_event", LoggingMiddleware(deleteEventHandler))
	http.HandleFunc("/events_for_day", LoggingMiddleware(eventsForDayHandler))
	http.HandleFunc("/events_for_week", LoggingMiddleware(eventsForWeekHandler))
	http.HandleFunc("/events_for_month", LoggingMiddleware(eventsForMonthHandler))

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
