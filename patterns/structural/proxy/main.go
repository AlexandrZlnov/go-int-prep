// Паттерн - Proxy (Прокси)
// Суть паттерна: создается объект-заместитель, который контролирует доступ к другому объекту.
// Proxy — это представитель объекта.
package main

import (
	"fmt"
	"time"
)

type DataService interface {
	GetData(id int) string
}

type RealDataService struct{}

func NewRealDataSevice() *RealDataService {
	fmt.Println("Создание подключения к БД")
	time.Sleep(1 * time.Second) //Имитация тяжелой инициализации
	return &RealDataService{}
}

func (d *RealDataService) GetData(id int) string {
	return fmt.Sprintf("Данные из БД для ID = %d", id)
}

// Proxy (владеет и создаёт real object)
type DataServiceProxy struct {
	real *RealDataService
}

func (p *DataServiceProxy) GetData(id int) string {
	if p.real == nil {
		fmt.Println("RealDataSevice еще не создан, создаем его сейчас--->")
		p.real = NewRealDataSevice()
	}

	// Контроль доступа (можно добавить auth, логирование и т.п.)
	fmt.Println("Proxy: обработка запроса")

	return p.real.GetData(id)
}

func NewDataServiceProxy() *DataServiceProxy {
	return &DataServiceProxy{}
}

func main() {
	// Используем интерфейс, чтобы клиент был независим от реализации и мог
	// использовать любую реализацию без изменения кода.
	var service DataService = NewDataServiceProxy()

	fmt.Println("Proxy создан. Реального объекта ещё нет.")
	fmt.Println("----")

	fmt.Println(service.GetData(1))
	fmt.Println("----")

	fmt.Println(service.GetData(2))
}
