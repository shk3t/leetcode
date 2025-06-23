package main

func containsNearbyDuplicate(nums []int, k int) bool {
	history := map[int]int{}
	for i, v := range nums {
		if j, ok := history[v]; ok && i - j <= k {
			return true
		}
		history[v] = i
	}
	return false
}
