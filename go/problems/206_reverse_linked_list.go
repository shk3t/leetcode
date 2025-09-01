package problems

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev, next *ListNode = nil, nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}

	return prev
}
