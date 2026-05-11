// Задача
// Реализовать `func WorkerPool(jobs []int, workers int) []int`
//    - фиксированное число workers
//    - jobs через channel
//    - graceful shutdown
//    - Написать тесты на: корректность результатов, отсутствие зависаний, отмену
//    - Ограниченное число workers
//    - Должен быть results channel

// Базовое решение.
// Без context

package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Index int
	Value int
}

type Result struct {
	Index int
	Value int
}

func worker(jobCh <-chan Job, resultCh chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobCh {
		resultCh <- Result{job.Index, job.Value * 2}
	}

}

func WorkerPool(jobs []int, workers int) []int {
	if workers <= 0 {
		workers = 1
	}

	jobsCh := make(chan Job)
	resultCh := make(chan Result, workers)

	result := make([]int, len(jobs))

	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(jobsCh, resultCh, &wg)

	}

	go func() {
		defer close(jobsCh)
		for i, job := range jobs {
			jobsCh <- Job{i, job}
		}
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		result[res.Index] = res.Value
	}

	return result
}

func main() {

	jobs := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	const workers = 3

	result := WorkerPool(jobs, workers)

	fmt.Println("Результать таботы:", result)

}

// Усложнение:
// Добавить ctx cancellation
/*
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Index int
	Value int
}

type Result struct {
	Index int
	Value int
}

func worker(
	ctx context.Context,
	jobs <-chan Job,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			result := Result{
				Index: job.Index,
				Value: job.Value * 2,
			}

			select {
			case <-ctx.Done():
				return

			case results <- result:
			}
		}
	}
}

func WorkerPoolCtx(
	ctx context.Context,
	jobs []int,
	workers int,
) []int {
	if workers <= 0 {
		workers = 1
	}

	jobsCh := make(chan Job)
	resultsCh := make(chan Result)

	results := make([]int, len(jobs))

	var wg sync.WaitGroup

	// workers
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(ctx, jobsCh, resultsCh, &wg)
	}

	// producer
	go func() {
		defer close(jobsCh)

		for i, job := range jobs {
			select {
			case <-ctx.Done():
				return

			case jobsCh <- Job{
				Index: i,
				Value: job,
			}:
			}
		}
	}()

	// closer
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	// aggregator
	for {
		select {
		case <-ctx.Done():
			return results

		case result, ok := <-resultsCh:
			if !ok {
				return results
			}

			results[result.Index] = result.Value
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Millisecond,
	)
	defer cancel()

	jobs := []int{
		10, 20, 30, 40, 50,
		60, 70, 80, 90,
	}

	results := WorkerPoolCtx(ctx, jobs, 3)

	fmt.Println(results)
}
*/
