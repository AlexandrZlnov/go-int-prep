// --------------------------------------------------------------------------------------------------
// из списка задач к собеседования - Avtostopom-po-Go
//1. Иногда приходят нули. В чем проблема? Исправь ее
//2. Если функция bank_network_call выполняется 5 секунд,
//то за сколько выполнится balance()? Как исправить проблему?
//3. Представим, что bank_network_call возвращает ошибку дополнительно.
//Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.

// func balance() int {
// 	x := make(map[int]int, 1)
// 	var m sync.Mutex

// 	// call bank
// 	for i := 0; i < 5; i++ {
// 		i := i
// 		go func() {
// 			m.Lock()
// 			b := bank_network_call(i)
// 			x[i] = b
// 			m.Unlock()
// 		}()
// 	}

//		// Как-то считается сумма значений в мапе и возвращается
//		return sumOfMap
//	}

package main

import (
	"context"
	"sync"
)

func balance() (int, error) {
	x := make(map[int]int, 1)
	var (
		m         sync.Mutex
		wg        sync.WaitGroup
		errReturn error
		sOnce     sync.Once
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// call bank
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			b, err := bank_network_call(i)
			if err != nil {
				sOnce.Do(func() {
					errReturn = err
					cancel()
				})
				return
			}

			select {
			case <-ctx.Done():
				return
			default:
			}

			m.Lock()
			x[i] = b
			m.Unlock()
		}()
	}
	wg.Wait()
	if errReturn != nil {
		return 0, errReturn
	}
	// Как-то считается сумма значений в мапе и возвращается
	return sumOfMap, nil
}
