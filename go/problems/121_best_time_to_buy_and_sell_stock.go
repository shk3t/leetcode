package problems

import (
	"math"
)

func maxProfit(prices []int) int {
	best := 0
	low, high := math.MaxInt, math.MaxInt

	for _, v := range prices {
		if v > high {
			high = v
		} else if v < low {
			best = max(best, high-low)
			low, high = v, v
		}
	}

	return max(best, high-low)
}
