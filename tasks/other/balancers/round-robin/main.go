// Секция кодинг (1 час 1 задача)
// Задача написать balancer
// Round Robin - равномерно отправлять задачи по кругу всем воркерам.
// Не учитывает загрузку воркеров

package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID int
}

type Worker struct {
	ID int
	ch chan Task
}

func NewWorker(id int) *Worker {
	w := &Worker{
		ID: id,
		ch: make(chan Task),
	}

	go func() {
		for task := range w.ch {
			fmt.Printf("Worker %d обрабатывает задачу %d\n", w.ID, task.ID)
		}
	}()

	return w
}

type Balancer struct {
	workers []*Worker
	next    int
	mu      sync.Mutex
}

func NewBalancer(w []*Worker) *Balancer {
	return &Balancer{workers: w}
}

func (b *Balancer) Dispatch(task Task) {
	b.mu.Lock()
	worker := b.workers[b.next]
	b.next = (b.next + 1) % len(b.workers)
	b.mu.Unlock()

	worker.ch <- task
}

func main() {
	workers := []*Worker{
		NewWorker(1),
		NewWorker(2),
		NewWorker(3),
	}

	b := NewBalancer(workers)

	for i := 1; i <= 10; i++ {
		b.Dispatch(Task{ID: i})
	}
}
