package main

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
		}
	}
	return lastIndex + 1
}
