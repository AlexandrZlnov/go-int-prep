// из списка задач к собеседования - Avtostopom-po-Go
//Что выведет? Как исправить?

package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person **Person) {
	*person = &Person{
		Name: "Alice",
	}
}

func main() {

	person := &Person{ // person это указатель на структуру Person      ///
		Name: "Bob",
	}

	fmt.Println(person.Name) // Выведет "Bob"
	changeName(&person)
	fmt.Println(person.Name) // Выведет "Bob"
}
