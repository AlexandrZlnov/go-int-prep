// Задача
// Найти подмассив с максимальной суммой.
// [-2,1,-3,4,-1,2,1,-5,4]
// -> 6
// Потому что: [4,-1,2,1]

// Решение:
// Используем алгоритм Кадане
// Принцип:
// Проходим циклом по массиву
// Если текущая сумма < 0 то имеет смысл начать суммирвоание сначала
// А если > 0 суммируем следующий элемент
// На каждом шаге сравниваем с максимальной суммо и обновляем ее если текущая больше максимальной

package main

import "fmt"

func main() {
	testArr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubarray(testArr))
}

func maxSubarray(a []int) int {
	currentSum := a[0]
	maxSum := a[0]

	for i := 1; i < len(a); i++ {

		if currentSum < 0 {
			currentSum = a[i]
		} else {
			currentSum += a[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// Альтернативное решение, но с тем же принципом
/*
func maxSubarray(nums []int) int {
    maxSum, currSum := nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        currSum = max(nums[i], currSum+nums[i])
        maxSum = max(maxSum, currSum)
    }
    return maxSum
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
*/
