// Задание:
// Реализовать `func fetch(ctx context.Context) (int, error)` (100ms, отмена)
// Запустить 3 fetch параллельно, отменить через timeout
// Добавить логирование: какие goroutine успели завершиться, какие отменились

// Вариант 1
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func fetch(ctx context.Context, id int) (int, error) {
	select {
	case <-time.After(50 * time.Millisecond):
		result := id * 10
		log.Printf("Fetch %d completed.", id)
		return result, nil
	case <-ctx.Done():
		log.Printf("Fetch %d cancelled: %v", id, ctx.Err())
		return 0, ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100))
	defer cancel()

	var wg sync.WaitGroup
	const workers = 3

	results := make(chan int, 3)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			result, err := fetch(ctx, id)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					log.Printf("Goroutine %d stopped by Timeout: %v", id, err)
				}
				return
			}
			results <- result
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0

	for r := range results {
		fmt.Printf("Result received: %d\n", r)
		total += r
	}

	fmt.Println("TOTAL:", total)

}

// Ваариант 2
/*
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func fetch(ctx context.Context, id int) (int, error) {
	work := 10

	timer := time.After(100 * time.Millisecond)

	for {
		work *= id

		select {
		case <-timer:
			log.Printf("Timer stop. Fetch: %d.\n", id)
			return work, nil
		case <-ctx.Done():
			log.Printf("Timeout stop. Fetch: %d.\n", id)
			return work, ctx.Err()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100))
	defer cancel()
	result := make(chan int, 3)

	total := 0

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			work, err := fetch(ctx, id)
			if err != nil {
				fmt.Printf("Fetch %d, on wokr: %d, canceled: %v\n", id, work, err)
			} else {
				log.Printf("Fetch %d completed on work: %d", id, work)
			}
			result <- work
		}(i)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for w := range result {
		total += w

	}

	fmt.Println("TOTAL:", total)

}
*/
