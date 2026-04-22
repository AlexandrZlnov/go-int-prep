package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				fmt.Println("goroutine", id, "работает на потоке")
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	for {
		fmt.Println("горутин сейчас:", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}
