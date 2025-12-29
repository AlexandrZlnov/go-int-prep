// По видео канала Skill Issue не тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s
// Запустить 5 воркеров для обработки входных данных и процессит с processData
// И результать processData писать в канал out
// И вся операция должна выполняться не более 5 секунд
// Если более 5 секунд то дропаем и возвращаем 0
//
// func processData(val int) int {
//     time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
//     return val * 2
// }

// func main() {
//     in := make(chan int)
//     out := make(chan int)

//     go func() {
//         for i := range 100 {
//             in <- i
//         }
//         close(in)
//     }()

//     now := time.Now()
//     processParallel(in, out, 5)

//     for val := range out {
//         fmt.Println(val)
//     }
//     fmt.Println(time.Since(now))
// }
// func processParallel(in <-chan int, out chan<- int, numWorkers int){
// }

//Решение
//есть сомнение в правильности правильности

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	now := time.Now()
	go processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}
func processParallel(in <-chan int, out chan<- int, numWorkers int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	worker := func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				result := processData(val)

				select {
				case <-ctx.Done():
					out <- 0
					return
				case out <- result:
				}
			}
		}
	}

	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	wg.Wait()
	close(out)

	// if ctx.Err() != nil {
	// 	return 0
	// }

	//return 1
}
