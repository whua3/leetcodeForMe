package other

import "math"

func maxProfit(prices []int) int {
	var minPrice, maxProfile = math.MaxInt64, 0

	for _, item := range prices {
		if item < minPrice {
			minPrice = item
		} else {
			if item-minPrice > maxProfile {
				maxProfile = item - minPrice
			}
		}
	}

	return maxProfile
}
