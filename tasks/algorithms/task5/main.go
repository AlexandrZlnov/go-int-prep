// Задача:
// Посчитать количество каждого символа в строке.
//

// Решение:
// Временная сложность	O(n)
// Память	O(k) - мапа хранит только уникальные заначения, тоесть k<=n
// Принцип: символы строки записываются в мапу как ключи, а в значении счетчик
// по каждому символу
package main

import "fmt"

func main() {
	words := []string{
		"разработка",
		"ННрррРРЙх/",
		"language",
		"Concurrency",
	}

	for _, s := range words {
		result := SymbolCount(s)
		fmt.Printf("\nКоличество символов в строке - \"%s\"\n", s)
		for k, v := range result {
			fmt.Printf("%q: %d\n", k, v)
		}
	}
}

func SymbolCount(str string) map[rune]int {
	result := make(map[rune]int)

	for _, r := range str {
		result[r]++

	}
	return result
}
