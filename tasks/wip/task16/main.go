// Что выведет программа?

package main

import "fmt"

func add(s []int) {
	s = append(s, 99)
}

func main() {
	a := make([]int, 2, 4)
	a[0], a[1] = 1, 2

	add(a)

	fmt.Println("a:", a)
	fmt.Println("a[:3]:", a[:3])
}

// Ответ:
// fmt.Println("a:", a)         // 1 2
// fmt.Println("a[:3]:", a[:3]) // 1 2 99
