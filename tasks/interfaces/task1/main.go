/*
Задача:
Реализуй интерфейс Profitable и структуру StatisticProfit, которая будет создавать этот интерфейс.

Интерфейс Profitable содержит следующие методы:

SetProduct(p *Product) —  устанавливает продукт для статистики прибыли.
GetAverageProfit() float64 —  возвращает среднюю прибыль.
GetAverageProfitPercent() float64 —  должен возвращать не средний процент от прибыли, а процент средней прибыли, таким образом просто считаем от средней прибыли и средней стоимости закупки (не нужно считать проценты по каждой сделке и потом искать среднее значение).
GetCurrentProfit() float64 —  возвращает текущую прибыль.
GetDifferenceProfit() float64 -—  возвращает разницу между текущей ценой продукта и средней ценой продажи.
GetAllData() []float64 —  возвращает все данные о прибыли в виде среза чисел.
Average(prices []float64) float64 —  вычисляет среднее значение из среза чисел.
Sum(prices []float64) float64 —  вычисляет сумму чисел в срезе.
Структура StatisticProfit содержит следующие поля:

product *Product —  продукт для статистики прибыли.
getAverageProfit func() float64 —  функция для вычисления средней прибыли.
getAverageProfitPercent func() float64 —  функция для вычисления среднего процента прибыли.
getCurrentProfit func() float64 —  функция для вычисления текущей прибыли.
getDifferenceProfit func() float64 —  функция для вычисления разницы между текущей ценой продукта и средней ценой продажи.
getAllData func() []float64 —  функция для получения всех данных о прибыли в виде среза чисел.
Реализуй методы интерфейса Profitable в структуре StatisticProfit и добавь необходимые функции для вычисления указанных значений.

Критерии приемки:
Реализован интерфейс Profitable.
Реализована структура StatisticProfit.
Методы интерфейса Profitable возвращают корректные значения.
Все функции в структуре StatisticProfit работают правильно.
Все данные о прибыли возвращаются в виде среза чисел.
Все необходимые значения для статистики прибыли вычисляются правильно.
Ресиверы методов интерфейса Profitable имеют тип *StatisticProfit.

Исходный код:
package main

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable

// WithAverageProfit sets the average profit of the product
//
// Average Profit = Average Sells - Average Buys
func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = // ваш код здесь
}

// WithAverageProfitPercent sets the average profit percent of the product
//
// Average Profit Percent = Average Profit / Average Buys * 100
func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = // ваш код здесь
}

// WithCurrentProfit sets the current profit of the product
//
// Current Profit = Current Price - Current Price * (100 - Profit Percent) / 100
func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = // ваш код здесь
}

// WithDifferenceProfit sets the difference profit of the product
//
// Difference Profit = Current Price - Average Sells
func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = // ваш код здесь
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfit != nil {
			res = append(res, s.getAverageProfit())
		}
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite

	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)
}

*/

// Решение:

package main

import "fmt"

// "fmt"

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	//GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	//GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func NewStatisticProfit(product *Product, opts ...func(*StatisticProfit)) Profitable {
	s := &StatisticProfit{product: product}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithAverageProfit sets the average profit of the product
//
// Average Profit = Average Sells - Average Buys
func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
	//getAverage(s.product.Sells) - getAverage(s.product.Buys)
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	return s.Sum(prices) / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	var sum float64 = 0
	for _, x := range prices {
		sum += x
	}
	return sum
}

func (s *StatisticProfit) GetAverageProfit() float64 {
	return s.getAverageProfit()
}

func (s *StatisticProfit) GetCurrentProfit() float64 {
	return s.getCurrentProfit()
}

func (s *StatisticProfit) GetDifferenceProfit() float64 {
	return s.getDifferenceProfit()
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

// WithAverageProfitPercent sets the average profit percent of the product
//
// Average Profit Percent = Average Profit / Average Buys * 100
func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		return s.getAverageProfit() / s.Average(s.product.Buys) * 100
	}
	// s.getAverageProfit / getAverage(s.product.Buys) * 100
}

// WithCurrentProfit sets the current profit of the product
//
// Current Profit = Current Price - Current Price * (100 - Profit Percent) / 100
func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
	}
	// s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
}

// WithDifferenceProfit sets the difference profit of the product
//
// Difference Profit = Current Price - Average Sells
func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
	// s.product.CurrentPrice - getAverage(s.product.Sells)
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfit != nil {
			res = append(res, s.getAverageProfit())
		}
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite

	statProfit := NewStatisticProfit(
		product,
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)

	fmt.Println(product)
	fmt.Println(*statProfit)

	fmt.Println("Average Profit:", statProfit.GetAverageProfit())
	fmt.Println("Current Profit:", statProfit.GetCurrentProfit())
	fmt.Println("Difference Profit:", statProfit.GetDifferenceProfit())
}
