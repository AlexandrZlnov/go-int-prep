// есть сотрудники делающие бургеры, код учитывает сколько они сделали бургеров
// по жалобе клиента, у них могут уменьшать счетчик бургеров

// Исходный код
/*
var workers map[string]uint

func makeBurger(workerName string) {
    workers[workerName] = workers[workerName]++
}

func rejectBurger(workerName string) {
    workers[workerName] = workers[workerName]
}

func currentCount(workerName string) {
    fmt.Printf("Worker: %s made %d burgers\n", workerName, workers[workerName])
}
*/

// Ответ:

// есть сотрудники делающие бургеры, код учитывает сколько они сделали бургеров
// по жалобе клиента, у них могут уменьшать счетчик бургеров

package main

import (
	"fmt"
	"sync"
)

// инициализируем мапу
var workers = make(map[string]uint)

// добавить Mutex
var mu sync.RWMutex

func makeBurger(workerName string) {
	//workers[workerName] = workers[workerName]++
	mu.Lock()
	defer mu.Unlock()

	workers[workerName]++
}

func rejectBurger(workerName string) {
	mu.Lock()
	defer mu.Unlock()
	// необходима проверка знечение не может быть меньше 0
	if workers[workerName] > 0 {
		workers[workerName]--
	}
}

func currentCount(workerName string) {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Printf("Worker: %s made %d burgers\n", workerName, workers[workerName])
}
