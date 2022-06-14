package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	node := head
	nodeLen := 0

	for node != nil {
		nodeLen++
		node = node.Next
	}

	if k > nodeLen {
		return nil
	}
	node = head
	for i := 1; i <= nodeLen; i++ {
		if i == (nodeLen - k + 1) {
			return node
		}

		node = node.Next
	}

	return nil
}