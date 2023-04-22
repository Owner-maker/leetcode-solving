package main

// doesn't work for cases like [0,1,1,0]
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	length := len(nums)
	var zeros int
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			zeros++
			continue
		}
		if i+1 != length {
			if nums[i] != 0 && nums[i+1] == 0 {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				zeros++
			}
		}
	}

	if zeros == 0 {
		return
	}

	j := length - 1
	for i := zeros; i < (zeros+length)/2; i++ {
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}

	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
