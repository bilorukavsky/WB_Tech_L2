Что выведет программа? Объяснить вывод программы.

```go
package main
type customError struct {
	msg string
}
func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет: error

Ситуация аналогична listing03: при вызове фунции test(), мы, правда, возвращаем значение nil, но
затем оно преобразуется в интерфейсный тип error, и при сравнении интерфейсного типа оба поля типа
интерфейс не равняются nil, что выполнит условие if и выведет сообщение error. 
```