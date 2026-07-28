// Задача:
//
// Что выведет программа?
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	addNum(nums[0:2])
	fmt.Println(nums) // 1 2 4

	addNums(nums[0:2])
	fmt.Println(nums) // 1 2 4
}

func addNum(nums []int) {
	nums = append(nums, 4) // 1 2 4
}

func addNums(nums []int) {
	nums = append(nums, 5, 6) // 1 2 5 6 - new_arr
}
