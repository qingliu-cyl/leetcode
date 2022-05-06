package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	addrMap := make(map[string]bool)
	addr := fmt.Errorf("%p", &head)
	addrMap[addr.Error()] = true
	node := head
	for node.Next != nil {
		addr := fmt.Errorf("%p", &(node.Next))
		if addrMap[addr.Error()] {
			return true
		}
		addrMap[addr.Error()] = true
		node = node.Next
	}
	return false
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
	fmt.Println(hasCycle(node4))


}
