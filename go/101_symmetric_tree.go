package main

// Solution 1

func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, invertTree(root.Right))
}

// Solution 2

func isSymmetric2(root *TreeNode) bool {
	return isInvertedEqual(root.Left, root.Right)
}

func isInvertedEqual(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) ||
		(p != nil && q != nil && p.Val == q.Val && isInvertedEqual(p.Left, q.Right) && isInvertedEqual(p.Right, q.Left))
}
