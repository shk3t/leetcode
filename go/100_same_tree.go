package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) ||
		(p != nil && q != nil && p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
}
