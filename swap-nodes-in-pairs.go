package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	var pre *ListNode
	for {
		if head != nil && head.Next != nil {
			node := head.Next
			head.Next = node.Next
			node.Next = head
			if pre != nil {
				pre.Next = node
			}
			pre = head
			head = head.Next
		} else {
			return res
		}
	}
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	res := swapPairs(list)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
