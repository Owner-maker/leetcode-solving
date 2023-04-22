package main

import "fmt"

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	res := make([]int, len(nums), len(nums))

	left := 0
	right := len(nums) - 1
	currentIndex := right

	var leftValue int
	var rightValue int

	for left <= right {
		leftValue = nums[left] * nums[left]
		rightValue = nums[right] * nums[right]

		if leftValue > rightValue {
			res[currentIndex] = leftValue
			left++
		}
		if rightValue >= leftValue {
			res[currentIndex] = rightValue
			right--
		}
		currentIndex--
	}
	return res
}

func main() {
	f := sortedSquares([]int{-7, -3, 2, 3, 11})
	fmt.Print(f)
}
