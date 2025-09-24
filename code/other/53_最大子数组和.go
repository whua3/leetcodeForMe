package other

import "math"

func maxSubArray(nums []int) int {
	var sumValue int
	var maxValue int = math.MinInt64
	for _, num := range nums {
		if sumValue < 0 {
			sumValue = num
		} else {
			sumValue += num
		}
		if sumValue > maxValue {
			maxValue = sumValue
		}
	}
	return maxValue
}
