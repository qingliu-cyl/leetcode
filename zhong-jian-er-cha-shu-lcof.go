package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	findFitst := func(list []int) int {
		for _, i := range preorder {
			for j := 0; j < len(list); j++ {
				if list[j] == i {
					return i
				}
			}
		}

		return -1
	}

	var findChild func(list []int, target *TreeNode)
	findChild = func(list []int, target *TreeNode) {
		if len(list) == 0 {
			return
		}
		// 根据中序遍历获取左节点
		var left []int
		var right []int
		for idx, _ := range list {
			if list[idx] == target.Val {
				left = list[:idx]
				right = list[idx+1:]
			}
		}

		if len(left) > 0 {
			leftVal := findFitst(left)
			leftNode := &TreeNode{
				Val: leftVal,
			}
			target.Left = leftNode
			findChild(left, leftNode)
		}

		if len(right) > 0 {
			rightVal := findFitst(right)
			rightNode := &TreeNode{
				Val: rightVal,
			}
			target.Right = rightNode
			findChild(right, rightNode)
		}

	}

	findChild(inorder, root)
	return root
}

//官解
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main() {
	buildTree([]int{7,-10,-4,3,-1,2,-8,11}, []int{-4,-10,3,-1,7,11,-8,2})
}
