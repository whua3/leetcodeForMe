package array

// 双指针

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	var ans int

	for left < right {
		maxLeft = max(height[left], maxLeft)
		maxRight = max(height[right], maxRight)

		if height[left] < height[right] {
			ans += maxLeft - height[left]
			left++
		} else {
			ans += maxRight - height[right]
			right--
		}
	}

	return ans
}
