package main

// Solution 1

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right-1 {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] >= target {
		return left
	}
	return right
}

// Solution 2

func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left != right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// Tip 1: mid usually is floor rounded
// Tip 2: substitute nums[mid] with side-index to estimate target position
// Tip 3: do not forget about edge case in case when left != right
