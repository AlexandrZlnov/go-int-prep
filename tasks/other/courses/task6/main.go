/*
Задание
Разработай программу для планирования интерфейса на языке программирования Golang.
Программа должна генерировать SQL-запросы для создания таблицы и вставки данных в базу данных SQLite.

package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"reflect"
	"strings"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	//
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	//
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	//
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}

Критерии завершенности:
- Создай интерфейс Tabler для получения имени таблицы.
- Интерфейс Tabler содержит метод TableName, который возвращает имя таблицы.
- Создан интерфейс SQLGenerator для генерации SQL-запросов и интерфейс для генерации фейковых данных.
- Интерфейс SQLGenerator содержит метод CreateTableSQL, который принимает аргумент table с типом Tabler и возвращает строку с SQL-запросом для создания таблицы.
- Интерфейс SQLGenerator содержит метод CreateInsertSQL, который принимает аргумент model с типом Tabler и возвращает строку с SQL-запросом для вставки данных в таблицу.
- Создан интерфейс FakeDataGenerator для генерации фейковых данных.
- Интерфейс FakeDataGenerator содержит метод GenerateFakeUser, который возвращает значение типа User с фейковыми данными.
*/

package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"reflect"
	"strings"
)

const (
	userTable = "users"
)

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"INTEGER PRIMARY KEY AUTOINCREMENT"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type SQLiteGenerator struct{}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

type Tabler interface {
	TableName() string
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

func (u *User) TableName() string {
	return userTable
}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	t := reflect.TypeOf(table).Elem()
	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbField := field.Tag.Get("db_field")
		dbType := field.Tag.Get("db_type")

		if dbField != "" && dbType != "" {
			fields = append(fields, dbField+" "+dbType)
		}
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXIST %s (%s);",
		table.TableName(),
		strings.Join(fields, ", "),
	)
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	v := reflect.ValueOf(model).Elem()
	t := v.Type()

	var fields []string
	var values []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbField := field.Tag.Get("db_field")

		if dbField == "id" {
			continue
		}

		fields = append(fields, dbField)
		value := v.Field(i).Interface()
		switch v.Field(i).Kind() {
		case reflect.String:
			values = append(values, fmt.Sprintf("'%v'", value))
		default:
			values = append(values, fmt.Sprintf("%v", value))

		}

	}
	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		model.TableName(),
		strings.Join(fields, ", "),
		strings.Join(values, ", "),
	)

}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}

	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
