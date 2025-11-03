package other

func quickSort(nums []int, low, high int) {
	i, j := low, high
	tmp := nums[i]

	for i < j {
		for i < j && tmp < nums[j] {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && tmp > nums[i] {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = tmp
	if low < i {
		quickSort(nums, low, i-1)
	}
	if high > i {
		quickSort(nums, i+1, high)
	}
}
