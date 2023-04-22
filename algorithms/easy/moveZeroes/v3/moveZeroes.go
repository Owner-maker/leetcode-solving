package main

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	var lastZeroPlace int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastZeroPlace] = nums[lastZeroPlace], nums[i]
			lastZeroPlace++
		}
	}
}
