// Что выведет программ?

package main

import (
	"fmt"
)

type S struct {
	v int
}

func (s S) Value() int {
	return s.v
}

type I interface {
	Value() int
}

func mutate(i I) {
	// Мы думаем, что меняем реальный объект...
	if s, ok := i.(S); ok {
		s.v = 100
	}
}

func main() {
	s := S{v: 10}
	var i I = s

	mutate(i)

	fmt.Println(s.v)       // ??
	fmt.Println(i.Value()) // ??
}

// Вывод:
// 10
// 10

// в интерфейсе i хранится копия структуры S так как она не передается через указатель
// поэтому первый выдод = 10
// в функции mutate создается еще одна локальная копия if s, ok := i.(S) структны
// поэтому значение в структуре интерфейса не меняется
// что бы значения изменились:
//		if s, ok := i.(*S); ok {
//		s := &S{v: 10}
