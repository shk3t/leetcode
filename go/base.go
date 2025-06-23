package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func main() {
// 	tree := &TreeNode{
// 		Val: 4,
// 		Left: &TreeNode{
// 			Val: 2,
// 			Left: &TreeNode{
// 				Val: 1,
// 			},
// 			Right: &TreeNode{
// 				Val: 3,
// 			},
// 		},
// 		Right: &TreeNode{
// 			Val: 6,
// 		},
// 	}
//
// 	diff := getMinimumDifference(tree)
// 	fmt.Println(diff)
// }
