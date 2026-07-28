// Задача:
// Разбери код, что получим на выводе?

package main

import "fmt"

type NamedParams struct {
	Params map[string]int
	Name   string
}

func (s NamedParams) Clear() {
	s.Name = "no name"
	for k := range s.Params {
		delete(s.Params, k)
	}
}

func main() {
	p := NamedParams{Name: "box", Params: map[string]int{"length": 100, "width": 200}}
	fmt.Println(p.Name, p.Params["length"], p.Params["other"]) // box 100 0 -> 0 выводиться по уколчаюнию как zero value в случае отсутствия ключа в мапе
	p.Clear()
	fmt.Println(p.Name, p.Params["length"], p.Params["other"]) // box 0 0 -> остается box поскольку ресивер в методе вызывается по значению
	// соответственно передается копия структуры => no name значение присвоено в копии структуры. 100 замениться на 0 потому что map это указатель на hmap
	// и он тоже передан как копия, но значени указателя отправляет нас к исходной мапе, поэтому значения в ней изменяться.
}
