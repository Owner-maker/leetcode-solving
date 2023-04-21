package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target < nums[0] {
			return 0
		}
		return target - nums[0]
	}

	left := 0
	right := len(nums) - 1
	middle := len(nums) / 2
	res := -1

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

	if res == -1 {
		if target < nums[middle] {
			return 0
		}
		return middle + 1
	}

	return res
}
