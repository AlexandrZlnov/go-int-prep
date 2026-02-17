// Пример реализации паттерна Adapter
// Adapter (Адаптер) — это структурный паттерн, который позволяет объектам
// с несовместимыми интерфейсами работать вместе. Он создаёт промежуточный
// объект-адаптер, который реализует ожидаемый интерфейс системы и внутри
// вызывает методы стороннего или устаревшего класса, преобразуя данные и
// сигнатуры вызовов. Обычно адаптер содержит ссылку на «несовместимый» объект
// и переводит один интерфейс в другой. Паттерн используется при интеграции
// сторонних библиотек, legacy-кода или разных API, чтобы изолировать
// бизнес-логику от внешних зависимостей.

package main

import (
	"fmt"
)

// Целевой интерфейс (Target)
// Вся бизнес-логика работает только через этот интерфейс
type PaymentProvider interface {
	Pay(amount float64) error
}

// Штатная (внутренняя) платежная система
type InternalPaymentSystem struct {
}

// Эта система уже реализует нужный интерфейс
func (ips InternalPaymentSystem) Pay(amount float64) error {
	fmt.Printf("====---InternalService---====\nPaid: %.2f USD\n", amount)
	return nil
}

// 1-я сторонняя система Stripe (несовместимый интерфейс)
// имитация создания нового платежа
// через внешнюю платежную систему
// которая не имеет метода Pay по умолчанию
type StripeClient struct {
}

func (sc StripeClient) CreateCharge(cents int64, currency string) error {
	fmt.Printf("====---Strite---====\nCharge: %d cents %s\n", cents, currency)
	return nil
}

// Adapter для Stripe
type StripeAdapter struct {
	client   *StripeClient
	currency string
}

func (sa *StripeAdapter) Pay(amount float64) error {
	cents := int64(amount * 100)
	sa.client.CreateCharge(cents, sa.currency)
	return nil
}

// 2-я сторонняя система PayPal (другой интерфейс)
type PayPalClient struct {
}

func (pp *PayPalClient) SendPayment(amount float64) error {
	fmt.Printf("====---PayPal---====\nPayment: %.2f USD\n", amount)
	return nil
}

// Adapter для PayPal
type PayPalAdapter struct {
	client *PayPalClient
}

func (ppa *PayPalAdapter) Pay(amount float64) error {
	ppa.client.SendPayment(amount)
	return nil
}

// Бизнес-логика (Client)
func MakePayment(provider PaymentProvider, amount float64) {
	err := provider.Pay(amount)
	if err != nil {
		fmt.Println("Ошибка платежа", err)
		return
	}
	fmt.Println("Платеж прошел удачно")
}

func main() {
	fmt.Println("\n-------InternalService-------")
	internal := &InternalPaymentSystem{}
	MakePayment(internal, 10.50)

	fmt.Println("\n-------StripeClient-------")
	stripe := &StripeClient{}
	stripeAdapter := &StripeAdapter{
		client:   stripe,
		currency: "USD",
	}
	MakePayment(stripeAdapter, 10.50)

	fmt.Println("\n-------PayPal-------")
	paypal := &PayPalClient{}
	paypalAdapter := &PayPalAdapter{
		client: paypal,
	}
	MakePayment(paypalAdapter, 10.50)

}
