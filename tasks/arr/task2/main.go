// Собес: Авито Платформа
// Задача:
// Дано 2 отсортированных (по возрастанию) массива А и В длины М и N.
// Нужно слить их в один отсортированный (по возрастанию) массив,
// состоящий из элементов первых двух.
// -- Пример 1
// Ввод: [1, 2, 5] [1, 2, 3, 4, 6]
// Вывод: [1, 1, 2, 2, 3, 4, 5, 6]
// -- Пример 2
// Ввод: [4, 7, 13] [3, 5, 8, 9, 11]
// Вывод: [3, 4, 5, 7, 8, 9, 11, 13]

package main

import (
	"fmt"
)

func main() {
	arr1_1 := []int{1, 2, 5}
	arr1_2 := []int{1, 2, 3, 4, 6}

	arr2_1 := []int{4, 7, 13}
	arr2_2 := []int{3, 5, 8, 9, 11}

	fmt.Println(merge(arr1_1, arr1_2))
	fmt.Println(merge(arr2_1, arr2_2))

}

// Вариант 1
func merge(arr1, arr2 []int) []int {
	merge := make([]int, len(arr1)+len(arr2))
	i := 0
	j := 0

	for i < len(arr1) || j < len(arr2) {
		if j == len(arr2) && i < len(arr1) {
			merge[i+j] = arr1[i]
			i++
		} else if i == len(arr1) && j < len(arr2) {
			merge[i+j] = arr2[j]
			j++
		} else if arr1[i] > arr2[j] {
			merge[i+j] = arr2[j]
			j++
		} else if arr1[i] < arr2[j] {
			merge[i+j] = arr1[i]
			i++
		} else {
			merge[i+j] = arr1[i]
			merge[i+j+1] = arr2[j]
			i++
			j++
		}

	}

	return merge
}

// Вариант 2
// 	for range len(arr1) + len(arr2) {
// 		if j == len(arr2) && i < len(arr1) {
// 			merge[i+j] = arr1[i]
// 			i++
// 			continue
// 		}
// 		if i == len(arr1) && j < len(arr2) {
// 			merge[i+j] = arr2[j]
// 			j++
// 			continue
// 		}
// 		if i == len(arr1) && j == len(arr2) {
// 			break
// 		}
// 		if arr1[i] > arr2[j] {
// 			merge[i+j] = arr2[j]
// 			j++
// 			continue
// 		}
// 		if arr1[i] < arr2[j] {
// 			merge[i+j] = arr1[i]
// 			i++
// 			continue
// 		}
// 		if arr1[i] == arr2[j] {
// 			merge[i+j] = arr1[i]
// 			merge[i+j+1] = arr2[j]
// 			i++
// 			j++
// 			continue
// 		}

// 	}

// 	return merge
// }

// Вариант 3 - так не стоит, это излишне.
// func merge(arr1, arr2 []int) []int {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		for _, v := range arr1 {
// 			ch1 <- v
// 		}
// 		close(ch1)
// 	}()

// 	go func() {
// 		for _, v := range arr2 {
// 			ch2 <- v
// 		}
// 		close(ch2)
// 	}()

// 	result := make([]int, 0, len(arr1)+len(arr2))

// 	v1, ok1 := <-ch1
// 	v2, ok2 := <-ch2

// 	for ok1 || ok2 {
// 		switch {
// 		case !ok2 || (ok1 && v1 <= v2):
// 			result = append(result, v1)
// 			v1, ok1 = <-ch1
// 		default:
// 			result = append(result, v2)
// 			v2, ok2 = <-ch2
// 		}
// 	}

// 	return result
// }
