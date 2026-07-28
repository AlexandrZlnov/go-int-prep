// Задача
// Даны два массива (или слайса) целых чисел.
// Нужно найти пересечение — множество элементов, которые
// присутствуют в обоих массивах одновременно.
// [1,2,2,3]
// [2,2,4]
// -> [2]

// Решение
// Принцип
// Создается контрольная мапа по первому слайсу
// Значения 2го слайса сравниваются с мапой по ключу
// Если есть совпадение и значение в мапе == true добавляем ключ результирующий слайс
// У добавленного элемента в контрольной мапе значение меняем на false
// Нуждно для исключения повторений чисел в результирующем слайсе
// Упрощение: для исключения дублирований в результирующем слайсе вместо true/false можно
// сделать control := make(map[int]struct{}) и после добавления элемента удалять значение из
// контрольной мапы delete(control, n)

package main

import "fmt"

func main() {
	testArr1 := []int{1, 2, 2, 3, 70}
	testArr2 := []int{2, 2, 6, 8, 10, 14, 70}

	fmt.Println("Пересечения - ", Intersection(testArr1, testArr2))
}

func Intersection(a1, a2 []int) []int {
	result := []int{}
	control := make(map[int]bool)

	if len(a1) == 0 || len(a2) == 0 {
		fmt.Println("Нет пересечений")
		return nil
	}

	for _, n := range a1 {
		if _, ok := control[n]; !ok {
			control[n] = true
		}
	}

	for _, n := range a2 {
		val, ok := control[n]
		if ok && val == true {
			result = append(result, n)
			control[n] = false
		}
	}

	return result
}
