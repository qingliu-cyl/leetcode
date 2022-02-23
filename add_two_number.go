package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	res := &ListNode{} // 返回值中的最后一个结构体
	head := res        // 返回的结构体的头部
	carry := 0         // 进位变量
	pre := &ListNode{} // 上一个结构体， 用于在最后一次循环将next置nil
	for {
		// 当l1 l2 已经carry都为nil时，说明已经不用计算了， 将最后一个结构体的next置为nil
		if l1 == nil && l2 == nil && carry == 0 {
			pre.Next = nil
			return head
		}

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		res.Val = carry % 10
		res.Next = &ListNode{}
		pre = res
		res = res.Next
		carry = carry / 10
	}
}
