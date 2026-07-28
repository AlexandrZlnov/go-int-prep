// Задача
// Скомпилируется или нет?
// Если да, то что выведет?

// Ответ:
// Скомпилируется
// f просвоим значение 1
// вызываем метод и передаем аргуметном foo3 = 4 или 100
// поскольку в методе flag передается по значение а не как указатель изменения в f за пределами методу не будут видны.
// следовательно вывод дасть первоначальное значение f в двоичном виде по глаголу %b - 1

package main

import "fmt"

var (
	foo1 flag = 1
	foo2 flag = 1 << 1
	foo3 flag = 1 << 2
)

type flag uint8

func (f flag) set(ff flag) {
	f |= ff
}

func main() {
	f := foo1
	f.set(foo3)
	fmt.Printf("%b\n", f)
}

type customer struct {
	balance float64
}

func (c *customer) add(v float64) {
	c.balance += v
}

func add(c *customer, v float64) {
	c.balance += v
}
