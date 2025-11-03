package array

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	maxLen := 1
	currentLen := 1

	tmp := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > tmp {
			currentLen++
			tmp = nums[i]
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 1
			tmp = nums[i]
		}
	}

	return maxLen
}
