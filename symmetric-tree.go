package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var check func(node1, node2 *TreeNode) bool
	check = func(node1, node2 *TreeNode) bool {
		// 都为nil一定对称
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil {
			return false
		}

		return node1.Val == node2.Val && check(node1.Left, node2.Right) &&
			check(node1.Right, node2.Left)
	}

	return check(root, root)
}

func main() {

}
