// Задача
// Проверить, являются ли две строки анаграммами.
// Анаграмма — одинаковые символы в одинаковом количестве.

// Решение
// Сложность:
// O(n) по времени
// O(n) по памяти

package main

import (
	"fmt"
	"strings"
)

func main() {
	test := []struct{ w1, w2 string }{
		{"Апельсин", "апельсиН"},
		{"algorithm", "rithmalgo"},
		{"qwerty", "werty"},
		{"ac", "bb"},
	}

	for _, words := range test {
		if IsAnаgram(words.w1, words.w2) {
			fmt.Printf("Слова: %s и %s являются анограммами\n", words.w1, words.w2)
			continue
		}
		fmt.Printf("Слова: %s и %s не анограммы\n", words.w1, words.w2)
	}

}

func IsAnаgram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		fmt.Printf("Разная длина слов. ")
		return false
	}

	w1 = strings.ToLower(w1)
	w2 = strings.ToLower(w2)

	count := make(map[rune]int)

	for _, w := range w1 {
		count[w]++
	}

	for _, w := range w2 {
		count[w]--

		if count[w] < 0 {
			return false
		}
	}

	return true
}
