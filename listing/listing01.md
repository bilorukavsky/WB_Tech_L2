Что выведет программа? Объяснить вывод программы.

```go
package main
import (
    "fmt"
)
func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Индексация элеметов массива "a" имеют следующий вид:
```

| Индекс   | 0   | 1   | 2   | 3   | 4   |
|----------|-----|-----|-----|-----|-----|
| Элемент  | 76  | 77  | 78  | 79  | 80  |

```
Когда берем слайс [A:B] - формируется слайс, нижняя граница которого - A (включительно)
а верхняя граница - индекс B (не включительно). 

В нашем случае будет формироваться слайс из элементов с индексами 1, 2, 3: [77 78 79]
```