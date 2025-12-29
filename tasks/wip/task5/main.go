// что выведет код?
package main

import (
	"log/slog"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("Panic occurred", "error", err)
		}
	}()
	go func() { panic("something bad happened") }()
}

//Ответ:
// Возможно 2 варианта:
// Наиболее вероятный - не выведен ничего. Panic останавливает выполнение только текущей горутины
// он никак не попадет в defer который в main, поэтому recover даст nil
// Маловероятный - если горутина отработает быстрее чем main, тогда получим панику
// panic: something bad happened
