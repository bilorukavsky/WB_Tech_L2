package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

// Context - контекст, использующий различные состояния
type Context1 struct {
	state State
}

// State - интерфейс состояния
type State interface {
	Handle(c *Context1)
}

func NewContext() *Context1 {
	return &Context1{state: &ConcreteStateA{}}
}

func (c *Context1) setState(state State) {
	c.state = state
}

func (c *Context1) Request() {
	c.state.Handle(c)
}

// ConcreteStateA - конкретное состояние A
type ConcreteStateA struct{}

func (csa *ConcreteStateA) Handle(c *Context1) {
	fmt.Println("Обработка запроса в состоянии A")
	c.setState(&ConcreteStateB{})
}

// ConcreteStateB - конкретное состояние B
type ConcreteStateB struct{}

func (csb *ConcreteStateB) Handle(c *Context1) {
	fmt.Println("Обработка запроса в состоянии B")
	c.setState(&ConcreteStateA{})
}

/*
func main() {
	context := NewContext()

	// Последовательное выполнение запросов
	context.Request() // Ожидаемый вывод: Обработка запроса в состоянии A
	context.Request() // Ожидаемый вывод: Обработка запроса в состоянии B
	context.Request() // Ожидаемый вывод: Обработка запроса в состоянии A
}
*/
