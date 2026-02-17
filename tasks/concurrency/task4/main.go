// Собес: Магнит

// 1. Смотрим код до функции GetFile
// 1.1 Опиши что делает код?
// 2. Учитывая функцию GetFile, сколько будет выполняться код, точность +-20%?
// 3. Можно ли сократить время выполнения?
// 4. Рефакторинг кода:
// 4.1 Реазлизуй многопоточность.
// 	   Условие:
//	   - как только встречаем ошибку прекращаем обработку и возвращаем ошибку
//	   - обработать обшибки

/*
//-----------------------------------------------------------------------------------------------------
//<Исходный код:>
//
package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	m, err := GetFiles(context.TODO(), "1", "2", "3", "4", "5")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m)
	fmt.Println(time.Since(start))
}

// GetFiles пример функции, которую нужно оптимизировать.
func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetFile является примером функции, которая относительно
// недолго выполняется при единичном вызове. Но достаточно
// долго если вызывать последовательно.
// Предположим, что оптимизировать в этой функции нечего.
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	if name == "invalid" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}
//
// </Исходный код>
//-----------------------------------------------------------------------------------------------------
*/

package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	m, err := GetFiles(context.TODO(), "1", "2", "3", "4", "5")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m)
	fmt.Println(time.Since(start))
}

// GetFiles пример функции, которую нужно оптимизировать.
// Решение Вариант 1:
func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		errorNew error
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	result = make(map[string][]byte, len(names))

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			res, err := GetFile(ctx, name)
			if err != nil {
				mu.Lock()
				firstErr := errorNew == nil // дополнительная логика для выноса cancel из unlock
				if firstErr {
					errorNew = err
				}
				mu.Unlock()

				if firstErr {
					cancel()
				}
				return
			}
			mu.Lock()
			result[name] = res
			mu.Unlock()

		}(name)

	}

	wg.Wait()

	return result, errorNew
}

// Решение Вариант 2: c errgroup
/*
func GetFiles(ctx context.Context, names ...string) (map[string][]byte, error) {
    if len(names) == 0 {
        return nil, nil
    }

    g, ctx := errgroup.WithContext(ctx)

    result := make(map[string][]byte, len(names))
    var mu sync.Mutex

    for _, name := range names {
        name := name

        g.Go(func() error {
            res, err := GetFile(ctx, name)
            if err != nil {
                return err
            }

            mu.Lock()
            result[name] = res
            mu.Unlock()

            return nil
        })
    }

    if err := g.Wait(); err != nil {
        return nil, err
    }

    return result, nil
}
*/

// GetFile является примером функции, которая относительно
// недолго выполняется при единичном вызове. Но достаточно
// долго если вызывать последовательно.
// Предположим, что оптимизировать в этой функции нечего.
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	if name == "invalid" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}
