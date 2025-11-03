package array

import "sort"

// 按区间的头进行排序，然后遍历合并

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按左边界升序排序, 要记住这个sort函数怎么写的
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge2(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		left, right := intervals[i][0], intervals[i][1]
		if len(ans) == 0 || ans[len(ans)-1][1] < left {
			ans = append(ans, intervals[i])
		} else {
			if ans[len(ans)-1][1] < right {
				ans[len(ans)-1] = []int{ans[len(ans)-1][0], right}
			}
		}
	}

	return ans
}
