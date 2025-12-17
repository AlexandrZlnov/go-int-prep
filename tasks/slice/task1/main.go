package main

import (
	//"errors"
	"fmt"
)

func addNum(nums []int) {
	nums = append(nums, 4)
}

func addNums(nums []int) {
	fmt.Printf("In addNums nums - %p\n", nums)
	nums = append(nums, 5, 6)
	fmt.Printf("In addNums after append nums - %p\n", nums)
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Printf("Isxod nums - %p\n", nums)

	addNum(nums[0:2])
	fmt.Println(nums)
	fmt.Printf("After addNum nums - %p\n", nums)

	addNums(nums[0:2])
	fmt.Println(nums)
	fmt.Printf("After addNum nums - %p\n", nums)
}
