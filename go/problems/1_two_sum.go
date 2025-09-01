package problems

func twoSum(nums []int, target int) []int {
	need := make(map[int]int)

	for i, n := range nums {
		if j, ok := need[n]; ok {
			return []int{j, i}
		} else {
			need[target-n] = i
		}
	}

	return nil
}
