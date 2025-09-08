package problems

func pivotIndex(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	sum := 0
	for i, v := range nums {
		sum += v
		prefixSum[i+1] = sum
	}

	for i := range len(nums) {
		pivot := prefixSum[i+1]
		if prefixSum[i] == sum-pivot {
			return i
		}
	}

	return -1
}
