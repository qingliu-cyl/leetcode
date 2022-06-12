package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	list := make([][]*TreeNode, 1)
	list[0] = []*TreeNode{root}
	for i :=0 ; i< len(list); i++ {
		cow := make([]*TreeNode, 0)
		resCow := make([]int, 0)
		for j := 0; j < len(list[i]);j++ {
			if list[i][j].Left != nil {
				cow = append(cow, list[i][j].Left)
			}

			if list[i][j].Right != nil {
				cow = append(cow, list[i][j].Right)
			}
			resCow = append(resCow, list[i][j].Val)
		}
		if len(resCow) > 0 {
			res = append(res, resCow)
		}

		if len(cow) > 0 {
			list = append(list, cow)
		}
	}

	return res
}

func main()  {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrder(node))
}