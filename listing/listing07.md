Что выведет программа? Объяснить вывод программы.

```go
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет сначала значения из каналов а и b, а затем будет выводить нули в бесконечном цикле.

Функция asChan() создает новый канал, закрывает его и возвращает из функции.
Функция merge() берет значения из закрытых каналов и записывает их в новый канал с
в канал с будут успешно записаны все значения из каналов а и b, но затем из каналов будут
браться значения по умаолчанию (ноль) и тоже записываться в канал c.
И, так как канал с не был закрыт, при извлечении из него возникнет бесконечный цикл.
```