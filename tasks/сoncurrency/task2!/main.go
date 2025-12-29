// Задача с собеса в X5 - 18.12.2025

// Реализуйте структуру "once", функцию "new" и потокобезопасный метод "do".

// Реализация "once" должна использовать каналы, не использовать пакет "sync".
// Функция "new" должна возвращать указатель на структуру "once".
// Метод "do":
//   - получает на вход функцию "f",
//   - выполняет функцию "f" только в том случае, если "do" вызывается в "первый раз" для данного экземпляра.

// Функция "main" должна вывести "call" в консоль ровно один раз.

// Исходный код к заданию:
/*
package main

import (
	"fmt"
	"sync"
)

const goroutinesNumber = 10

type once struct {
}

func new() *once {
}

func (o *once) do(f func()) {
}

func funcToCall() {
	fmt.Println("call")
}

func main() {
	wg := sync.WaitGroup{}
	so := new()

	wg.Add(goroutinesNumber)
	for i := 0; i < goroutinesNumber; i++ {
		go func() {
			defer wg.Done()
			so.do(funcToCall)
		}()
	}

	wg.Wait()
}
*/

// Решение:
// Не утверждено.
// Пояснение:
// Мы используем буферизированный канал ёмкостью 1 как семафор.
// В new мы кладём в канал один токен.
// Первая горутина успешно забирает токен и выполняет функцию.
// После этого канал пуст, и все последующие вызовы do не могут получить токен и сразу выходят через default.
// Таким образом, функция выполняется ровно один раз.

package main

import (
	"fmt"
	"sync"
)

const goroutinesNumber = 10

type once struct {
	ch chan struct{}
}

func new() *once {
	o := &once{
		ch: make(chan struct{}, 1),
	}
	o.ch <- struct{}{}
	return o

}

func (o *once) do(f func()) {
	select {
	case <-o.ch:
		f()
	default:
	}
}

func funcToCall() {
	fmt.Println("call")
}

func main() {
	wg := sync.WaitGroup{}
	so := new()

	wg.Add(goroutinesNumber)
	for i := 0; i < goroutinesNumber; i++ {
		go func() {
			defer wg.Done()
			so.do(funcToCall)
		}()
	}

	wg.Wait()
}
