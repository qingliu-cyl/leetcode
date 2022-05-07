package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	resMap := make(map[*ListNode]bool)

	for headA != nil || headB != nil{
		nodeA := headA
		if nodeA != nil {
			if resMap[nodeA] {
				return nodeA
			} else {
				resMap[nodeA] = true
				headA = headA.Next
			}
		}

		nodeB := headB
		if nodeB != nil {
			if resMap[nodeB] {
				return nodeB
			} else {
				resMap[nodeB] = true
				headB = headB.Next
			}
		}
	}
	return nil
}

func main() {

}
