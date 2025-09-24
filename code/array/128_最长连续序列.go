package array

// 官方题解方法一：哈希表
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, 0)
	for _, item := range nums {
		numMap[item] = true
	}

	var currentLen, maxLen, currentVal int
	for k, _ := range numMap {
		if _, exist := numMap[k-1]; exist {
			continue
		}

		currentLen = 1
		currentVal = k
		for {
			if _, exist := numMap[currentVal+1]; exist {
				currentLen++
				currentVal++
			} else {
				if currentLen > maxLen {
					maxLen = currentLen
				}
				break
			}
		}
	}

	return maxLen
}
