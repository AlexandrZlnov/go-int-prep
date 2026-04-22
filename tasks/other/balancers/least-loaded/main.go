// Секция кодинг (1 час 1 задача)
// Задача написать balancer
// Least loaded - распределяет задачи по наименее загруженным воркерам
// Это упрощенное решение в котором есть изъян - оно не атомарно в части анализа нагрузки воркеров.
// Тк значение load может измениться в вроцессе перебора слайса воркеров

package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type Worker struct {
	ID   int
	ch   chan Task
	load int        // нагрузка
	mu   sync.Mutex // mutex для операций коррекции нагрузки
}

func NewWorker(id int) *Worker {
	w := &Worker{
		ID: id,
		ch: make(chan Task, 3), // добавлен размер буфера канала
	}

	go func() {
		for task := range w.ch {
			fmt.Printf("Worker %d обрабатывает задачу %d\n", w.ID, task.ID)

			time.Sleep(time.Millisecond * 100)

			w.Dec()
		}
	}()

	return w
}

// добавлены методы учета нагрзуки ворекеров
func (w *Worker) Inc() {
	w.mu.Lock()
	w.load++
	w.mu.Unlock()
}

func (w *Worker) Dec() {
	w.mu.Lock()
	w.load--
	w.mu.Unlock()
}

func (w *Worker) Load() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.load
}

type Balancer struct {
	workers []*Worker
	//next    int			// не нужен для реализации least loaded
	mu sync.Mutex
}

func NewBalancer(w []*Worker) *Balancer {
	return &Balancer{workers: w}
}

// измененна логика распределения относительно реализации round robin
// задача передается наименее нагруженному worker
func (b *Balancer) Dispatch(task Task) {
	var bestWorker *Worker

	minLoad := int(^uint(0) >> 1)

	for _, worker := range b.workers {
		if l := worker.Load(); l < minLoad {
			minLoad = l
			bestWorker = worker
		}
	}

	bestWorker.Inc()
	bestWorker.ch <- task
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

	time.Sleep(time.Second)
}
