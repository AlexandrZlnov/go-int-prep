// Задача:
// Условие:
// Есть:
// - список задач []int
// - нужно обработать их параллельно
// но:
// - максимум 3 горутины одновременно
// - результат нужно сохранить в slice
// Требования:
// результат строго по индексу input
// без race conditions (go test -race должен быть clean)
// корректный shutdown всех goroutine
// нельзя глобальные переменные

/*
Исходный код:
func process(n int) int {
    return n * 2
}

func main() {
    tasks := []int{1,2,3,4,5,6,7,8,9,10}

    results := make([]int, len(tasks))

    // TODO: обработать параллельно с лимитом 3 goroutine

    fmt.Println(results)
}
*/

package main

import (
	"fmt"
	"sync"
)

func process(n int) int {
	return n * 2
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	results := make([]int, len(tasks))

	// TODO: обработать параллельно с лимитом 3 goroutine
	var wg sync.WaitGroup
	maxWorkers := 3
	jobs := make(chan struct {
		ind   int
		value int
	})

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results[job.ind] = process(job.value)
			}
		}()
	}

	for ind, value := range tasks {
		jobs <- struct {
			ind   int
			value int
		}{ind, value}
	}

	close(jobs)
	wg.Wait()

	fmt.Println(results)
}
