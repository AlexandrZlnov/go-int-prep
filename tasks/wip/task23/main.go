// Задача:
// Скомпилируется или нет?
// Если да, то что выведет?

// Ответ:
// Скомпилируется
// Вывод: 0 1 2 3 4
// break в данном случае завершит только switch, а общий цикл продолжиться.
// что бы выйти из цикла нужно использовать метку loop: for...  break loop

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break
		}
	}
}
