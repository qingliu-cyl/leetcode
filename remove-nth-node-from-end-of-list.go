package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := 0
	list := head
	for list != nil {
		len++
		list = list.Next
	}

	if len == n {
		return head.Next
	}

	list = head
	for i := 0; i < (len - n - 1); i++ {
		list = list.Next
	}

	if n == 1 {
		list.Next = nil
	} else {
		next := list.Next.Next
		list.Next = next
	}
	return head
}

func main() {

}
