// Задача:
// Что выведет код?
package main

import "fmt"

func main() {
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f()
	}
}

// Ответ:
// 0 1 2
// в версии Go 1.22 и младще вывод будет 3 3 3
