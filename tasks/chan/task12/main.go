// Написать 2–3 кейса на порядок выполнения (что может напечататься и почему)

// кейс 1:
// конструкция select case не гарантирует последовательность выводоа данных
/*
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan string, 3)

	go func() {
		defer close(ch1)
		for i := 0; i < cap(ch1); i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 0; i < cap(ch2); i++ {
			ch2 <- fmt.Sprintf("word_%d", i)
		}
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case i, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println(i)
		case w, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(w)
		}
	}

}
*/

// кейс 2:
// последовательность передачи данных гарантировано сохранятется
package main

import (
	"fmt"
)

func SomeWork(work []int, ch chan int) {
	for _, w := range work {
		ch <- w * w
	}
	close(ch)

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int)

	go SomeWork(arr, ch)

	for res := range ch {
		fmt.Println(res)
	}

}
