// По видео канала Skill Issue не тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s

package main

import (
	"fmt"
	"sync"
)

func Writer() chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 11
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	ch := Writer()

	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("v = ", v, "ok= ", ok)
	// }

	// или

	for v := range ch {
		fmt.Println("v = ", v)
	}

}
