package array

// 官方题解方法一：哈希表

// 先用hash表存储nums，然后遍历Map，找到一个没有前一个数在map中的数开始找
func longestConsecutive(nums []int) int {
	// 使用hash表存储nums
	numsMap := make(map[int]bool, 0)
	for _, item := range nums {
		numsMap[item] = true
	}

	var ans int
	for k := range numsMap {
		if numsMap[k-1] {
			// 找到没有前一个数的k，从头开始 (这一步最关键)
			continue
		}

		currentLen := 1
		currentVal := k

		for numsMap[currentVal+1] {
			currentLen++
			currentVal++
		}

		if currentLen > ans {
			ans = currentLen
		}
	}
	return ans
}
