// --------------------------------------------------------------------------------------------------
// По видео канала Skill Issue не тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s
// Задача: написать функции: writer-генерит числа от 1-10 double умножает их на 2 reader читает из dubler

package main

import (
	"fmt"
	"time"
)

func writer() chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	return ch
}

func double(ch chan int) chan int {
	chW := make(chan int)

	go func() {
		defer close(chW)
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			chW <- v * 2
		}
	}()

	return chW
}

func reader(chW chan int) {
	for v := range chW {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Double v =", v)
	}
}

func main() {
	reader(double(writer()))
}
