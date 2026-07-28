// Задача:
// Собес: Магнит
// Смержить произвольное количество каналов в 1 канал.
// Исходный код:
/*
func merge(chList ...chan int) chan int {
}

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	out := merge(ch1, ch2)

	for v := range out {
		fmt.Println(v)
	}
}
*/

// Решение:
package main

import (
	"fmt"
	"sync"
)

func merge(chList ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	out := merge(ch1, ch2)

	for v := range out {
		fmt.Println(v)
	}
}
