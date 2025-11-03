package array

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

// 快排
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

// 实现一个堆排序

// 堆排序主函数
func heapSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 1. 构建最大堆（从最后一个非叶子节点开始）
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// 2. 排序阶段
	for i := n - 1; i > 0; i-- {
		// 将堆顶元素（最大值）与当前堆的最后一个元素交换
		nums[0], nums[i] = nums[i], nums[0]

		// 调整剩余元素为最大堆，堆大小减1
		heapify(nums, i, 0)
	}
}

// 堆化函数：确保以i为根的子树是最大堆
func heapify(nums []int, heapSize, i int) {
	largest := i     // 初始化最大值为根节点
	left := 2*i + 1  // 左子节点索引
	right := 2*i + 2 // 右子节点索引

	// 如果左子节点大于根节点，更新最大值索引
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}

	// 如果右子节点大于当前最大值，更新最大值索引
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	// 如果最大值不是根节点，则交换并递归堆化受影响的子树
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, heapSize, largest)
	}
}
