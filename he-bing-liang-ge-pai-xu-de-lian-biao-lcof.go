package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head

	p1, p2 := l1, l2

	for p1 != nil || p2 != nil {
		node.Next = &ListNode{}
		node = node.Next
		if p2 == nil || p1 != nil && p1.Val < p2.Val {
			node.Val = p1.Val
			p1 = p1.Next
		} else {
			node.Val = p2.Val
			p2 = p2.Next
		}

	}

	return head.Next
}

func main() {
	fmt.Println(mergeTwoLists(nil, &ListNode{Val: 0}))
}
