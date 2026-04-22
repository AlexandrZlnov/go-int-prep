// Что выведет код?

package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	t := s[:2]
	u := s[2:]

	// Изменяем через t
	t = append(t, 9)

	// Изменяем через u
	u[0] = 99

	fmt.Println("s:", s)
	fmt.Println("t:", t)
	fmt.Println("u:", u)
}

// Вывод:
// s: [1 2 99 4]
// t: [1 2 99]
// u: [99 4]
