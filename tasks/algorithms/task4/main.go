// Задача:
// Проверить, является ли строка палиндромом.
// Палиндром — это слово, фраза или число, которые
// одинаково читаются в обе стороны: слева направо и справа налево.

// Решение:
// Временная сложность	O(n)
// Память	O(n)
package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{
		"level",
		"казаК",
		"шалаши",
		"1236321",
	}

	for _, word := range words {
		if !IsPalindrome(word) {
			fmt.Printf("Строка %s - не палиндром\n", word)
			continue
		}
		fmt.Printf("Строка %s - палиндром\n", word)
	}
}

func IsPalindrome(word string) bool {
	wordRune := []rune(strings.ToLower(word))
	leftInd := 0
	rightInd := len(wordRune) - 1

	for leftInd < rightInd {
		if wordRune[leftInd] != wordRune[rightInd] {
			return false
		}

		leftInd++
		rightInd--

	}

	return true
}
