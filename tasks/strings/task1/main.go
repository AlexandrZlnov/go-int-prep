// Package challenge6 contains the solution for Challenge 6.
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
//"The quick brown fox jumps over the lazy dog."
//"Hello, hello! How are you doing today? Today is a great day."

package main

import (
	"fmt"
	"strings"
)

func CountWordFrequency(text string) map[string]int {
	words := make(map[string]int)
	splitFunc := func(c rune) bool {
		return c == ' ' || c == '!' || c == '?' || c == ',' || c == '.' || c == ':' || c == ';' || c == '-' ||
			c == '_' || c == '\\' || c == '\t' || c == '\n' || c == '\''
	}

	cleanedText := strings.ReplaceAll(text, "'", "")

	splitedText := strings.FieldsFunc(strings.ToLower(cleanedText), splitFunc)

	for _, w := range splitedText {
		if _, ok := words[w]; !ok {
			words[w] = 1
		} else {
			words[w] += 1
		}

	}

	return words
}

func main() {
	input1 := "The quick brown fox jumps over the lazy dog."
	input2 := "Hello, hello! How are you doing today? Today is a great day."
	input3 := "Go, go, go! Let's learn Go programming."
	input4 := "  Spaces,   tabs,\t\tand\nnew-lines are ignored!  "
	fmt.Println(CountWordFrequency(input1))
	fmt.Println(CountWordFrequency(input2))
	fmt.Println(CountWordFrequency(input3))
	fmt.Println(CountWordFrequency(input4))
}
