package main

func search(nums []int, target int) int {
	res := -1
	if len(nums) == 0 {
		return res
	}

	left := 0
	right := len(nums) - 1
	middle := len(nums) / 2

	for left <= right {
		if nums[middle] == target {
			res = middle
			break
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
		middle = (right + left) / 2
	}
	return res
}
