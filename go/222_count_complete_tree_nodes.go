package main

// Solution 1

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// Solution 2

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	cur := root
	pos := 0

	for !(cur.Left == nil && cur.Right == nil) {
		if getLeftDepth(cur.Left) == getLeftDepth(cur.Right) {
			cur = cur.Right
			pos = 2*pos + 1
		} else {
			cur = cur.Left
			pos = 2 * pos
		}
	}

	return pos + (1 << (getLeftDepth(root) - 1))
}

func getLeftDepth(node *TreeNode) int {
	depth := 0
	for node != nil {
		node = node.Left
		depth++
	}
	return depth
}
