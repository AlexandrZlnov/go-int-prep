// Задача:
// Собес: Магнит
// Слить 2 упорядоченных массива в 1 упорядоченный массив.
//
// Пример:
// Дано:
// [1, 2, 3, 4, 5]
// [4, 5, 6, 7, 8]
//
// Результат:
// [1, 2, 3, 4, 4, 5, 5, 6, 7, 8]
//
// func mergeSlices(a []int, b []int) []int {}
//
// Решение:

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}

	result := mergeSlices(arr1, arr2)

	fmt.Println("Result:", result)
}

func mergeSlices(a []int, b []int) []int {
	result := make([]int, 0, len(a)+len(b))

	i, j := 0, 0

	for i < len(a) && j < len(b) {

		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else if b[j] < a[i] {
			result = append(result, b[j])
			j++
		} else {
			result = append(result, a[i], b[j])
			i++
			j++
		}
	}

	if i >= len(a) {
		result = append(result, b[j:]...)
	} else if j >= len(b) {
		result = append(result, a[i:]...)

	}

	return result

}
