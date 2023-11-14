package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
***Взаимодействие с ОС***
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func changeDirectory(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("cd: missing argument")
	}

	err := os.Chdir(args[1])
	return err
}

func printCurrentDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(dir)
}

func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Println("kill: missing argument")
		return
	}

	pid := args[1]
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("kill: invalid process ID")
		return
	}

	process, err := os.FindProcess(pidInt)
	if err != nil {
		fmt.Println("kill:", err)
		return
	}

	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Println("kill:", err)
		return
	}

	fmt.Println("Process killed:", pid)
}

func showProcesses() {
	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("ps:", err)
		return
	}

	fmt.Println(string(output))
}

func runCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	fmt.Println("Simple Shell. Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		args := strings.Fields(input)

		switch args[0] {
		case "cd":
			err := changeDirectory(args)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "pwd":
			printCurrentDirectory()
		case "echo":
			echo(args)
		case "kill":
			killProcess(args)
		case "ps":
			showProcesses()
		default:
			runCommand(args)
		}
	}
}
