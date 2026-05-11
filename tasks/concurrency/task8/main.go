// Реализация worker pool

// Нужно написать:

// функцию:
// func WorkerPool(jobs []int, workers int) []int
// Требования:
// - есть список чисел

// каждый worker:
// - умножает число на 2

// нужно:
// - параллельно обработать jobs
// - сохранить результат
// - порядок НЕ важен
// - нельзя data race
// - нужно корректно завершать goroutines

//Вариант 1

package main

import (
	"fmt"
	"sync"
)

type Result struct {
}

func WorkerPool(jobs []int, workers int) []int {
	var wg sync.WaitGroup
	result := make([]int, len(jobs))
	jobCh := make(chan struct {
		i int
		v int
	})

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				result[job.i] = job.v * 2
			}

		}()

	}

	for i, job := range jobs {
		jobCh <- struct {
			i int
			v int
		}{
			i: i,
			v: job,
		}
	}

	close(jobCh)

	wg.Wait()

	return result

}

func main() {
	jobs := []int{2, 3, 4, 5, 6, 7, 8, 9}
	workers := 3

	result := WorkerPool(jobs, workers)

	fmt.Println(result)

}

// Вариант 2
// Интересный вариант но не workerpool с ограниченным количеством воркеро.
// Тут скорее ограничение параллелизма через семафор.
// Количество горутин будет = количеству элементов слайса.
// Но не более 3х одновременно.

/*
package main

import (
	"fmt"
	"sync"
)

func WorkerPool(jobs []int, workers int) []int {
	if workers == 0 {
		workers = 1
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, workers)
	result := make([]int, len(jobs))

	for i, job := range jobs {

		wg.Add(1)
		sem <- struct{}{}

		go func(i, job int) {
			defer wg.Done()
			result[i] = job * 2
			<-sem
		}(i, job)

	}
	wg.Wait()
	return result
}

func main() {
	jobs := []int{2, 3, 4, 5, 6, 7, 8, 9}
	workers := 3

	result := WorkerPool(jobs, workers)

	fmt.Println(result)

}
*/
