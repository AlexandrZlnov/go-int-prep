// Builder Pattern

package main

import "fmt"

//Outfit — это конечный продукт, который мы собираем.
//Это структура, которая хранит итоговое состояние объекта
//Метод Show() — это просто способ показать результат.
//Outfit ничего не знает о процессе сборки.
//Он просто хранит данные.

type Outfit struct {
	Shoes      string
	Jacket     string
	Underewear string
}

func (o Outfit) Show() {
	fmt.Printf("Outfit:\nShoes: %s\nJacket: %s\nUnderwear: %s\n\n", o.Shoes, o.Jacket, o.Underewear)
}

type OutfitBuilder interface {
	SetShoes()
	SetJacket()
	SetUnderwear()
	GetOutfit() Outfit
}

//SportOutfitBuilder — это конкретный строитель, который:
//Хранит внутри объект Outfit (композиция).
//Пошагово заполняет его поля
//Возвращает готовый объект через GetOutfit().
//Его задача: знать, как собрать спортивный комплект.

type SportOutfitBuilder struct {
	outfit Outfit
}

func (sob *SportOutfitBuilder) SetShoes() {
	sob.outfit.Shoes = "Sport Shoes"
}

func (sob *SportOutfitBuilder) SetJacket() {
	sob.outfit.Jacket = "Sport Jacket"
}

func (sob *SportOutfitBuilder) SetUnderwear() {
	sob.outfit.Underewear = "Sport Underwear"
}

func (sob *SportOutfitBuilder) GetOutfit() Outfit {
	return sob.outfit
}

// Director — это управляющий сборкой объект, который:
// - хранит ссылку на интерфейс строителя (OutfitBuilder);
// - через метод SetBuilder() получает конкретную реализацию строителя;
// не знает, какой именно это строитель: спортивный классический зимний и тд
// просто вызывает шаги сборки.
// То есть:
// Director хранит не SportOutfitBuilder,
// а интерфейс OutfitBuilder.
// И через этот интерфейс вызывает методы.

type Director struct {
	builder OutfitBuilder
}

func (d *Director) SetBuilder(b OutfitBuilder) {
	d.builder = b
}

// Build — метод, который управляет процессом сборки.
// То есть Build не создаёт структуру напрямую,
// а заставляет строителя собрать её шаг за шагом.

func (d *Director) Build() Outfit {
	d.builder.SetShoes()
	d.builder.SetJacket()
	d.builder.SetUnderwear()

	return d.builder.GetOutfit()
}

func main() {
	sportBuilder := &SportOutfitBuilder{}
	director := Director{}
	director.SetBuilder(sportBuilder)

	sportOutfit := director.Build()
	sportOutfit.Show()
}
