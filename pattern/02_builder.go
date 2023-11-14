package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Пример структуры для объекта, который будет строиться
type Product struct {
	PartA string
	PartB string
	PartC string
}

// Интерфейс Строителя
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// Конкретный Строитель #1
type ConcreteBuilder1 struct {
	product *Product
}

func NewConcreteBuilder1() *ConcreteBuilder1 {
	return &ConcreteBuilder1{product: &Product{}}
}

func (b *ConcreteBuilder1) BuildPartA() {
	b.product.PartA = "PartA built by ConcreteBuilder1"
}

func (b *ConcreteBuilder1) BuildPartB() {
	b.product.PartB = "PartB built by ConcreteBuilder1"
}

func (b *ConcreteBuilder1) BuildPartC() {
	b.product.PartC = "PartC built by ConcreteBuilder1"
}

func (b *ConcreteBuilder1) GetProduct() *Product {
	return b.product
}

// Конкретный Строитель #2
type ConcreteBuilder2 struct {
	product *Product
}

func NewConcreteBuilder2() *ConcreteBuilder2 {
	return &ConcreteBuilder2{product: &Product{}}
}

func (b *ConcreteBuilder2) BuildPartA() {
	b.product.PartA = "PartA built by ConcreteBuilder2"
}

func (b *ConcreteBuilder2) BuildPartB() {
	b.product.PartB = "PartB built by ConcreteBuilder2"
}

func (b *ConcreteBuilder2) BuildPartC() {
	b.product.PartC = "PartC built by ConcreteBuilder2"
}

func (b *ConcreteBuilder2) GetProduct() *Product {
	return b.product
}

// Директор, который управляет процессом строительства
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *Product {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
	return d.builder.GetProduct()
}

/*
func main() {
	// Пример использования паттерна Строитель
	builder1 := NewConcreteBuilder1()
	director := NewDirector(builder1)
	product1 := director.Construct()
	fmt.Println("Product built by ConcreteBuilder1:")
	fmt.Printf("PartA: %s\nPartB: %s\nPartC: %s\n", product1.PartA, product1.PartB, product1.PartC)

	fmt.Println()

	builder2 := NewConcreteBuilder2()
	director = NewDirector(builder2)
	product2 := director.Construct()
	fmt.Println("Product built by ConcreteBuilder2:")
	fmt.Printf("PartA: %s\nPartB: %s\nPartC: %s\n", product2.PartA, product2.PartB, product2.PartC)
}
*/
