package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node.Left != nil {
			traversal(node.Left)
		}
		res = append(res, node.Val)
		if node.Right != nil {
			traversal(node.Right)
		}
	}

	traversal(root)

	return res
}

func main() {

}
