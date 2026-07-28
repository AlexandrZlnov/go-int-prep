//Задача:
// Дан список чисел, в котором некоторые числа повторяются.
// Нужно оставить только уникальные значения, убрав все копии.
// Порядок чисел должен сохраниться.
// Было:  [1, 2, 2, 3, 1]
// Стало: [1, 2, 3]

// Решение:
// Временная сложность	O(n)
// Память	O(k)
// Принцип:
// Уникальные числа записываются в результирующий слай
// Проверка уникальности через мапу вида map[int]bool
// Если число есть в мапе они не пишеться в результирующий слайс

package main

import "fmt"

func main() {
	testNum := []int{1, 2, 2, 3, 1, 4, 5, 5, 6, 7, 7, 8, 8, 8, 9, 0, 7, 4, 3, 2, 1, 6}
	fmt.Printf("Слайс %v без дубликатов - %v\n", testNum, Unique(testNum))
}

func Unique(numbers []int) []int {

	// Вариант 1
	check := make(map[int]bool)
	result := []int{}

	for _, num := range numbers {
		if !check[num] {
			check[num] = true
			result = append(result, num)
		}
	}

	/*
		// Вариант 2
		check := make(map[int]struct{})
		result := []int{}

		for _, n := range numbers {
			if _, ok := check[n]; ok {
				continue
			}
			check[n] = struct{}{}
			result = append(result, n)
		}
	*/

	return result
}
