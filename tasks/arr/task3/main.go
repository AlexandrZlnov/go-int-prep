// Задача с Leetcode 243
// Дан массив слов и два слова word1 и word2. Найти минимальное расстовяние между ними в массиве.
// Пример: words["practice", "makes", "pefect", "coding", "makes"]
// word1 = "coding" word2="makes"
// coding на индексе 3
// makes на индеке 1, 4
// |3-1| = 2
// |3-4| = 1 - минимум

package main

import (
	"fmt"
)

// Вариант 1
func wordDistance(words []string, w1, w2 string) int {
	last := -1
	minDist := len(words)

	for i, word := range words {
		if word != w1 && word != w2 {
			continue
		}

		if last != -1 && words[last] != word {
			dist := i - last
			if dist < minDist {
				minDist = dist
			}
		}

		last = i
	}

	return minDist
}

// Вариант 2
/*
func wordDistance(words []string, w1, w2 string) int {
	var minDist int = len(words)
	var index1, index2 = -1, -1

	if w1 == w2 || len(w1) == 0 || len(w2) == 0 {
		fmt.Printf("Неверный форма слов: %s, %s", w1, w2)
		return -1
	}

	for i, word := range words {
		if word == w1 {
			index1 = i
		} else if word == w2 {
			index2 = i
		}

		if index1 != -1 && index2 != -1 {
			dist := index1 - index2
			if dist < 0 {
				dist = -dist
			}

			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}
*/

func main() {
	words := [...]string{"practice", "makes", "pefect", "coding", "makes"}
	word1 := "coding"
	word2 := "makes"

	fmt.Printf("Минимальное расстояние между словами %s & %s = %d \n", word1, word2, wordDistance(words[:], word1, word2))
}
