// Что выведет код.
// Как сделать так чтобы s2 s3 не влияли друг на друга

package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s[0], s[1], s[2] = 1, 2, 3

	s2 := append(s, 4)
	s3 := append(s, 5)

	fmt.Println(s2)
	fmt.Println(s3)
}

// Ответ:
// 1 2 3 5
// 1 2 3 5
// Вариант 1 через make и coy:
// s2 := make([]int, 3, 5)
// copy(s2, s)
// s2 = append(s2, 4)
// Вариант 2 через full slice expression s[low:high:max]:
// s2 := s[:len(s):len(s)]
// s2 = append(s2, 4)
