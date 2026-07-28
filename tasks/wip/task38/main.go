// Задача:
// Что выведет код?

package main

import "fmt"

type Counter struct {
	value int
}

func (c Counter) Increment() {
	c.value++
}

func (c *Counter) IncrementPtr() {
	c.value++
}

func main() {
	c := Counter{value: 0}
	c.Increment()
	fmt.Println(c.value)

	c.IncrementPtr()
	fmt.Println(c.value)
}

// Ответ:
// 0
// 1
