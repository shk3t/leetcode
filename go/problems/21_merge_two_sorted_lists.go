package problems

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resHead := new(ListNode)
	res := resHead

	for !(list1 == nil && list2 == nil) {
		if list2 == nil || (list1 != nil && list1.Val < list2.Val) {
			res.Next = list1
			res = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			res = list2
			list2 = list2.Next
		}
	}

	return resHead.Next
}
