package main

func rotate(nums []int, k int) {
	sz := len(nums)
	k = k % len(nums)

	for k > 0 {
		for i := sz - k - 1; i >= 0; i-- {
			nums[i], nums[i+k] = nums[i+k], nums[i]
		}
		sz, k = k, (k-sz%k)%k
	}
}
