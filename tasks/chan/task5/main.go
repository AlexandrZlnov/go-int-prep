// --------------------------------------------------------------------------------------------------
// По видео канала Skill Issue не тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s
// предотвращение утечки го рутин

package main

import (
	"context"
	"fmt"
	"time"
	//"time"
)

func main() {
	ctx, cansel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cansel()

	ch := make(chan int)

	// тут select завершит горутину по контексту
	go func() {
		for i := range 10000 {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
		close(ch)
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println("v =", v)
		case <-ctx.Done():
			return
		}
	}

}
