// Что выведет программа?

package main

import "fmt"

func add(s []int) {
	s = append(s, 99)
	s[0] = 42
}

func main() {
	a := []int{1, 2}
	add(a)

	fmt.Println("a:", a)
}

// Ответ:
// fmt.Println("a:", a) // 1 2
