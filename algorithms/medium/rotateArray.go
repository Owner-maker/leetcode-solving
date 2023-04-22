package main

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}

	length := len(nums)
	k %= length

	edge := length - k
	for i := 0; i < edge/2; i++ {
		nums[i], nums[edge-i-1] = nums[edge-i-1], nums[i]
	}

	j := length - 1
	for i := edge; i < edge+k/2; i++ {
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}

	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}
