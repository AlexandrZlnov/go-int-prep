// Singleton:
// гарантирует, что у структуры существует только один экземпляр,
// и предоставляет к нему глобальную точку доступа.

// Вариант 1 ---------------------------
// Правильный, канонический, с защитой от rece condition

package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value int
}

var (
	instance *Singleton
	once     sync.Once
	wg       sync.WaitGroup
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			value: 33,
		}
	})
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(grNum int) {
			defer wg.Done()
			inst := GetInstance()
			fmt.Printf("Gorutine number = %d, instance addr = %p\n", grNum, inst)
		}(i)
	}
	wg.Wait()
}

// Вариант 2 ---------------------------
// Плохой вариант с double-check locking и
// с race condition в при первой провеке на nil
// в случае конкурентного вызова
// проверить можно через go run -race main.go
/*

package main

import (
	"fmt"
	"sync"
)

var mu = &sync.Mutex{}

type config struct {
	//config variables
}

var counter int = 1
var singleConfigInstance *config

func getConfigInstance() *config {
	if singleConfigInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleConfigInstance == nil {
			fmt.Println("Creating single instance now, and counter =", counter)
			singleConfigInstance = &config{}
			counter = +1
		} else {
			fmt.Println("Single instance alredy crated - 1, returning that one")
		}
	} else {
		fmt.Println("Single instance alredy crated - 2, returning that same")
	}
	return singleConfigInstance
}

func main() {
	for i := 0; i < 10; i++ {
		getConfigInstance()
	}
	fmt.Scanln()
}
*/
