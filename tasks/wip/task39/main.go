// Задача:
// Что выведет код?

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 42

	go func() {
		ch <- 100
	}()

	time.Sleep(time.Millisecond * 100)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Ответ:
// 42
// 100
