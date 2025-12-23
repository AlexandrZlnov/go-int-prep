/*
1. Сделать структуры Base и Child.
2. Структура Base должна содержать строковое поле name.
3. Структура Child должна содержать строковое поле lastName.
4. Сделать функцию Say у структуры Base, которая распечатывает на экране: Hello, %name!
5. Пронаследовать Child от Base.
6. Инициализировать экземпляр b1 Base. присвоить name значение Parent.
7. Инициализировать экземпляр c1 Child. присвоить name значение Child, присвоить lastName значение Inherited.
8. Вызвать у обоих экземпляров метод Say.
9. Переопределить метод Say для структуры Child, чтобы он выводил на экран: Hello, %lastName %name!
10. Сделать массив, содержащий b1 и c1.
11. Вызвать Say у всех элементов массива из шага 10.
12. Сделать метод NewObject для создания экземпляров Base и Child в зависимости от входного параметра.
13. Написать юнит-тесты для метода NewObject.
14. Сделать генератор объектов Base и Child такой, чтобы:
    - объекты Base создавались в фоновом потоке с задержкой 1 секунда;
    - объекты Child создавались в фоновом потоке с задержкой 2 секунды;
    - общее время генерации объектов не превышало 11 секунд.
15. Сделать асинхронный обработчик сгенерированных объектов такой, чтобы:
    - метод Say вызывался в порядке генерации объектов;
    - не приводил к утечкам памяти.
*/

// Решение:

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Base struct {
	name string
}

type Child struct {
	Base
	lastName string
}

func (b *Base) Say() {
	fmt.Printf("Hello, %s\n", b.name)

}

func (c *Child) Say() {
	fmt.Printf("Hello, %s %s\n", c.lastName, c.name)
}

type Sayer interface {
	Say()
}

func main() {
	b1 := Base{
		name: "Parent",
	}

	c1 := Child{
		Base:     Base{name: "Child"},
		lastName: "Inherited",
	}

	b1.Say()
	c1.Say()

	arr := [2]Sayer{&b1, &c1}

	for _, s := range arr {
		s.Say()
	}

	// 14. Сделать генератор объектов Base и Child такой, чтобы:
	// - объекты Base создавались в фоновом потоке с задержкой 1 секунда;
	// - объекты Child создавались в фоновом потоке с задержкой 2 секунды;
	// - общее время генерации объектов не превышало 11 секунд.
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()

	ch := make(chan Sayer, 10)

	wg.Add(2)

	go func(ctx context.Context, ch chan Sayer) {
		defer wg.Done()

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		i := 0

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ch <- NewObject("base", fmt.Sprintf("NameBase %d", i), "")
				i++
			}
		}
	}(ctx, ch)

	go func(ctx context.Context, ch chan Sayer) {
		defer wg.Done()

		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		j := 0

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ch <- NewObject("child", fmt.Sprintf("NameChild %d", j), fmt.Sprintf("lastNameChild %d", j))
				j++
			}
		}

	}(ctx, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	// 15. Сделать асинхронный обработчик сгенерированных объектов такой, чтобы:
	//     - метод Say вызывался в порядке генерации объектов;
	//     - не приводил к утечкам памяти.

	var wgConsumer sync.WaitGroup

	wgConsumer.Add(1)

	go func(ch chan Sayer) {
		defer wgConsumer.Done()
		for obj := range ch {
			obj.Say()
		}
	}(ch)

	wgConsumer.Wait()

}

// 12. Сделать метод NewObject для создания экземпляров Base и Child в зависимости от входного параметра.
func NewObject(objType, name, lastName string) Sayer {
	if objType == "base" {
		return &Base{
			name: name,
		}
	} else if objType == "child" {
		return &Child{
			Base:     Base{name: name},
			lastName: lastName,
		}
	}

	return nil
}
