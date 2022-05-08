package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		l, dl := dfs(node.Left)
		r, dr := dfs(node.Right)

		count := dl + dr + 1
		total := l + r + node.Val
		ave := total / count
		if ave == node.Val {
			res+=1
		}
		return total, count
	}
	dfs(root)

	return res
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(averageOfSubtree(root))

}
