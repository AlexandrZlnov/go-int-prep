// Собес: Мосдеп
// Задача:
// Что выведет код?

// Решение:
// Будет напечатан сериализованный json в виде строки. Тоесть поле значение.
// Строка будет содержать только экспртируемые поля структуры.
// Имена полей в соответствии с тэгами если тега нет по названию поля струкруты.
// Не заданные поля получат null или не будет сериалозовано вообзе в случае omitempty

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name     string   // не сериализуется, не экспортируемое
	Surname  string   // сериализуется по имени поля, нет тэга
	Address  string   `json:"address"`         // сериализуется с именем поля по тэгу, с мальнокой буквы
	Phone    string   `json:"-"`               // поля с тэгом "-" не подлежать сериализации или десериализации, не попадет в json
	Parents  []string `json:"parents"`         // значение не задано, сериализуется как null, имя поля по тэгу
	Children []string `json:"children"`        // в json попадет пустой слайс
	Extra    string   `json:"extra,omitempty"` // значение не задано, по причине omitempty не попадет в json
}

func main() {
	p := Person{
		name:     "John",
		Surname:  "Doe",
		Address:  "Moscow",
		Phone:    "+712345",
		Children: make([]string, 0),
	}

	res, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}
