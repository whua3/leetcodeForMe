package array

// 双指针指向头尾，高度更矮的边向内移动
func maxArea(height []int) int {
	length := len(height)
	var maxAreaVal int
	for i, j := 0, length-1; i < j; {
		currentArea := min(height[i], height[j]) * (j - i)
		if currentArea > maxAreaVal {
			maxAreaVal = currentArea
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxAreaVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
