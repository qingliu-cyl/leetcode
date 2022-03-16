package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)

	var addVal func(node *TreeNode)
	addVal = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			addVal(node.Left)
		}
		res = append(res, node.Val)
		addVal(node.Right)
	}

	addVal(root)

	if len(res) == 0 {
		return true
	}

	min := res[0]
	for i := 1; i < len(res); i++ {
		if res[i] <= min {
			return false
		}
		min = res[i]
	}

	return true
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(isValidBST(root))
}
