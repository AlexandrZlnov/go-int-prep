// из списка задач к собеседования - Avtostopom-po-Go
//Условие задачи
//Дан массив целых чисел nums и целое число k. Нужно написать функцию,
//которая вынимает из массива nums k наиболее часто встречающихся элементов.

//Задание:
//# ввод
//nums = [1,1,1,2,2,3]
//k = 2
//# вывод (в любом порядке)
//[1, 2]

package main

import (
	"fmt"
	//"sort"
)

type pair struct {
	key   int
	value int
}

//вариан 1
// func topKFrequent(nums []int, k int) []int {
// 	freq := make(map[int]int)
// 	//distNew := []alfa{}
// 	//out := []int{}

// 	for _, i := range nums {
// 		freq[i]++
// 	}

// 	pairs := make([]pair, 0, len(freq))
// 	for num, count := range freq {
// 		pairs = append(pairs, pair{num, count})
// 	}

// 	fmt.Println(pairs)

// 	sort.Slice(pairs, func(i, j int) bool {
// 		return pairs[i].value > pairs[j].value
// 	})

// 	out := make([]int, 0, k)
// 	for i := 0; i < k && i < len(pairs); i++ {
// 		out = append(out, pairs[i].key)
// 	}

// 	// fmt.Println(out)

// 	return out

// }

// вариант 2 - bucket sort
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	freq := make([][]int, len(nums)+1)
	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	out := make([]int, 0, k)

	for i := len(freq) - 1; i >= 0; i-- {
		for _, n := range freq[i] {
			out = append(out, n)
			if len(out) == k {
				return out
			}
		}

	}

	return out
}

func main() {
	//arr := []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 4, 4}
	arr := []int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 7, 7, 7, 7, 7, 7, 7, 4, 4, 4, 4, 4, 4, 4}
	k := 3
	fmt.Println("Top", k, "> > >", topKFrequent(arr, k))

}
