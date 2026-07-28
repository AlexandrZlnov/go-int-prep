// Задача:
// Требуется реализовать функцию которая генерирует слайс длинной n уникальных и рандомных числе
// Исходный код:
/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqRandom(10))
}

func uniqRandom(n int) []int {

}
*/

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println(uniqRandom(10))
}

// Вариант 1
func uniqRandom(n int) []int {
	nums := make([]int, 0, n)

	seen := make(map[int]struct{})

	max := n * 5

	for len(nums) < n {
		new := rand.IntN(max)
		if _, ok := seen[new]; ok {
			continue
		}
		seen[new] = struct{}{}
		nums = append(nums, new)
	}
	return nums
}

// Вариант 2
/*
func uniqRandom(n int) []int {
	max := n * 5

	nums := make([]int, max)

	for i := range nums {
		nums[i] = i
	}

	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	return nums[:n]

}
*/
