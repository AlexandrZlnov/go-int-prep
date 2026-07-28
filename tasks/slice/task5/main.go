// Задача
// Отсортируете массив структур по алфавиту по полю Name?

// Решение

package main

import (
	"fmt"
	"sort"
)

// Структура Person
type Person struct {
	Name string
	Age  int
}

func main() {
	// Исходный массив (срез) структур
	people := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 20},
		{"alex", 35}, // обратите внимание: с маленькой буквы
		{"Анна", 28}, // кириллица
	}

	// Сортировка по полю Name (по алфавиту)
	// задаем условие отвечая на вопрос:
	// "элемент с индексом i должен быть перед элементом с индексом j,
	// если значение его поля Name лексикографически меньше".
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	// Вывод результата
	fmt.Println("Отсортировано по имени:")
	for _, p := range people {
		fmt.Printf("%s - %d лет\n", p.Name, p.Age)
	}
}
