// Задача: rate limiter (token bucket упрощённый)
// Условие:
// - Нужно реализовать limiter
// Требования:
// - не более rps операций в секунду
// - безопасно для concurrent использования
// - без внешних библиотек

// Исходный код:
/*
type Limiter struct {
    // TODO
}

func NewLimiter(rps int) *Limiter {
    // rps = сколько запросов в секунду можно
}

func (l *Limiter) Allow() bool {
    // true если можно выполнить запрос
    // false если лимит превышен
}
*/

// Решение:

package main

import (
	"sync"
	"time"
)

type Limiter struct {
	// TODO
	rate       int
	tokens     float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewLimiter(rps int) *Limiter {
	// rps = сколько запросов в секунду можно
	return &Limiter{
		rate:       rps,
		tokens:     float64(rps),
		lastRefill: time.Now(),
		mu:         sync.Mutex{},
	}
}

func (l *Limiter) Allow() bool {
	// true если можно выполнить запрос
	// false если лимит превышен
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(l.lastRefill)

	newTokens := elapsed.Seconds() * float64(l.rate)

	l.tokens = min(l.tokens+newTokens, float64(l.rate))
	l.lastRefill = now

	if l.tokens >= 1 {
		l.tokens -= 1
		return true
	}

	return false
}
