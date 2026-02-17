// Abstract-Factory Pattern

package main

import (
	"fmt"
)

type Shoes interface { //	<---------------------------
	GetShoesInfo() string
}

type Jacket interface { //	<---------------------------
	GetJacketInfo() string
}

// Конкретные продукты: Sport
type SportShoes struct {
}

func (sh SportShoes) GetShoesInfo() string {
	return "Sport Shoes"
}

type SportJacket struct {
}

func (jc SportJacket) GetJacketInfo() string {
	return "Sport Jackets"
}

// Конкретные продукты: Classic
type ClassicShoes struct {
}

func (sh ClassicShoes) GetShoesInfo() string {
	return "Classic Shoes"
}

type ClassicJacket struct {
}

func (jc ClassicJacket) GetJacketInfo() string {
	return "Classic Jackets"
}

// Абстрактная фабрика
type ClothingFactory interface { //	<---------------------------
	CreateShoes() Shoes
	CreateJacket() Jacket
}

// Конкретные фабрики
type SportFactory struct{}

func (sf SportFactory) CreateShoes() Shoes {
	return SportShoes{}
}

func (sf SportFactory) CreateJacket() Jacket {
	return SportJacket{}
}

type ClassicFactory struct{}

func (sf ClassicFactory) CreateShoes() Shoes {
	return ClassicShoes{}
}

func (sf ClassicFactory) CreateJacket() Jacket {
	return ClassicJacket{}
}

// Клиентский код
func useFactory(factory ClothingFactory) {
	shoes := factory.CreateShoes()
	jacket := factory.CreateJacket()

	fmt.Println(shoes.GetShoesInfo())
	fmt.Println(jacket.GetJacketInfo())
}

func main() {
	var factory ClothingFactory

	// выбираем фабрику
	factory = SportFactory{}
	useFactory(factory)

	fmt.Println("-----")

	factory = ClassicFactory{}
	useFactory(factory)
}

/*
          interface Shoes
        +------------------+
        | GetShoesInfo()   |
        +------------------+
                 ^
                 |
            struct SportShoes
            struct ClassicShoes

          interface Jacket
        +------------------+
        | GetJacketInfo()  |
        +------------------+
                 ^
                 |
           struct SportJacket
           struct ClassicJacket

          interface ClothingFactory
        +---------------------------+
        | CreateShoes() Shoes       |
        | CreateJacket() Jacket     |
        +---------------------------+
                 ^
                 |
      +----------+----------+
      |                     |
   struct SportFactory    struct ClassicFactory
   +-----------------+   +------------------+
   | CreateShoes()   |   | CreateShoes()    |
   | CreateJacket()  |   | CreateJacket()   |
   +-----------------+   +------------------+

CLIENT (main)
-----------------
factory := SportFactory{}
shoes := factory.CreateShoes()    // SportShoes
jacket := factory.CreateJacket()  // SportJacket
*/
