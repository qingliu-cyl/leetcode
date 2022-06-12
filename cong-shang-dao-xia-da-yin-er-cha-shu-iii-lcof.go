package main

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	list := make([][]*TreeNode, 1)
	list[0] = []*TreeNode{root}
	revert := false
	for i :=0 ; i< len(list); i++ {
		cow := make([]*TreeNode, 0)
		resCow := make([]int, 0)
		if revert {
			for j := len(list[i]) - 1; j >= 0;j-- {
				resCow = append(resCow, list[i][j].Val)
			}
		} else {
			for j := 0; j < len(list[i]);j++ {
				resCow = append(resCow, list[i][j].Val)
			}
		}

		for j := 0; j < len(list[i]);j++ {
			if list[i][j].Left != nil {
				cow = append(cow, list[i][j].Left)
			}

			if list[i][j].Right != nil {
				cow = append(cow, list[i][j].Right)
			}
		}

		revert = !revert

		if len(resCow) > 0 {
			res = append(res, resCow)
		}

		if len(cow) > 0 {
			list = append(list, cow)
		}
	}

	return res
}