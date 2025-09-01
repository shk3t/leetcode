package problems

func hasPathSum(root *TreeNode, targetSum int) bool {
	return root != nil &&
		((root.Val == targetSum && root.Right == nil && root.Left == nil) ||
			hasPathSum(root.Left, targetSum-root.Val) ||
			hasPathSum(root.Right, targetSum-root.Val))
}
