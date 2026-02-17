// техскрин МТС
// что выведет код

package main

import "fmt"

func main() {
	f1 := func(x []int) {
		x[0]++
	}

	f2 := func(x []int) {
		x = append(x, 100)
		x[0]++
	}

	f3 := func(x *[]int) {
		(*x)[0]++
	}

	f4 := func(x *[]int) {
		*x = append(*x, 100)
		(*x)[0]++
	}

	a := []int{1}

	fmt.Println("len", "item[0]")

	f1(a)
	fmt.Println(len(a), a[0]) // 1 2

	f2(a)
	fmt.Println(len(a), a[0]) // 1 2

	f3(&a)
	fmt.Println(len(a), a[0]) // 1 3

	f4(&a)
	fmt.Println(len(a), a[0]) // 2 4
}
