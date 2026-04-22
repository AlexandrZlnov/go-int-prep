// Что выведет код и почему?

package main

import "fmt"

func main() {
	s := make([]int, 2, 3)
	s[0], s[1] = 1, 2

	a := s[:2]
	b := append(s, 3)

	modify(a)
	modify(b)

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
}

func modify(x []int) {
	x = append(x, 100) // 2: 1 2 100 100
	x[0] = 999         // 2: 999 1 2 100 100
}

//Вывод:
// 999 2
// 999 2
// 999 2 100
