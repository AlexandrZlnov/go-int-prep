// Паттерн Facade (Фасад)
// Фасад — это объект, который: даёт простой интерфейс к сложной подсистеме.
// Фасад, упрощает использование нескольких компонентов.

// есть процесс регистрации, который состоит из нескольких шагов: Валидация, Хэширование, Сохранение, Отправка письма
// Каждый шаг — отдельная подсистема.
// Фасад объединяет их в один простой метод Register.

package main

import "fmt"

// Валидатор
type Validator struct{}

func (v *Validator) ValidateEmail(email string) bool {
	fmt.Println("Валидация пользователя:", email)
	return len(email) > 5
}

// Хэширование пароля
type PaswordHasher struct{}

func (p *PaswordHasher) Hash(password string) string {
	fmt.Println("Хэширование пароля")
	return "Хэш_пароля_" + password
}

// Репозиторий
type UserRepository struct{}

func (*UserRepository) Save(email, hashedPassword string) {
	fmt.Printf("Сохранение в БД пользователя: %s, %s\n", email, hashedPassword)
}

// Сервис информирования
type EmailService struct{}

func (e *EmailService) SendWelcome(email string) {
	fmt.Printf("Отправка приветственного сообщения: %s\n", email)
}

// Фасад
type UserRegistrationFacade struct {
	validator *Validator
	hasher    *PaswordHasher
	repo      *UserRepository
	mailer    *EmailService
}

func NewValidateUserService() *UserRegistrationFacade {
	return &UserRegistrationFacade{
		validator: &Validator{},
		hasher:    &PaswordHasher{},
		repo:      &UserRepository{},
		mailer:    &EmailService{},
	}
}

func (n *UserRegistrationFacade) Register(email, password string) {
	fmt.Println("=== Регистрация пользователя ===")

	if !n.validator.ValidateEmail(email) {
		fmt.Println("Нет пользователя:", email)
		return
	}

	hpass := n.hasher.Hash(password)
	n.repo.Save(email, hpass)
	n.mailer.SendWelcome(email)

}

func main() {
	user := "ivan@gmail.ru"
	password := "Ivan_123"

	facade := NewValidateUserService()
	facade.Register(user, password)

}

// Вывод:
// === Регистрация пользователя ===
// Валидация пользователя: ivan@tmail.ru
// Хэширование пароля
// Сохранение в БД пользователя: ivan@tmail.ru, Хэш_пароля_Ivan_123
// Отправка приветственного сообщения: ivan@tmail.ru
