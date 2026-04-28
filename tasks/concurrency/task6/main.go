// Вопросы:
// - Есть ли здесь race condition
// - Что может произойти в runtime (конкретно)
// - Как это исправить (минимум 2 способа)

//Исходный код:
/*
package main

func main() {
	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			m[i] = i
		}(i)
	}
}
*/

// Ответ:
// - race condition есть, несколько горутин пишут в одну мапу
// - ошибка: одновременная запись в мапу
// - ниже 3 варианта реализации

// Ответ Вариант 1:

package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}

// Ответ Вариант 2:
/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait()

	m.Range(func(key, val any) bool {
		fmt.Printf("%v:%v ", key, val)
		return true
	})

}
*/

// Ответ Вариант 3 без примитивов синхронизации (плохой вариант):
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)
	ch := make(chan struct {
		k int
		v int
	}, 1000)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			ch <- struct {
				k int
				v int
			}{i, i}
		}(i)

	}

	time.Sleep(time.Microsecond * 50)
	close(ch)

	for kv := range ch {
		m[kv.k] = kv.v
	}

	fmt.Println(len(m))

}
*/
