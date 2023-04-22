package main

// doesn't work for cases like [4,2,4,0,0,3,0,5,1,0]
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	length := len(nums)
	edge := length - 1

	for i := 0; i < edge; i++ {
		if nums[edge] == 0 {
			edge--
			i--
			continue
		}

		if nums[i] == 0 && nums[edge] != 0 {
			nums[i], nums[edge] = nums[edge], nums[i]
			edge--
		}
	}

	if edge == length-1 || edge == length-2 {
		return
	}

	for i := 0; i < edge; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}

func main() {
	moveZeroes([]int{0, 0, 1})
}
