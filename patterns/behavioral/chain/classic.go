// Паттерн - Chain of Responsibility (Цепочка обязанностей)
// Пример Классической (упрощенной) реализации
//
// AuthHandler реализует интерфейс Handler, хотя напрямую определяет только метод Handle.
// Метод SetNext получается через embedding BaseHandler. Благодаря продвижению методов (method promotion)
// методы BaseHandler становятся доступными у AuthHandler, поэтому структура удовлетворяет интерфейсу Handler.

// Каждый обработчик содержит ссылку на следующий обработчик через поле next в BaseHandler.
// Тип поля — интерфейс Handler, поэтому в цепочку можно включать любые структуры, реализующие этот интерфейс.
// Таким образом формируется динамическая цепочка обработки запроса.

// Получается такая схема зависимости структур:
// AuthHandler
//    └── next → RoleHandler

// RoleHandler
//    └── next → RateLimitHandler

// RateLimitHandler
//    └── next → nil

package main

import "fmt"

// Структура обработчика
type Handler interface {
	SetNext(Handler)
	Handle(request *Request)
}

// Структура запроса
type Request struct {
	User      string
	Token     string
	UserRole  string
	RequestID int
}

// Базовый обработчик
type BaseHandler struct {
	next Handler
}

func (bh *BaseHandler) SetNext(handler Handler) {
	if bh.next == nil {
		bh.next = handler
	}
}

func (bh *BaseHandler) callNext(request *Request) {
	if bh.next != nil {
		bh.next.Handle(request)
	}
}

// Handler 1 — Проверка авторизации
type AuthHandler struct {
	BaseHandler
}

func (ah *AuthHandler) Handle(req *Request) {
	if req.Token == "" {
		fmt.Println("Auth: пользователь не авторизован")
		return
	}

	fmt.Println("Auth: пользователь авторизован")
	ah.callNext(req)
}

// Handler 2 — Проверка роли
type RoleHandler struct {
	BaseHandler
}

func (rh *RoleHandler) Handle(req *Request) {
	if req.UserRole != "admin" {
		fmt.Println("Role: Доступ запрещен")
		return
	}

	fmt.Println("Role: роль подтверждена")
	rh.callNext(req)
}

// Handler 3 — Проверка лимита
type RateHandler struct {
	BaseHandler
}

func (rh *RateHandler) Handle(req *Request) {
	if req.RequestID > 5 {
		fmt.Println("Rate: превышен лимит запросов")
		return
	}

	fmt.Println("Rate: лимит в норме")
	rh.callNext(req)
}

// Основная фукнция
// Запускае классический пример Chain of Responsibility
func RunClassicExample() {
	fmt.Println("Запущен классический пример Chain of Responsibility")

	auth := &AuthHandler{}
	role := &RoleHandler{}
	rate := &RateHandler{}

	// определяем цепочку последовательности
	auth.SetNext(role)
	role.SetNext(rate)

	// тестовый запрос
	newRequest := &Request{
		User:      "Kirill",
		Token:     "12345.qwerty.asdf123",
		UserRole:  "admin",
		RequestID: 2,
	}

	auth.Handle(newRequest)

}
