// мерджер нескольких каналов в 1

package main

//"fmt"
import (
	"sync"
)

// "golang.org/x/text/number"

func mergeChan(cs ...<-chan int) <-chan int {
	//number := 0
	out := make(chan int, len(cs))

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, v := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(v)
	}
	go func() {
		wg.Wait()
	}()
	return out
}

func main() {

}
