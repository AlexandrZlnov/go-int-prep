// --------------------------------------------------------------------------------------------------
// По видео канала Skill Issue не тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s
// Задача: имеется функция которая работает неопределенно долго (до 100 секунд)
// func randomTimeWork() {
// time.Sleep(time.Duration(rand.Intn(100)) * time.Second)}
// написать обертку для этой функции, которая будет прерывать выполнение, если
// функция работает больше 3 секунд, и возвращать ошибку
// func predictableTimeWork() {}

package main

import (
	//"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Решение варант 2 (автора - доработанный)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() error {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("Работа выполнена за 3 секунды")
		return nil
	case <-time.After(3 * time.Second):
		return errors.New("Ожидание более 3 секунд")
	}
}

func main() {
	err := predictableTimeWork()
	if err != nil {
		fmt.Println(err)
	}
}

// Решение вариант 1 (мой)
/*
func randomTimeWork(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
	ch <- "work done"

}

func predictableTimeWork(ctx context.Context) error {
	ch := make(chan string)
	go randomTimeWork(ch)

	select {
	case <-ctx.Done():
		return errors.New("Прошло 3 секунды")
	case ms := <-ch:
		fmt.Println(ms)
		return nil
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := predictableTimeWork(ctx)
	if err != nil {
		fmt.Println(err)
	}

}
*/
