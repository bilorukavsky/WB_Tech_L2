package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import (
	"fmt"
)

// Receiver - получатель команды
type Receiver struct {
	Name string
}

func (r *Receiver) Action() {
	fmt.Printf("%s: выполнение действия\n", r.Name)
}

// Command - интерфейс команды
type Command interface {
	Execute()
}

// ConcreteCommand - конкретная реализация команды
type ConcreteCommand struct {
	receiver *Receiver
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Invoker - вызывающий (исполнитель) команд
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.command = cmd
}

func (i *Invoker) ExecuteCommand() {
	fmt.Println("Исполнитель: вызываю команду")
	i.command.Execute()
}

/*
func main() {
	// Создание объектов получателя и команды
	receiver := &Receiver{Name: "Получатель"}
	command := &ConcreteCommand{receiver: receiver}

	// Создание вызывающего объекта и назначение команды
	invoker := &Invoker{}
	invoker.SetCommand(command)

	// Вызов команды
	invoker.ExecuteCommand()
}
*/
