// Factory-Method Pfttern
//
// Суть:
// Не создавать объект напрямую через new или литерал типа,
// а делегировать создание отдельному методу/фабрике
//
// Фабрика может:
// - создавать разные реализации интерфейса
// - инкапсулировать сложную логику создания
// - упростить расширение кода без изменения клиентского кода

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

type ClothingFactory interface {
	CreateClothing() Clothing
}

type ShoesFactory struct{}

func (clF ShoesFactory) CreateClothing() Clothing {
	return Shoes{
		brend:  "Lacoste",
		season: "summer",
		size:   50,
	}
}

type JacketFactory struct{}

func (clF JacketFactory) CreateClothing() Clothing {
	return Jacket{
		brend:  "DG",
		season: "summer",
		size:   48,
	}
}

type UnderwearFactory struct{}

func (clF UnderwearFactory) CreateClothing() Clothing {
	return Underwear{
		brend: "Calvin Klein",
		size:  "XS",
	}
}

func main() {
	var factory ClothingFactory

	factory = ShoesFactory{}
	item1 := factory.CreateClothing()
	fmt.Print(item1.GetInfo())

	factory = JacketFactory{}
	item2 := factory.CreateClothing()
	fmt.Print(item2.GetInfo())

	factory = UnderwearFactory{}
	item3 := factory.CreateClothing()
	fmt.Print(item3.GetInfo())

}

/*
          interface Clothing
        +------------------+
        | GetInfo() string |
        +------------------+
                 ^
                 |
        +--------+--------+
        |                 |
      struct Shoes        struct Jacket
      +----------------+ +----------------+
      | GetInfo()      | | GetInfo()      |
      +----------------+ +----------------+

          interface ClothingFactory
        +---------------------+
        | Create() Clothing   |
        +---------------------+
                 ^
                 |
      +----------+-----------+
      |                      |
  struct ShoesFactory      struct JacketFactory
  +------------------+    +------------------+
  | Create() Shoes   |    | Create() Jacket  |
  +------------------+    +------------------+

CLIENT (main)
-----------------
factory := ShoesFactory{}
shoes := factory.Create()
shoes.GetInfo()
*/
