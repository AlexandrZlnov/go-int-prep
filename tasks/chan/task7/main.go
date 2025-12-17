// --------------------------------------------------------------------------------------------------
// из списка задач к собеседованию - Avtostopom-po-Go
//1. Merge n channels
//2. Если один из входных каналов закрывается,
//то нужно закрыть все остальные каналы
//func case3(channels ...chan int) chan int {}

package main

import (
	"context"
	"fmt"
	"sync"
)

//"fmt"

func case3(parentCtx context.Context, channels ...chan int) chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	// создаём локальный контекст, который можно отменить при закрытии любого входного канала
	ctx, cancel := context.WithCancel(parentCtx)

	// Обработчик каждого входного канала
	for _, in := range channels {
		in := in // захватываем переменную цикла
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-ch:
					if !ok {
						// один канал закрылся → сигнализируем остальным
						cancel()
						return
					}
					select {
					case out <- v:
						// успешно отправили
					case <-ctx.Done():
						// если кто-то уже закрыл пайплайн
						return
					}
				case <-ctx.Done():
					// сигнал завершения от другого канала
					return
				}
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	out := case3(a, b)

	go func() {
		a <- 1
		a <- 2
		close(a) // закрытие одного из каналов остановит весь merge
	}()

	go func() {
		b <- 10
		b <- 20
		// b не закрываем, merge всё равно остановится
	}()

	for v := range out {
		fmt.Println(v)
	}

}
