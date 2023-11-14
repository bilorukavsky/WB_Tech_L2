package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
***Утилита telnet***
Реализовать простейший telnet-клиент.
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

func connectToServer(host, port string, timeout time.Duration) (net.Conn, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func readFromSocket(conn net.Conn) {
	// чтения данных из сокета и вывода их в STDOUT
	io.Copy(os.Stdout, conn)
	fmt.Println("Connection closed by server.")
	os.Exit(0)
}

func writeToSocket(conn net.Conn) {
	// Копирование данных из STDIN в сокет
	io.Copy(conn, os.Stdin)
	fmt.Println("Connection closed by user.")
}

func main() {
	// Обработка флагов командной строки
	host := flag.String("host", "", "Host to connect to")
	port := flag.String("port", "23", "Port to connect to")
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	// Проверка наличия обязательного параметра - хоста
	if *host == "" {
		fmt.Println("Please provide a valid host.")
		return
	}

	// Подключение к серверу
	conn, err := connectToServer(*host, *port, *timeout)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", fmt.Sprintf("%s:%s", *host, *port))

	// Запуск горутины для чтения данных из сокета и вывода их в STDOUT
	go readFromSocket(conn)

	// Копирование данных из STDIN в сокет
	writeToSocket(conn)
}
