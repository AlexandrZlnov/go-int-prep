package main

import (
	"fmt"
	"log"
	"time"

	//"log"
	"net/http"
	//"time"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("Logging: старт", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
		fmt.Println("Logging: продолжительность", time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Auth: Пользователь не аторизован", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("limit") == "true" {
				http.Error(w, "Превышено количество запросов", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
}

func FinalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Final: Привет! Запрос прошел все middlewares!")
}

func RunMiddlewareExample() {
	handler := Chain(
		http.HandlerFunc(FinalHandler),
		LoggingMiddleware,
		AuthMiddleware,
		RateLimitMiddleware,
	)

	http.Handle("/hello", handler)

	fmt.Println("Запускаем сервер на порту 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
