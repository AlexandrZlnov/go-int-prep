// функция дожидается закрытия обоих каналов
// версия с Wait.group

package main

import "sync"

func WaitChannels(a, b chan struct{}) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for range a {
		}
		wg.Done()
	}()

	go func() {
		for range b {
		}
		wg.Done()
	}()

	wg.Wait()
}

// функция дожидается закрытия обоих каналов
// версия с select case
/*
package main

func WaitChannels(a, b chan struct{}) {
	doneA, doneB := false, false

	for !(doneA && doneB) {
		select {
		case _, ok := <-a:
			if !ok {
				doneA = true
			}
		case _, ok := <-b:
			if !ok {
				doneB = true
			}
		}

	}
}

func main() {
	var a, b chan struct{}

	WaitChannels(a, b)

}
*/
