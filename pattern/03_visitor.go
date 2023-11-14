package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import "fmt"

// Интерфейс посетителя
type Visitor interface {
	VisitElementA(element *ElementA)
	VisitElementB(element *ElementB)
}

// Интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ElementA struct {
	Name string
}

func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

// Конкретный элемент B
type ElementB struct {
	Number int
}

func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitElementA(element *ElementA) {
	fmt.Printf("Visitor is processing ElementA with name: %s\n", element.Name)
}

func (cv *ConcreteVisitor) VisitElementB(element *ElementB) {
	fmt.Printf("Visitor is processing ElementB with number: %d\n", element.Number)
}
/*
func main() {
	// Создаем элементы
	elementA := &ElementA{Name: "ElementA"}
	elementB := &ElementB{Number: 54}

	// Создаем посетителя
	visitor := &ConcreteVisitor{}

	// Применяем посетителя к элементам
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
*/
