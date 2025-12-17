package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	//fmt.Println("Person In function:", &person.Name)
	person = &Person{
		Name: "Alice",
	}
	//
	//fmt.Println("In function:", &person.Name)

}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	//fmt.Println("Our of function:", &person.Name)
	changeName(person)
	fmt.Println(person.Name)
}
