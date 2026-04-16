// Паттерн Flyweight (Легковес) или (Приспособленец)
// Flyweight — это разделение общего и уникального состояния для экономии памяти.
// Позволяет разделять общие неизменяемые данные между множеством объектов через фабрику,
// чтобы уменьшить потребление памяти.
// В данном примере Flyweight = кэш общих тяжёлых объектов + ссылка на них из лёгких объектов
// Country — тяжёлый общий объект (intrinsic state).
// User — лёгкий объект с уникальными данными (extrinsic state).
// CountryFactory — контролирует создание и повторное использование.
// map[string]*Country — кэш.
package main

import "fmt"

// TaxRules и Permission - Лёгкие вспомогательные структуры.
// Часть "тяжёлого" состояния Country (intrinsic state),
// нужны чтобы показать, что Country — большой (тяжелый)объект.
type TaxRules struct {
	Name   string
	Amount float64
}

type Permission struct {
	Name string
}

// Flyweight-объект.
// Это тяжёлый, общий для многих пользователей объект (intrinsic state).
// Он НЕ должен создаваться заново для каждого User.
type Country struct {
	Name       string
	Population int64
	Currency   string
	Languages  []string
	TaxRules   []TaxRules
	Permitions []Permission
}

// Фабрика Flyweight.
// Хранит уже созданные объекты Country в мапе
// и гарантирует, что для одного имени (страны) будет только один экземпляр.
type CountryFactory struct {
	countries map[string]*Country
}

func NewCountryFactory() *CountryFactory {
	return &CountryFactory{
		countries: make(map[string]*Country),
	}
}

// Ключевой метод паттерна Flyweight.
// 1) Проверяет, создан ли уже объект.
// 2) Если да — возвращает существующий.
// 3) Если нет — создаёт, сохраняет в map и возвращает.
// Таким образом обеспечивается повторное использование объектов.
func (cf *CountryFactory) GetCountry(name string) *Country {
	if c, ok := cf.countries[name]; ok {
		fmt.Printf("Страна %s уже есть в списке\n", name)
		return c
	}

	var c *Country

	switch name {
	case "USA":
		c = &Country{
			Name:       "USA",
			Population: 340000000,
			Currency:   "USD",
			Languages:  []string{"American English", "Spanish"},
			TaxRules: []TaxRules{
				{"Income_Tax", 0.2},
				{"VAT", 0.1},
			},
			Permitions: []Permission{
				{"CanVote"},
				{"CanWork"},
			},
		}
	case "Canada":
		c = &Country{
			Name:       "Canada",
			Population: 40385000,
			Currency:   "USD",
			Languages:  []string{"English", "French"},
			TaxRules: []TaxRules{
				{"Income_Tax", 0.4},
				{"VAT", 0.3},
			},
			Permitions: []Permission{
				{"CanBayRealEstate"},
				{"CanWork"},
			},
		}

	}
	cf.countries[name] = c
	return c
}

// Лёгкий объект (контекст).
// Хранит своё уникальное состояние (ID, Name)
// и ссылку на общий Flyweight-объект Country.
type User struct {
	ID      int
	Name    string
	Country *Country
}

// Демонстрация работы Flyweight:
// Несколько пользователей получают один и тот же объект Country.
// Проверка через == показывает, что это один и тот же указатель в памяти.
// Экономия памяти достигается за счёт совместного использования Country.
func main() {
	factory := NewCountryFactory()

	users := []User{
		{ID: 1, Name: "Viktor", Country: factory.GetCountry("USA")},
		{ID: 2, Name: "Boris", Country: factory.GetCountry("USA")},
		{ID: 3, Name: "Lesha", Country: factory.GetCountry("Canada")},
		{ID: 4, Name: "Sergey", Country: factory.GetCountry("USA")},
		{ID: 5, Name: "Kira", Country: factory.GetCountry("Canada")},
	}

	fmt.Println("USA для users[0] == users[1]:", users[0].Country == users[1].Country)            // true
	fmt.Println("Canada для users[2] == users[4]:", users[2].Country == users[4].Country)         // true
	fmt.Println("USA для users[0] == Canada для users[2]:", users[0].Country == users[2].Country) // false

	// Выводим данные пользователей
	for _, u := range users {
		fmt.Printf("User %s, Country: %s, Population: %d\n", u.Name, u.Country.Name, u.Country.Population)
	}
}

// Вывод:
/*
Страна USA уже есть в списке
Страна USA уже есть в списке
Страна Canada уже есть в списке
USA для users[0] == users[1]: true
Canada для users[2] == users[4]: true
USA для users[0] == Canada для users[2]: false
User Viktor, Country: USA, Population: 340000000
User Boris, Country: USA, Population: 340000000
User Lesha, Country: Canada, Population: 40385000
User Sergey, Country: USA, Population: 340000000
User Kira, Country: Canada, Population: 40385000
*/
