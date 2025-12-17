// --------------------------------------------------------------------------------------------------
// из списка задач к собеседования - Avtostopom-po-Go
//Написать код, который будет выводить
//коды ответов на HTTP-запросы по двум URL
//главная страница Google и главная страница WB.
//Запросы должны осуществляться параллельно.

//Решение 1:

package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	url  string
	code int
	err  error
}

func main() {
	urls := []string{"https://google.ru", "https://wildberries.ru"}

	var wg sync.WaitGroup

	results := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchStatus(url, results, &wg)

	}

	for i := 0; i < len(urls); i++ {
		res := <-results
		if res.err != nil {
			fmt.Printf("---> %s ошибка: %s\n", res.url, res.err.Error())

		} else {
			fmt.Println(res)
		}
	}

	wg.Wait()

}

func fetchStatus(url string, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	var resp Result
	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		resp = Result{url, 0, err}
		results <- resp
		return
	}

	resp = Result{url, res.StatusCode, err}

	results <- resp

	return
}

// Решение 2
/*
package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{"https://google.ru", "https://wildberries.ru"}

	ctx, cansel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))

	defer cansel()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			checkURL(url, ctx)
		}(url)
	}

	wg.Wait()

}

func checkURL(url string, ctx context.Context) {
	req, err := http.NewRequestWithContext(ctx, "Get", url, nil)

	if err != nil {
		fmt.Printf("Ошибка формирования запроса для %s: %s\n", url, err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка получения ответа для %s: %s\n", url, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("URL: %s, Код ответа: %d\n", url, resp.StatusCode)

	return
}
*/
