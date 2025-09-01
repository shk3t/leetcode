package problems

func removeElement(nums []int, val int) int {
	shift := 0

	for i, v := range nums {
		if v == val {
			shift--
		} else {
			nums[i+shift] = v
		}
	}

	return len(nums) + shift
}
