package main

import (
	"fmt"
	"time"
)

/*
***Or channel ***
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь,
однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая,
приняв на вход один или более or-каналов, реализовывала бы весь функционал.
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orCh := make(chan interface{})
	for _, ch := range channels {
		go func(c <-chan interface{}) {
			for {
				select {
				case <-c:
					select {
					case orCh <- struct{}{}:
					case <-orCh:
					}
					return
				case <-orCh:
					return
				}
			}
		}(ch)
	}
	return orCh
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
