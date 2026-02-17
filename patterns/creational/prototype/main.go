package main

import (
	"fmt"
)

// Prototype интерфейс с методом Clon
type Prototype interface {
	Clone() Prototype
	Show()
}

// Concrete Prototype — персонаж
type Character struct {
	Name   string
	Health int
	Power  int
	Skills []string
}

// Реализация Clone (поверхностная копия)
// Или глубокая
func (c *Character) Clone() Prototype {
	//copy := *c
	//return &copy

	copyChar := *c

	copyChar.Skills = make([]string, len(c.Skills))
	copy(copyChar.Skills, c.Skills)

	return &copyChar
}

// Метод для отображения состояния персонажа
func (c *Character) Show() {
	fmt.Printf("Character:\nName: %s\nHealth: %d\nPower: %d\nSkills: %v\n\n", c.Name, c.Health, c.Power, c.Skills)
}

func main() {
	// Оригинал
	original := &Character{
		Name:   "Gendalf",
		Health: 150,
		Power:  200,
		Skills: []string{"Magic", "Sword"},
	}

	// Клонируем персонажа
	clone1 := original.Clone().(*Character)
	clone2 := original.Clone().(*Character)

	clone1.Name = "Bilbo"
	clone1.Skills = append(clone1.Skills[1:], "Small")

	clone2.Name = "Legolas"
	clone2.Skills = append(clone2.Skills, "Bow")

	original.Show()
	clone1.Show()
	clone2.Show()

}
