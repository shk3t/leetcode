package main

func removeDuplicates2(nums []int) int {
	i, j := 1, 2
	for j < len(nums) {
		if nums[i] != nums[j] || nums[i-1] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return min(i + 1, len(nums))
}
