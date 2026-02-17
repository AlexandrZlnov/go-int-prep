// Simple Factory-method
// Упрощенный вариант - нарушает принцип Open/Close из SOLID

package main

import (
	"fmt"
)

type Clothing interface {
	GetInfo() string
}

type Shoes struct {
	brend  string
	season string
	size   int
}

func (cl Shoes) GetInfo() string {
	return fmt.Sprintf("Shoes brend - %s, season - %s, size - %d\n", cl.brend, cl.season, cl.size)
}

type Jacket struct {
	brend  string
	season string
	size   int
}

func (cl Jacket) GetInfo() string {
	return fmt.Sprintf("Jacket brend - %s, season - %s, size - %d\n", cl.brend, cl.season, cl.size)
}

type Underwear struct {
	brend string
	size  string
}

func (cl Underwear) GetInfo() string {
	return fmt.Sprintf("Underwear brend - %s, size - %s\n", cl.brend, cl.size)
}

func GetClothing(clothing string) Clothing {
	switch clothing {
	case "shoes":
		item := Shoes{
			brend:  "Lacoste",
			season: "summer",
			size:   50,
		}
		return item
	case "jacket":
		item := Jacket{
			brend:  "DG",
			season: "summer",
			size:   48,
		}
		return item
	case "underwear":
		item := Underwear{
			brend: "Calvin Klein",
			size:  "XS",
		}
		return item
	default:
		return nil
	}
}

func main() {
	clothing := []string{"shoes", "jacket", "hat", "underwear"}

	for _, cl := range clothing {
		item := GetClothing(cl)
		if item == nil {
			fmt.Printf("Нет такого типа одежды: %s\n\n", cl)
			continue
		}

		fmt.Println(item.GetInfo())
	}
}
