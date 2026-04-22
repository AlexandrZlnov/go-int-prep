// Что выведет программ?

package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // ляжет в стэк defer: 2, 1, 0
	}

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // с учетом версии go выше 1.22 будет замыкание ОТДЕЛЬНЫХ переменных и стек defer дополнится функциями с результатом: 2, 1, 0
		}()
	}

	for i := 0; i < 3; i++ {
		i := i
		defer func() {
			fmt.Println(i) // дополнит стэк defer: 2, 1, 0
		}()
	}
}

// вывод будет в столбец: 2 1 0 2 1 0 2 1 0
