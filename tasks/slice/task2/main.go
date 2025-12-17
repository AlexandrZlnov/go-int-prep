// Слайы и мапы
package main

import "fmt"

func main() {
	a := make([]int, 0, 3)
	var a []int
	fmt.Println(a == nil)
	fmt.Printf("Array: %v, Len: %v, Cap: %v, Pointer: %p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	//fmt.Println(a, len(a), cap(a))
	fmt.Printf("Array: %v, Len: %v, Cap: %v, Pointer: %p\n", a, len(a), cap(a), a)
	a = append(a, 2, 3, 4)
	fmt.Printf("Array: %v, Len: %v, Cap: %v, Pointer: %p\n", a, len(a), cap(a), a)

	var m map[int]int
	m = make(map[int]int)
	fmt.Println(m[1])
	fmt.Println(m == nil)
	m[0] = 100
	fmt.Println(m[0])
}
