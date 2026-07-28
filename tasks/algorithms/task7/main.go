// Задача
// Дана строка. Нужно изменить порядок символов на противоположный:
// - первый символ становится последним,
// - второй — предпоследним и так далее.

// Решение:
// Временная сложность	O(n)
// Память	O(k)
// Принцип:
// Используются два указателя: символы с начала и конца строки попарно меняются местами

package main

import "fmt"

func main() {
	testStr := "Шла Маша по шоссе и сосала сушку!"

	fmt.Printf("Реверс строки %s\n--> %s\n", testStr, Reverse(testStr))
}

func Reverse(str string) string {
	strRune := []rune(str)
	left := 0
	right := len(strRune) - 1

	for left < right {
		strRune[left], strRune[right] = strRune[right], strRune[left]
		left++
		right--
	}

	return string(strRune)
}
