// Задача
// Что выведет код.

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3} // len=4 cap=4
	b := a                 // [0 1 2 3] len=4 cap=4

	b[1] = b[len(b)-1] // [0 3 2 3] len=4 cap=4
	b[len(b)-1] = 0    // [0 3 2 0] len=4 cap=4
	b = b[:len(b)-1]   //  [0 3 2] len=3 cap=4
	b[0] = 4           // [4 3 2] len=3 cap=4

	fmt.Println(a, b)                           // [4 3 2 0] [4 3 2]
	fmt.Println(len(a), len(b), cap(a), cap(b)) // 4 3 4 4

}
