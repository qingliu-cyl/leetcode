package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {

	res := make([][]int, 0)
	var search func(node *TreeNode, list []int)
	search = func(node *TreeNode, list []int) {
		newList := make([]int, len(list))
		copy(newList, list)
		newList = append(newList, node.Val)

		if node.Right == nil && node.Left == nil {
			sum := 0
			for _, i := range newList {
				sum += i
			}

			if sum == target {
				res = append(res, newList)
			}
			return
		}

		if node.Left != nil {
			search(node.Left, newList)
		}

		if node.Right != nil {
			search(node.Right, newList)
		}
	}

	if root == nil {
		return res
	}
	search(root, []int{})

	return res
}
