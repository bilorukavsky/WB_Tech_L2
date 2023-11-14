package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
*** Базовая задача ***
Создать программу, печатающую точное время с использованием NTP библиотеки.
Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена как go module.
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS.
*/

func GetTime(server string) (time.Time, error) {
	return ntp.Time(server)
}

func main() {
	ntpServer := "pool.ntp.org"
	ntpTime, err := GetTime(ntpServer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при запросе NTP-сервера: %s", err)
		os.Exit(1)
	}

	formattedTime := ntpTime.Format("2006-01-02 15:04:05")
	fmt.Printf("Текущее точное время: %v\n", formattedTime)

}
