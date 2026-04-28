package main

import (
	"testing"
	"time"
)

func TestLimiter_Basic(t *testing.T) {
	lim := NewLimiter(2)

	if !lim.Allow() {
		t.Fatal("ожидаемый первый запрос, который будет передан")
	}
	if !lim.Allow() {
		t.Fatal("ожидаемый второ запрос, который будет передан")
	}
	if lim.Allow() {
		t.Fatal("ожидается, что третий запрос будет отклонен")

	}

}

// Проверка восстановления токенов
func TestLimiter_Refill(t *testing.T) {
	lim := NewLimiter(2)

	t.Log("initial tokens:", lim.tokens)

	lim.Allow()
	lim.Allow()

	t.Log("after consuming:", lim.tokens)

	if lim.Allow() {
		t.Fatal("expected limit exceeded")
	}

	time.Sleep(time.Second)

	t.Log("before refill call:", lim.tokens)

	if !lim.Allow() {
		t.Fatal("expected token refill after 1 second")
	}

	t.Log("after refill:", lim.tokens)
}
