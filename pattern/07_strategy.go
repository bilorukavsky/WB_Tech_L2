package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Интерфейс стратегии
type Strategy interface {
	DoOperation(int, int) int
}

// Конкретная реализация стратегии сложения
type OperationAdd struct{}

func (oa *OperationAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

// Конкретная реализация стратегии умножения
type OperationMultiply struct{}

func (om *OperationMultiply) DoOperation(num1, num2 int) int {
	return num1 * num2
}

// Контекст, который использует стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

/*
func main() {
	context := &Context{}

	// Использование стратегии сложения
	context.SetStrategy(&OperationAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("Результат сложения:", result) // Ожидаемый вывод: 15

	// Использование стратегии умножения
	context.SetStrategy(&OperationMultiply{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("Результат умножения:", result) // Ожидаемый вывод: 50
}
*/
