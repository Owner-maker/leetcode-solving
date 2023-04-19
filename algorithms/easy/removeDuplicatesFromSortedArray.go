package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastIndex := 0
	currentNum := nums[0]

	for i := 0; i < len(nums); i++ {
		if currentNum != nums[i] {
			currentNum = nums[i]
			nums[i], nums[lastIndex+1] = nums[lastIndex+1], nums[i]
			lastIndex++
			//i++
		}
	}
	return lastIndex + 1
}

func main() {
	res := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Print(res)
}
