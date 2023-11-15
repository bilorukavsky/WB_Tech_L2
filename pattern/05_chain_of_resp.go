package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import (
	"fmt"
)

// Handler - интерфейс обработчика запроса
type Handler interface {
	SetNext(handler Handler)
	Handle(request int)
}

// ConcreteHandler - конкретный обработчик запроса
type ConcreteHandler struct {
	nextHandler Handler
}

func (ch *ConcreteHandler) SetNext(handler Handler) {
	ch.nextHandler = handler
}

func (ch *ConcreteHandler) Handle(request int) {
	// В данном примере обработчик проверяет, может ли он обработать запрос
	// Если нет, то передает запрос следующему обработчику
	if request < 10 {
		fmt.Println("ConcreteHandler обрабатывает запрос")
	} else if ch.nextHandler != nil {
		fmt.Println("ConcreteHandler не может обработать запрос, передача следующему обработчику")
		ch.nextHandler.Handle(request)
	} else {
		fmt.Println("Нет обработчика для запроса")
	}
}

/*
func main() {
	// Создание цепочки обработчиков
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	// Установка следующих обработчиков в цепочке
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Запуск обработки запросов
	handler1.Handle(5)  // Обработчик 1 обрабатывает запрос
	handler1.Handle(12) // Обработчик 1 не может обработать, передается следующему
	handler1.Handle(20) // Обработчик 1 и 2 не могут обработать, передается последнему
}
*/
