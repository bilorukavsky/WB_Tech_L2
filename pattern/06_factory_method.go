package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

// ProductId - перечисление идентификаторов продуктов
type ProductId int

const (
	MINE ProductId = iota
	YOURS
	// можно добавить другие идентификаторы продуктов по мере необходимости
)

// Product - интерфейс продукта
type Product1 interface {
	Print()
}

// Creator - интерфейс фабрики
type Creator interface {
	Create(id ProductId) Product1
}

// ConcreteProductMINE - конкретная реализация продукта MINE
type ConcreteProductMINE struct{}

func (p *ConcreteProductMINE) Print() {
	fmt.Printf("Print MINE\n")
}

// ConcreteProductYOURS - конкретная реализация продукта YOURS
type ConcreteProductYOURS struct{}

func (p *ConcreteProductYOURS) Print() {
	fmt.Printf("Print YOURS\n")
}

// ConcreteCreator - конкретная реализация фабрики
type ConcreteCreator struct{}

func (c *ConcreteCreator) Create(id ProductId) Product1 {
	switch id {
	case MINE:
		return &ConcreteProductMINE{}
	case YOURS:
		return &ConcreteProductYOURS{}
	default:
		return nil
	}
}

/*
func main() {
	// Создаем экземпляр конкретной фабрики
	var creator Creator = &ConcreteCreator{}

	// Создаем продукт MINE и вызываем его метод
	product := creator.Create(MINE)
	product.Print()

	// Создаем продукт YOURS и вызываем его метод
	product = creator.Create(YOURS)
	product.Print()
}
*/
