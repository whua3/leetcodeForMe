package array

// 二分查找，分段有序，比较high
func findMin(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2

		if nums[high] > nums[mid] { // 因为数组整体是升序的，要比较high而不是low
			// 右边有序
			high = mid
		} else {
			low = mid + 1 // 注意这个是mid+1，不是mid
		}
	}

	return nums[low]
}
