// Идиоматичные конструкторы в Go

// Подход с функциональными опциями - functional options.
// Option — это функция, которая принимает *Object и что-то с ним делает
// WithField1 — это функция, которая создаёт и возвращает функцию
// То есть она не настраивает объект, она создаёт инструкцию, как его настраивать.
// WithFieldX — это фабрика функций-настроек

package main

type Object struct {
	field1 string
	field2 int
}

type Option func(*Object)

func WithField1(value string) Option {
	return func(obj *Object) {
		obj.field1 = value
	}

}

func WithField2(value int) Option {
	return func(obj *Object) {
		obj.field2 = value
	}
}

func NewObject(options ...Option) *Object {
	obj := &Object{
		field1: "Default",
		field2: 0,
	}

	for _, opt := range options {
		opt(obj)
	}
	return obj
}

func main() {
	NewObject(
		WithField1("Поле-1"),
		WithField2(1000),
	)
}
