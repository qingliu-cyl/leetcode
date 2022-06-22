package main

func maxDepth(root *TreeNode) int {

	max := 0
	var search func(node *TreeNode, deep int)
	search = func(node *TreeNode, deep int) {
		if node == nil {
			if deep > max {
				max = deep
			}
			return
		}

		deep++
		search(node.Right, deep)
		search(node.Left, deep)
	}

	search(root, 0)
	return max
}