package main

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	pre := head
	node := head.Next
	for node != nil {
		if node.Val == val {
			pre.Next = node.Next
			return head
		}
		pre = node
		node = node.Next
	}

	return nil
}
