package main

import "encoding/asn1"

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

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, lists := range lists {
		res = mergeTwoLists(res, lists)
	}
	return res
}

func main() {

}
