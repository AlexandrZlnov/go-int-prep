// Задача:
// Напишите функцию, которая принимает строку и возвращает первый символ,
// который встречается в строке только один раз. Если все символы повторяются,
// вернуть пустую строку

// Решение:
// Принцип: через мапу посчитали количесвто каждого символа
// Втормы циклом по строке нашли первый счет которого = 1

package main

import "fmt"

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
)

func main() {
	words := []string{"leetcode", "aabbcc", "qwerty", "qqwweertty"}
	for _, w := range words {
		fmt.Printf("В слове %s%q%s первый уникальный символ - %s%q%s\n", ColorRed, w, ColorReset, ColorRed, IsFirstUnique(w), ColorReset)
	}
}

func IsFirstUnique(w string) string {
	count := make(map[rune]int)

	for _, letter := range w {
		count[letter]++
	}

	for _, letter := range w {
		if count[letter] == 1 {
			return string(letter)
		}
	}
	return ""
}
