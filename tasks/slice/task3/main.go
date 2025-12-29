// Собес: Авито Платформа
/*
Максимальные элементы в массиве
Условие задачи: Дан массив чисел nums и некоторое неизвестное науке число K.
Нужно написать функцию, которая вынимает K самых больших чисел из массива nums.

	Пример:
	# ввод
	nums = [100, 50, 0, 150, 100, 0, –30, 70]
	k = 3

	# ожидаемый вывод (в любом порядке)
	expected = [100, 150, 100]

Алгоритмическая сложность:
Бейзлайн — сортировка + слайс (N + logN);
Необходимо обогнать этот вариант по сложности.
*/

package main

import (
	"container/heap"
	"fmt"
	//"slices"		// для варианта 2
)

func main() {
	nums := []int{100, 50, 0, 150, 100, 0, -30, 70}
	k := 3

	fmt.Println(topK(nums, k)) // вызов варанта 1
	// fmt.Println(max(nums, k))   // вызов варианта 2

}

// Вариант 1
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func topK(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}

	if k >= len(nums) {
		return nums
	}

	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
			//fmt.Printf("after Push i = %d, heap = %v\n", i, *h)
		}
	}

	return *h
}

// Вариант 2 - плохой, так не стоит
// func max(nums []int, k int) []int {
// 	slices.Sort(nums)

// 	return nums[len(nums)-k:]
// }
