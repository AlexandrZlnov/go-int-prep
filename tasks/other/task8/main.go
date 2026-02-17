package main

import (
	"fmt"
)

func f(a []int) {
	a = append(a, 1)
}

func main() {
	a := make([]int, 10, 10)
	f(a)
	fmt.Println("Value a =", a)
}
