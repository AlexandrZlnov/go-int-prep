// Вопросы:
// • Сколько байт занимает a на x64?
// • Сколько байт занимает s?
// • Сколько байт занимает i?
// • Почему размер i не равен размеру Foo?
// • Что реально хранится внутри interface{} в Go?

package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	A int64
}

type IFoo interface {
	Get() int64
}

func (f Foo) Get() int64 { return f.A }

func main() {
	var a int64 = 10
	var s Foo = Foo{A: 42}
	var i IFoo = s

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(i))
}

// Ответы.
// • Сколько байт занимает a на x64?
// ---- Ответ: 8 байт
// • Сколько байт занимает s?
// ---- Ответ: 8 байт (структура содержит одно поле int64)
// • Сколько байт занимает i?
// ---- Ответ: 16 байт
// • Почему размер i не равен размеру Foo?
// ---- Ответ: interface хранит не само значение, а: указатель на тип указатель на данные. Это всегда 2 слова по 8 байт = 16 байт на x64
// • Что реально хранится внутри interface{} в Go?
// ---- Ответ: Внутри: 1) указатель на тип (type / itab) 2) указатель на данные. interface = (type pointer, data pointer)
