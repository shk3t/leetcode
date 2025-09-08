package problems

func singleNumber(nums []int) int {
	acc := 0
	for _, n := range nums {
		acc ^= n
	}
	return acc
}
