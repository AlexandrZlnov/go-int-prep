/*
Создание конструктора объекта с функциональными опциями
Задание
Создай конструктор объекта User с функциональными опциями.
В коде уже представлен пример реализации конструктора с функциями опций.
Дополни код, чтобы конструктор правильно инициализировал объект User с заданными параметрами.

package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}

Критерии заверешнности:
- Конструктор NewUser правильно инициализирует объект User с заданными параметрами.
- Параметр ID установлен в переданное значение id.
- Параметр Username установлен в переданное значение username.
- Параметр Email установлен в переданное значение email.
- Параметр Role установлен в переданное значение role.
*/

package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(val string) UserOption {
	return func(u *User) {
		u.Username = val
	}
}

func WithEmail(val string) UserOption {
	return func(u *User) {
		u.Email = val
	}
}

func WithRole(val string) UserOption {
	return func(u *User) {
		u.Role = val
	}
}

func NewUser(id int, options ...UserOption) *User {
	user := &User{
		ID: id,
	}

	for _, opt := range options {
		opt(user)
	}
	return user
}

func main() {
	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
	fmt.Printf("User: %+v\n", user)
}
