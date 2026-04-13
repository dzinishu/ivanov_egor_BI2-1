package main

import "fmt"

// рекурсивная сумма
func Sum(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + Sum(nums[1:])
}

// рекурсивный минимум
func Min(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	minRest := Min(nums[1:])

	if nums[0] < minRest {
		return nums[0]
	}

	return minRest
}

// рекурсивный максимум
func Max(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	maxRest := Max(nums[1:])

	if nums[0] > maxRest {
		return nums[0]
	}

	return maxRest
}

func main() {

	numbers := []int{5, 2, 9, 7, 1}

	fmt.Println("Минимум:", Min(numbers))
	fmt.Println("Максимум:", Max(numbers))
	fmt.Println("Сумма:", Sum(numbers))
}
