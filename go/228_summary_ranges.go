package main

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	beg, prev := nums[0], nums[0]
	ranges := []string{}

	for _, v := range nums[1:] {
		if v != prev+1 {
			ranges = append(ranges, intervalString(beg, prev))
			beg = v
		}
		prev = v
	}
	ranges = append(ranges, intervalString(beg, prev))

	return ranges
}

func intervalString(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	} else {
		return strconv.Itoa(a) + "->" + strconv.Itoa(b)
	}
}
