// Задача
// Что выведет код?

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch1, ch2 := worker(), worker()

	<-ch1
	<-ch2

	fmt.Printf("time passed %v", time.Since(start))
}

func worker() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}

// Ответ:
// Вывод: воркеры вызваны с ингорированием возвращаемого значения.
// Возвращаемые каналы будут проигнорированы, чтение из них выполняться не будет.
// Программа выполнится почти мгновенно до завещения time Sleep в горутине.
// Распечатается времы выполнения близкое к нулю.
