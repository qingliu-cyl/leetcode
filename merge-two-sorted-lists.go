package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	list := &ListNode{}
	head := list
	last := head
	for {
		if list1 == nil && list2 != nil {
			list.Val = list2.Val
			list2 = list2.Next
		} else if list1 != nil && list2 == nil {
			list.Val = list1.Val
			list1 = list1.Next
		} else if list1 == nil && list2 == nil {
			last.Next = nil
			return head
		} else {
			if list1.Val < list2.Val {
				list.Val = list1.Val
				list1 = list1.Next
			} else {
				list.Val = list2.Val
				list2 = list2.Next
			}
		}
		list.Next = &ListNode{}
		last = list
		list = list.Next
	}
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list := mergeTwoLists(list1, list2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
