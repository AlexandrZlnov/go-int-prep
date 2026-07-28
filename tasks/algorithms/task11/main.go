// Задача
// Дана строка, содержащая только символы скобок.
// Проверить, правильно ли они закрыты и вложены.

// Решение:
// Идея:
// Идем range по строке
// Заполняем стэк открывающими скобками
// Как только получаем закрывающую, проверяем ее с послденей открывающей в стэке
// Если совпадаю, удаляем последнюю открывающую из стыка
// Если не совпадают - проверка не пройдена
// Вконце стэк открывающий скобок должен быть пустым

package main

import "fmt"

func main() {
	test := []string{
		"()",
		"()[]{}",
		"{[]}",
		"(]",
		"([)]",
		"{[]}",
		"()[]{([][])}",
	}

	for _, str := range test {
		fmt.Println("Проверяем -", str)
		result := CheckBrackets(str)

		switch result {
		case true:
			fmt.Printf("Вложенность скобок - правильная\n")
		case false:
			fmt.Printf("Вложенность скобок - не верна.\n")
		}

	}
}

func CheckBrackets(s string) bool {
	str := []rune(s)

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0, len(str))

	for _, ch := range str {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else if _, ok := pairs[ch]; ok {

			if len(stack) == 0 {
				fmt.Println("Закрывающая скобка не может быть первой - ", ch)
				return false
			}

			if pairs[ch] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}
