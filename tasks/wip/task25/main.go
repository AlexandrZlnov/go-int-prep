// Задача:
// Что выведет код?

// Ответ:
// Range запомнит длину слайса в самом начале поэтому будет ровно 3 цикла. Вывод: done: [0 1 2 10 10 10]

package main

import "fmt"

func main() {
	s := []int{0, 1, 2}

	for range s {
		s = append(s, 10)
	}

	fmt.Printf("done: %v\n", s)
}
