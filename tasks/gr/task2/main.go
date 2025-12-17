// --------------------------------------------------------------------------------------------------
// из собеса на midle в авито
// задача на каналы и го рутины

package main

import (
	"time"
)

func main() {
	timeStart := time.Now()
	// var wg sync.WaitGroup
	// wg.Add(3)
	// for i := 0; i < 3; i++ {
	// 	go func() {
	// 		_ = <-worker()
	// 		defer wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// chans := []chan int{worker(), worker(), worker()}

	// for _, ch := range chans {
	// 	<-ch
	// }

	_, _, _ = worker(), worker(), worker()
	<-worker()

	println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- 1
	}()
	return ch
}
