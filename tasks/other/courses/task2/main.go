/*
Задание
Создание конструктора объекта с functional options
Необходимо реализовать конструктор объекта Order с использованием функциональных опций.

package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}

Критерии завершенности:
- Конструктор NewOrder создает объект типа Order с заданным ID.
- При использовании функциональных опций конструктор применяет соответствующие опции к создаваемому объекту.
- Если опции не переданы, соответствующие поля объекта Order остаются пустыми.
- Возвращаемое значение конструктора является указателем на объект Order с примененными опциями.

*/

package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(val string) OrderOption {
	return func(ord *Order) {
		ord.CustomerID = val
	}
}

func WithItems(val []string) OrderOption {
	return func(ord *Order) {
		ord.Items = val
	}
}

func WithOrderDate(val time.Time) OrderOption {
	return func(ord *Order) {
		ord.OrderDate = val
	}
}

func NewOrder(id int, options ...OrderOption) *Order {
	order := &Order{
		ID:         id,
		CustomerID: "default",
		Items:      []string{},
		OrderDate:  time.Now(),
	}

	for _, opt := range options {
		opt(order)
	}

	return order
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}
