package problems

func maxArea(height []int) int {
	maxLeftArea, maxRightArea := 0, 0
	maxLeftAreaIdx, maxRightAreaIdx := -1, -1

	for i, h := range height {
		area := h * (len(height) - i)
		if area > maxLeftArea {
			maxLeftArea = area
			maxLeftAreaIdx = i
		}
	}

	for i, h := range height {
		area := h * (i + 1)
		if area > maxRightArea && i != maxLeftAreaIdx {
			maxRightArea = area
			maxRightAreaIdx = i
		}
	}

	return min(height[maxLeftAreaIdx], height[maxRightAreaIdx]) * (maxRightAreaIdx - maxLeftAreaIdx)
}
