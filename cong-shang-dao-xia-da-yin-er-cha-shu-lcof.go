package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	list := make([]*TreeNode, 1)
	list[0] = root

	for i := 0; i < len(list); i++ {
		if list[i].Left != nil {
			list = append(list, list[i].Left)
		}

		if list[i].Right != nil {
			list = append(list, list[i].Right)
		}

		res = append(res, list[i].Val)
	}

	return res
}
