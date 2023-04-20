package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var place int

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			place++
			continue
		}

		if i+1 != len(nums) {
			nums[place], nums[i+1] = nums[i+1], nums[place]
			if nums[place] != nums[i+1] {
				place++
			}
		}

	}

	return place
}
