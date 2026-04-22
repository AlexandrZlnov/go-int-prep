// Функция должна заполнить срез и мапу,
// но сейчас она падает с panic.
// В чём проблема?

package main

import (
	"fmt"
)

func fillData(s []int, m map[string]int) {
	// Добавляем в срез
	s = append(s, 10, 20, 30)
	m = make(map[string]int)

	// Добавляем в мапу
	m["count"] = len(s)
	m["total"] = 60

	fmt.Println("Срез:", s)
	fmt.Println("Мапа:", m)
}

func main() {
	fillData(nil, nil)
}

// Ответ:
// Мапа не инициализирована, некуда писать, нет указателя на hash таблицу.
