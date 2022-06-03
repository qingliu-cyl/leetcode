package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	now := head.Next
	pre := head
	head.Next = nil

	for now != nil {
		next := now.Next
		now.Next = pre
		pre = now
		now = next
	}

	return pre
}

func main() {

}
