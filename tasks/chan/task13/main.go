package main

import (
	"fmt"
	"reflect"
)

func main() {
	chSend := make(chan<- int, 1)

	chSend <- 10

	fmt.Println(reflect.TypeOf(chSend).ChanDir())

	close(chSend)

	fmt.Println("done")

}
