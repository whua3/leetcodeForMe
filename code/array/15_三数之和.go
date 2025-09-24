package array

import "sort"

// 官方题解一：双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 先排序
	ans := make([][]int, 0)
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return ans
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := length - 1

		for left < right {
			val := nums[i] + nums[left] + nums[right]
			if val == 0 {
				// found
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if val > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
