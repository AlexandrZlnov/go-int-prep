package main

import (
	"fmt"
)

func main() {
	m := append(make([]int, 0, 10), 1, 2, 3, 4) // создат слайс l=0 c=10. Добавит числа.
	_ = append(m, 5)                            // добавит 5 в исходный массив но не сохранит число в слайсе. m останется l=4 c=10

	fmt.Println(m)     // [1 2 3 4]
	fmt.Println(m[:6]) // [1 2 3 4 5 0]
}
