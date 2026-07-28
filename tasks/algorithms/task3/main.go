// Задача
// Дан массив чисел и target.
// Нужно вернуть индексы двух чисел, сумма которых равна target.

// Вариант 2
// Временная сложность	O(n)
// Память	O(n)
// Принцип действия:
// - Проходим по массиву через range:
// 		- Для текущего числа вычисляем diff = target - num
// 		- Смотрим в map: есть ли там diff?
// 		- Если есть → нашли пару! Возвращаем [индекс из map, текущий индекс из range]
//		- Если нет → запоминаем текущее число в map[число из слайса] = индек из слайка (из range) и идем дальше
// - Если прошли весь массив → возвращаем nil (решения нет)

package main

import "fmt"

func main() {
	tests := []struct {
		numbers []int
		target  int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 10},
		{[]int{20, 10, 40, 5, 15, 30}, 70},
		{[]int{5, 7, 9, 6, 1, 99, 10}, 15},
		{[]int{5, 7, 9, 6, 1, 99, 10}, 120},
	}

	for _, test := range tests {
		fmt.Printf("Target - %d дают числа с индексами %v\n", test.target, Summ(test.numbers, test.target))
	}
}

func Summ(numbers []int, target int) []int {
	m := make(map[int]int)

	for i, num := range numbers {
		diff := target - num

		if ind, ok := m[diff]; ok {
			return []int{ind, i}
		}

		m[num] = i

	}

	return nil

}

// Вариант 1
// Временная сложность O(n²)
// Память	O(1)
/*
package main

import "fmt"

func main() {
	tests := []struct {
		numbers []int
		target  int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 10},
		{[]int{20, 10, 40, 5, 15, 30}, 70},
		{[]int{5, 7, 9, 6, 1, 99, 10}, 15},
		{[]int{5, 7, 9, 6, 1, 99, 10}, 120},
	}

	for _, test := range tests {
		ind1, ind2 := Summ(test.numbers, test.target)
		fmt.Printf("Target - %d дают числа с индексами %d и %d\n", test.target, ind1, ind2)
	}
}

func Summ(numbers []int, target int) (int, int) {

	for ind1, number := range numbers {
		delta := target - number

		for ind2 := ind1 + 1; ind2 < len(numbers); ind2++ {
			if delta == numbers[ind2] {
				return ind1, ind2
			}

		}

	}

	return -1, -1 // Явно указываем что решение не найдено
}

*/
