package array

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	i, j := 0, len(nums)-1
	for i < j {
		mid := i + (j-i)/2
		last := nums[j]

		if nums[mid] < last {
			j = mid
		} else {
			i = mid + 1 // 注意这个是mid+1，不是mid
		}
	}
	return nums[i]
}
