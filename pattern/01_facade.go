package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import "fmt"

// Подсистема #1
type CPU struct{}

func (c *CPU) start() {
	fmt.Println("Запуск ЦПУ")
}

func (c *CPU) execute() {
	fmt.Println("Исполнение команд")
}

// Подсистема #2
type Memory struct{}

func (m *Memory) load() {
	fmt.Println("Загрузка данных в память")
}

// Подсистема #3
type HardDrive struct{}

func (hd *HardDrive) read() {
	fmt.Println("Чтение данных с жесткого диска")
}

// Фасад
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) StartComputer() {
	fmt.Println("Запуск компьютера...")
	cf.cpu.start()
	cf.memory.load()
	cf.hardDrive.read()
	cf.cpu.execute()
	fmt.Println("Компьютер запущен")
}

/*
func main() {
	// Использование Фасада
	computer := NewComputerFacade()
	computer.StartComputer()
}
*/
