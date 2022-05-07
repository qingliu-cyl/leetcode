package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	node := head.Next
	head.Next = nil
	for node != nil && node.Next != nil {
		next := node.Next
		node.Next = head
		head = node
		node = next
	}
	node.Next = head
	return node
}

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: nil,
	}
	node2 := &ListNode{
		Val: 2,
		Next: node1,
	}

	node3 := &ListNode{
		Val: 3,
		Next: node2,
	}

	node4 := &ListNode{
		Val: 4,
		Next: node3,
	}

	reverseList(node4)
}
