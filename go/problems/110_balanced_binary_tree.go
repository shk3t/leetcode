package problems

const disbalanced = -1

func isBalanced(root *TreeNode) bool {
	return getHeightIfBalanced(root) != disbalanced
}

func getHeightIfBalanced(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeightIfBalanced(node.Left)
	rightHeight := getHeightIfBalanced(node.Right)
	if leftHeight == disbalanced || rightHeight == disbalanced || abs(rightHeight-leftHeight) > 1 {
		return disbalanced
	}

	return 1 + max(leftHeight, rightHeight)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
