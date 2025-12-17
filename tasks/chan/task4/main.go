// По видео канала Skill Issue на тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s
// Каналы и select case

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// go func() {
	// 	ch1 <- 1
	// }()

	timer := time.NewTimer(1 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Microsecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("v =", v, "exited by ch1")
	case v := <-ch2:
		fmt.Println("v =", v, "exited by ch2")
	case <-time.After(1 * time.Second):
		fmt.Println("exited by after")
	case <-timer.C:
		fmt.Println("exited by timer")
	case <-ctx.Done():
		fmt.Println("exited by context")
		// default:
		// 	fmt.Println("exited by default")

	}

}
