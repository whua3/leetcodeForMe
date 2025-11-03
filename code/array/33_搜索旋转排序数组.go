package array

// 官方题解方法一：二分查找

// 旋转之后的数组在旋转点左右两边分别都是有序的
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左边有序
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右边有序
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
