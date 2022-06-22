package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res := true
	var search func(node *TreeNode, deep int) int
	search = func(node *TreeNode, deep int) int {
		if node == nil {
			return deep
		}

		deep++
		ldeep := search(node.Right, deep)
		rdeep := search(node.Left, deep)

		if ! (ldeep == rdeep || ldeep +1 == rdeep || rdeep +1 == ldeep) {
			res = false
		}

		if ldeep > rdeep {
			return ldeep
		}

		return rdeep

	}

	search(root, 0)
	return res
}
