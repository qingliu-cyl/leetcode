package main

func reversePrint(head *ListNode) []int {
	res1 := make([]int, 0)
	for head != nil {
		res1 = append(res1, head.Val)
		head = head.Next
	}

	res := make([]int, 0)
	for i := len(res1)-1; i >=0; i-- {
		res = append(res, res1[i])
	}
	return res
}

func main() {

}
