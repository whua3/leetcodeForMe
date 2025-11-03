package array

// 使用一个map存储nums，判断target-nums[i]是否存在
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, item := range nums {
		if v, ok := numsMap[target-item]; ok {
			return []int{i, v}
		}
		numsMap[item] = i
	}

	return nil
}
