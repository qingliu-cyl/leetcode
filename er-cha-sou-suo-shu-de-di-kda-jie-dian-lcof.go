package main

func kthLargest(root *TreeNode, k int) int {
	var sort func(node *TreeNode)
	list := make([]int, 0)
	sort = func(node *TreeNode) {
		if node == nil {
			return
		}

		sort(node.Right)
		list = append(list, node.Val)
		sort(node.Left)
	}

	sort(root)
	return list[k-1]
}