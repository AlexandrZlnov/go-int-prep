// Задача:
// Что выведет код?

package main

import "fmt"

func main() {
	foo := 5
	bar := &foo
	fmt.Println(*bar) // 5

	changePointer(bar)
	fmt.Println(*bar) // 5
}

func changePointer(p *int) {
	newP := 3
	p = &newP
}
