// Собес: Cloude
// Задача:
// Написать асинхронный обработчик задач в виде библиотеки.
// - Клиент библиотеки передаёт на вход некоторый объект (Task), у которого есть метод Do().
// - Клиент может передать несколько задач обработчику, задачи передаются:
// 		- по одной;
// 		- задачи могут передаваться из разных рутин.
// - Как только клиент передаст все задачи обработчику, он закрывает
// обработчик и ждёт завершения выполнения всех переданных задач.
// - Обработчик внутри имеет свою очередь, откуда берёт задачи на выполнение.
// - Очередь обработчика ограничена размером X; если при добавлении задачи
// она заполнена, обработчик сразу возвращает ошибку.
// - Обработчик выполняет не более N задач одновременно.

// Решение:

package main

import (
	"errors"
	"sync"
)

var (
	ErrQueueFull error = errors.New("Очередь заполнена")
	ErrClosed    error = errors.New("Обработчик закрыт")
)

type Task interface {
	Do()
}

type Handler interface {
	AddTask(task Task) error
	Close()
}

type handler struct {
	mu sync.Mutex
	wg sync.WaitGroup

	closed bool

	queue chan Task
	once  sync.Once
}

func New(maxQueue int, maxWorkers int) Handler {
	h := &handler{
		queue: make(chan Task, maxQueue),
	}

	h.wg.Add(maxWorkers)
	for range maxWorkers {
		go h.Worker()
	}

	return h
}

func (h *handler) Worker() {
	defer h.wg.Done()
	for task := range h.queue {
		task.Do()
	}

}

func (h *handler) AddTask(task Task) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.closed {
		return ErrClosed
	}

	select {
	case h.queue <- task:
		return nil
	default:
		return ErrQueueFull
	}

}

func (h *handler) Close() {
	h.once.Do(
		func() {
			h.mu.Lock()
			h.closed = true
			close(h.queue)
			h.mu.Unlock()

		})
	h.wg.Wait()

}
