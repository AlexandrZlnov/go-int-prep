// Задание
// Разобрать гонку и починить 3 способами (channel/mutex/atomic):
//
// Требования:
// - нельзя менять сигнатуру функций
// - нелья убрать горутину, x= и Print должны выполняться в горутинах обязательно
// - нельзя совмещать методы, можно: только каналы, только mutex, только atomic
// - нельзя time.Sleep
// - должен быть чистый go run -race main.go
// - вывод всегда должен быть 1
// - можно spin loop
//
// Исходный код:
/*
var x int
go func() { x = 1 }()
go func() { fmt.Println(x) }()
*/

package main

import (
	"fmt"
	//"sync/atomic"
	//"sync"
)

func main() {

	// -------- Решение с каналами -------
	var x int
	ch := make(chan struct{})
	done := make(chan struct{})

	go func() {
		x = 1
		ch <- struct{}{}
	}()

	go func() {
		<-ch
		fmt.Println(x)
		close(done)
	}()
	<-done

	// -------- Решение через mutex -------
	/*
		var (
			x    int
			mu   sync.Mutex
			done int = 2
		)

		go func() {
			mu.Lock()
			x = 1
			done -= 1
			mu.Unlock()
		}()

		go func() {
			for {
				mu.Lock()
				if done != 1 {
					mu.Unlock()
					continue
				}
				break
			}
			fmt.Println(x)
			done -= 1
			mu.Unlock()
		}()

		for {
			mu.Lock()
			if done != 0 {
				mu.Unlock()
				continue
			}
			break
		}
	*/

	// -------- Решение через atomic -------
	/*
		var (
			x    int32
			done int32
		)

		go func() {
			atomic.AddInt32(&x, 1)
		}()

		go func() {
			for atomic.LoadInt32(&x) != 1 {
			}

			fmt.Println(atomic.LoadInt32(&x))

			atomic.StoreInt32(&done, 1)
		}()

		for atomic.LoadInt32(&done) != 1 {
		}
	*/

}
