package other

// 不能用递归函数，会超时
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	var res int
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}
