// Привести примеры 3х deadlock кейсов с каналами

package main

import "fmt"

func main() {
	// кейс 1: запись или чтение в nil канал
	var ch chan int
	ch <- 55
	fmt.Println(<-ch)

	// кейс 2: запсиь в небуфиризированный канал без чтения
	// ch := make(chan int)
	// ch <- 55

	// кейс 3: чтение из канала без записи
	// ch := make(chan int)
	// val := <-ch
	// fmt.Println(val)
}
