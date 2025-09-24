package array

func findLengthOfLCIS(nums []int) int {
	var ans int
	if len(nums) == 1 {
		return 1
	}
	var start int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			// 停止
			if (i + 1 - start) > ans {
				ans = i + 1 - start
			}
			start = i + 1
		} else if i == len(nums)-2 {
			if (i + 2 - start) > ans {
				ans = i + 2 - start
			}
		}
	}
	return ans
}
