// Задача
// Реализуй worker pool:
// Условие:
// - есть N воркеров
// - есть входной канал задач jobs
// - каждый воркер обрабатывает задачи и пишет результат в results
// - нужно корректно завершить все горутины

// Исходный код:
/*
type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	// TODO
}

func main() {
	jobs := make(chan Job)
	results := make(chan Result)

	// TODO: start workers

	// TODO: send jobs

	// TODO: close + collect results
}
*/

// Решение:

package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	// TODO
	defer wg.Done()
	for job := range jobs {
		results <- Result{
			JobID: job.ID,
			Value: job.ID * 10,
		}
		fmt.Printf("Worker %d выполни задачу %d\n", id, job.ID)
	}

}

func main() {
	jobs := make(chan Job)
	results := make(chan Result)

	// TODO: start workers
	var wg sync.WaitGroup

	const numWorkers = 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)

	}

	// TODO: send jobs
	const numJobs = 10
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{
				ID: j,
			}

		}
		close(jobs)
	}()

	// TODO: close + collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Результать задачи %d = %d\n", result.JobID, result.Value)
	}

}
