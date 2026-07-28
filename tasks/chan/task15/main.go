// Задача:
// 1) За какое время отработает этот код?
// 2) Как сократить время выполения кода?

package main

import (
	"fmt"
	"time"
)

func doSmt() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 111
	}()

	return ch
}

func main() {
	timeStart := time.Now()

	_, _ = <-doSmt(), <-doSmt()

	fmt.Println(int(time.Since(timeStart).Seconds()))
}

// Ответ:
// 1) Отработает за время немного более 4х секунд. Тк doSmt вызывается последовательно
// 2)   ch1 := doSmt()
//  	ch2 := doSmt()
//		<-ch1
// 		<-ch2
