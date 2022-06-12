package main

import (
	"bufio"
	"math/big"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(A, B *TreeNode) bool {
	if  A == nil && B != nil || A != nil && B != nil &&A.Val != B.Val{
		return false
	}
	if A == nil && B == nil || A != nil && B == nil {
		return true
	}
	return dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	list := make([]*TreeNode, 1)
	list[0] = A

	for i := 0; i < len(list); i++ {
		if list[i].Left != nil {
			list = append(list, list[i].Left)
		}

		if list[i].Right != nil {
			list = append(list, list[i].Right)
		}

		if list[i].Val == B.Val {
			if dfs(list[i], B) {
				return true
			}
		}
	}

	return false
}
