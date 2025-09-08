package problems

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{prefixSum: make([]int, len(nums)+1)}
	sum := 0
	for i, v := range nums {
		sum += v
		arr.prefixSum[i+1] = sum
	}
	return arr
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}
