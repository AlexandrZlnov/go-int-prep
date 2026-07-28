// Задача: Яндекс
// Дан непустой массив из нулей и единиц. Нужно определить,
// какой максимальный по длине подынтервал единиц можно получить,
// удалив (пропустив) ровно один элемент массива.
// Вернуть 0, если такого подынтервала не существует.
// Удалять один элемент из массива обязательно!
// Пример: [0, 1, 1, 0, 1] => 3

package main

import (
	"fmt"
)

func SubArr(arr []int) int {
	currentLenght := 0 // 1
	maxLenght := 0     // 1
	zeroCount := 0     // 1
	left := 0          // 0
	right := 0         // 3

	for right = 0; right < len(arr); right++ {
		if arr[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if arr[left] == 0 {
				zeroCount--
			}
			left++
		}

		if zeroCount <= 1 {
			currentLenght = right - left
			if currentLenght > maxLenght {
				maxLenght = currentLenght
			}
		}

	}
	return maxLenght

}

func main() {
	var arr [][]int = [][]int{
		{0, 1, 1, 0, 1},          // 3
		{0, 1},                   // 1
		{1, 0},                   // 1
		{1, 0, 1},                // 2
		{1, 1, 0, 1, 1},          // 4
		{1, 0, 1, 0, 1},          // 2
		{0, 1, 1, 1},             // 3
		{1, 1, 1, 0},             // 3
		{0, 1, 0, 1, 0},          // 2
		{1, 0, 1, 1, 0, 1, 1, 1}, // 5
		{1, 1, 1, 1},             // 3
		{1, 0, 0, 1},             // 1
		{0, 0, 0, 0},             // 0
	}

	for _, s := range arr {
		fmt.Printf("Для слайса %v ответ = %d\n", s, SubArr(s))
	}
}
