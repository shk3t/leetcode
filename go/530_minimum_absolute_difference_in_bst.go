package main

import (
	"math"
)

// Solution 1

func getMinimumDifference(root *TreeNode) int {
	sortedValues := putToSortedArray(root, []int{})

	minDiff := math.MaxInt
	for i := range len(sortedValues) - 1 {
		minDiff = min(minDiff, sortedValues[i+1]-sortedValues[i])
	}

	return minDiff
}

func putToSortedArray(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}
	arr = putToSortedArray(node.Left, arr)
	arr = append(arr, node.Val)
	return putToSortedArray(node.Right, arr)
}

// Solution 2

func getMinimumDifference2(root *TreeNode) int {
	_, minDiff := minDiffBranch(math.MinInt/2, root)
	return minDiff
}

func minDiffBranch(prevPrev int, node *TreeNode) (next int, minDiff int) {
	if node == nil {
		return math.MinInt / 2, math.MaxInt
	}
	prev, lMinDiff := minDiffBranch(prevPrev, node.Left)
	next, rMinDiff := minDiffBranch(node.Val, node.Right)
	return max(node.Val, next), min(lMinDiff, node.Val-max(prev, prevPrev), rMinDiff)
}
