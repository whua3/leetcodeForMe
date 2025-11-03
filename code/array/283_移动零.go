package array

func MoveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func moveZeroes(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
