// из списка задач к собеседования - Avtostopom-po-Go
//Добавить код, который выведет тип переменной whoami

package main

import "fmt"

func printType(whoami interface{}) {
	fmt.Printf("Type is - %T\n", whoami)
	//fmt.Printf("Type of whoami: %v\n", reflect.TypeOf(whoami))
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
