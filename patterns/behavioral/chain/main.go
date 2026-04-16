// Паттерн - Chain of Responsibility (Цепочка обязанностей)
// Как работает:
//   - запуск через go run .
//   - Можно запустить 2 варианта исполнения
//   - Для првоерки варианта с Middleware в терминале можно ввести:
//   - - curl localhost:8080/hello
//   - - curl -H "Authorization: token" localhost:8080/hello
//   - - curl -H "Authorization: token" "localhost:8080/hello?limit=true"

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Код позволяет запустить 2 варианта реализации паттерна Chain of Responsibility (Цепочка обязанностей).")
	fmt.Println(`Выбери нужный вариант:
1 - Классическая реализация
2 - Production middleware реализация
Введите номер:`)

	scanner.Scan()
	input := scanner.Text()

	switch input {
	case "1":
		RunClassicExample()
	case "2":
		RunMiddlewareExample()
	default:
		fmt.Println("Неизвестный варинт.")
		return
	}
}
