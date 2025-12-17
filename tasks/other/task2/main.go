// из списка задач к собеседования - Avtostopom-po-Go
//Исправить функцию, чтобы она работала.
//Сигнатуру менять нельзя

package main

import "fmt"

// Задание
// func printNumber(ptrToNumber interface{}) {
// 	if ptrToNumber != nil {
// 		fmt.Println(*ptrToNumber.(*int))
// 	} else {
// 		fmt.Println("nil")
// 	}
// }

// Решение ↓
func printNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil {
		fmt.Println("nil")
		return
	}

	val, ok := ptrToNumber.(*int)
	if !ok || val == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(*ptrToNumber.(*int))
	}

	return
}

// Решение ↑

func main() {

	v := 10
	printNumber(&v)

	var pv *int

	printNumber(pv)
	pv = &v
	printNumber(pv)
}
