// Задача:
// Разбери код.
// Исправь ошибки.

// Исходный код.
/*
package main

import (
	"sync"
)

func main() {
	const n = 1 << 10

	var (
		balance = 100
		wg      sync.WaitGroup
	)

	add := func(v int) {
		defer wg.Done()
		balance += v
	}

	wg.Add(2 * n)

	for i := 0; i < n; i++ {
		go add(1)
	}
}
*/

// Ответ:
// Код создасть 1024 горутины
// Код будет пытаться асинхронно увеличивать баланс на 1
// Проблемы:
// - data race
// - нет wg.Wait()

// Решение:

package main

import (
	"fmt"
	"sync"
)

func main() {
	const n = 1 << 10

	var (
		balance = 100
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	add := func(v int) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		balance += v
	}

	wg.Add(2 * n)

	for i := 0; i < n; i++ {
		go add(1)
		go add(-1)
	}

	wg.Wait()
	fmt.Println(balance)
}
