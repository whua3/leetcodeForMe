package other

func maxProfile2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var maxProfile int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			maxProfile += prices[i+1] - prices[i]
		}
	}

	return maxProfile
}
