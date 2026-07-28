// Задача:
// Что выведет код?

package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// Ответ:
// В коде присутствует data race, результат выполнения будет меняться
// Ожидается число стремящяеся к 999.
