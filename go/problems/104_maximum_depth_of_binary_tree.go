package problems

// Solution 1

func maxDepth(root *TreeNode) int {
	path := appendLeftBranch([]*TreeNode{}, root)
	maxLen := len(path)

	for len(path) > 0 {
		cur := path[len(path)-1]
		if cur.Right != nil {
			path = appendLeftBranch(path, cur.Right)
		} else {
			path = removeRightBranchedEnding(path)
		}
		maxLen = max(maxLen, len(path))
	}

	return maxLen
}

func appendLeftBranch(path []*TreeNode, cur *TreeNode) []*TreeNode {
	for cur != nil {
		path = append(path, cur)
		cur = cur.Left
	}
	return path
}

func removeRightBranchedEnding(path []*TreeNode) []*TreeNode {
	i := len(path) - 2
	for i >= 0 && path[i].Right == path[i+1] {
		i--
	}
	return path[:i+1]
}

// Solution 2

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		maxDepth(root.Left),
		maxDepth(root.Right),
	) + 1
}
