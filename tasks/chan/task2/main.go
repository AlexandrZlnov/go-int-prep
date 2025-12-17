// из списка задач к собеседования - Avtostopom-po-Go
// func main() {
// 	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

// 	fmt.Println(callRequestsForURLs(urls, 3))
// }

// // дернуть N урлов с лимитом K (то есть не больше K активных запросов одновременных), сигнатура:
// func callRequestsForURLs(urls []string, K int) []*http.Response {}

//Решение:

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://yandex.ru", "https://yandex.ru", "https://yandex.ru", "https://yandex.ru", "https://yandex.ru", "https://yandex.ru", "https://yandex.ru"}
	fmt.Println(callRequestsForURLs(urls, 3))

}

func callRequestsForURLs(urls []string, K int) []*http.Response {
	var wg sync.WaitGroup
	var mu sync.Mutex

	response := make([]*http.Response, 0, len(urls))
	maxcall := make(chan struct{}, K)

	wg.Add(len(urls))

	for _, url := range urls {

		url := url

		maxcall <- struct{}{}

		go func() {
			defer func() {
				<-maxcall
				wg.Done()
			}()

			client := http.Client{
				Timeout: 2 * time.Second,
			}

			res, err := client.Get(url)
			if err != nil {
				return
			}

			mu.Lock()
			response = append(response, res)
			mu.Unlock()

		}()

	}

	wg.Wait()

	return response

}
