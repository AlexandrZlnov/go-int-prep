// Паттерн - Decorator
// Позволяет динамически добавлять поведение объекту, не изменяя его класс.
// Представлено 2 варианта реализации.

// -----------Вариант 1:
// Каждый слой реализует тот же интерфейс
// Каждый слой содержит следующий слой
// Можно оборачивать сколько угодно раз
// Поведение расширяется динамически

package main

import "fmt"

// Базовый интерфейс
// Любой тип, который реализует GetUser(int) string, может быть:
// - основной реализацией
// - декоратором
type UserService interface {
	GetUser(id int) string
}

// Основная реализация
// реально делает работу — ходит в БД (условно).
// и ничего не знает про логирование или кэш.
type DBUserService struct{}

func (u *DBUserService) GetUser(id int) string {
	fmt.Println("Выполням запроса по ID =", id)
	return fmt.Sprintf("DB: данные пользователя с ID = %d\n", id)
}

// LoggerDecorator
// Он содержит внутри next. Next — это тоже UserService
// внутри может быть:
// - DBUserService
// - CacheDecorator
// - другой Logger
// - что угодно
type LoggerDecorator struct {
	next UserService
}

func (l *LoggerDecorator) GetUser(id int) string {
	fmt.Println("LOG: Запрос пользователя с ID =", id)
	result := l.next.GetUser(id)
	fmt.Println("LOG: Результат =", result)
	return result
}

// CacheDecorator
// реализует UserService
// содержит next
// добавляет поведение
type CacheDecorator struct {
	next  UserService
	cache map[int]string
}

func (c *CacheDecorator) GetUser(id int) string {
	if val, ok := c.cache[id]; ok {
		fmt.Println("CACHE: ID найден в кэш")
		return val
	}

	fmt.Println("CACHE: ID не найден в кэш")
	result := c.next.GetUser(id)
	c.cache[id] = result

	return result
}

func NewCacheDecorator(next UserService) *CacheDecorator {
	return &CacheDecorator{
		next:  next,
		cache: make(map[int]string),
	}
}

func main() {
	db := &DBUserService{}

	withLogging := &LoggerDecorator{next: db}

	withCache := NewCacheDecorator(withLogging)

	fmt.Println(withCache.GetUser(1))
	fmt.Println("-------")
	fmt.Println(withCache.GetUser(2))
	fmt.Println("-------")
	fmt.Println(withCache.GetUser(1))
}

/*
Вывод:
CACHE: ID не найден в кэш
LOG: Запрос пользователя с ID = 1
Выполням запроса по ID = 1
LOG: Результат = DB: данные пользователя с ID = 1

DB: данные пользователя с ID = 1

-------
CACHE: ID не найден в кэш
LOG: Запрос пользователя с ID = 2
Выполням запроса по ID = 2
LOG: Результат = DB: данные пользователя с ID = 2

DB: данные пользователя с ID = 2

-------
CACHE: ID найден в кэш
DB: данные пользователя с ID = 1
*/

//
// -----------Вариант 2:
// Классическая реализация с определением общей структуры Декоратора
// которая не нужна в данном случае
/*
package main

import "fmt"

// Component
type Notifier interface {
	Send(message string)
}

// ConcreteComponent
type EmailNotifier struct {
}

func (d *EmailNotifier) Send(message string) {
	fmt.Println("Sending email: ", message)
}

// Decorator
// type NotifierDecorator struct {
// 	wrapper Notifier
// }

// func (d *NotifierDecorator) Send(message string) {
// 	d.wrapper.Send(message)
// }

// ConcreteDecorator
type LoggingDecorator struct {
	Notifier
}

func (l *LoggingDecorator) Send(message string) {
	fmt.Println("Log: sending message")
	l.Notifier.Send(message)
}

// ConcreteDecorator
type RetryDecorator struct {
	Notifier
}

func (r *RetryDecorator) Send(message string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Attempt %d:\n", i+1)
		r.Notifier.Send(message)
	}
}

func main() {
	email := &EmailNotifier{}

	logging := &LoggingDecorator{Notifier: email}

	retry := &RetryDecorator{Notifier: logging}

	retry.Send("Hello")
}
*/

/*
Вывод:
Attempt 1:
Log: sending message
Sending email:  Hello
Attempt 2:
Log: sending message
Sending email:  Hello
Attempt 3:
Log: sending message
Sending email:  Hello
*/
