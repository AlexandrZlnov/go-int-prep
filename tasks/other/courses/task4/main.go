/*
Задание
Управление счетом покупателей с функциональными опциями
Создай конструктор для объекта Customer, который позволяет задавать различные опции при создании экземпляра. В этом случае опции позволяют задавать имя и тип счета для клиента.

package main

import (
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	err := customer.Account.Withdraw(100)
	if err != nil {
        fmt.Println(err)
    }

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}

Критерии завершенности:
- Конструктор NewCustomer создает экземпляр Customer с заданными опциями.
- Опция WithName устанавливать имя клиента.
- Опция WithAccount устанавливает тип счета клиента.
- Метод Withdraw корректно снимает деньги со счета клиента, если на счете достаточно средств.
- SavingsAccount не может совершать операции снятия при балансе меньше 1000.
- CheckingAccount не может совершать операции снятия при балансе меньше запрашиваемой суммы.
- Обе реализации конкурентно-безопасны.
- При создании клиента с опциями имя и баланс счета установлены корректно.
- При выполнении кода в функции main выводится имя клиента и баланс его счета.
*/

package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type Customer struct {
	ID      int
	Name    string
	Account Account
}

type SavingsAccount struct {
	mu      sync.Mutex
	balance float64
}

func (s *SavingsAccount) Deposit(am float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.balance += am
}

func (s *SavingsAccount) Withdraw(am float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.balance-am >= 1000 {

		s.balance -= am
		return nil
	}
	return errors.New("не достаточно средств на счету")
}

func (s *SavingsAccount) Balance() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.balance
}

type CheckingAccount struct {
	mu      sync.Mutex
	balance float64
}

func (c *CheckingAccount) Deposit(am float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance += am
}

func (c *CheckingAccount) Withdraw(am float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.balance >= am {
		c.balance -= am
		return nil
	}
	return errors.New("не достаточно средств на счету")
}

func (c *CheckingAccount) Balance() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.balance
}

type Option func(*Customer)

func WithName(val string) Option {
	return func(cust *Customer) {
		cust.Name = val
	}
}

func WithAccount(acc Account) Option {
	return func(cust *Customer) {
		cust.Account = acc
	}
}

func NewCustomer(id int, options ...Option) *Customer {
	cust := &Customer{
		ID: id,
	}

	for _, opt := range options {
		opt(cust)
	}
	return cust
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1200)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	err := customer.Account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
