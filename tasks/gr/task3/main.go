// Собес: МТС
// Задача:
// Исправь код так что бы получить
// горутину и ее номер

// Исходный код:
/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string) // нельзя использовать буферизированный канал
	mu := sync.Mutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(group sync.WaitGroup, i int) {
			defer group.Done()

			mu.Lock()
			ch <- fmt.Sprintf("Goroutine %d", i)
			mu.Unlock()
		}(wg, i)
	}

	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		}
	}

	wg.Wait()
}

*/

// Решение:

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string) // нельзя использовать буферизированный канал
	//mu := sync.Mutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(group *sync.WaitGroup, i int) {
			defer group.Done()

			//mu.Lock()
			ch <- fmt.Sprintf("Goroutine %d", i)
			//mu.Unlock()
		}(&wg, i)
	}

	// добавим close
	go func() {
		wg.Wait()
		close(ch)
	}()

	// добавим выход из цикла
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(s)
		}

	}

}
