package main


func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	list := make([]*TreeNode, 1)
	list[0] = root

	for i := 0; i < len(list); i++ {
		l, r := list[i].Left, list[i].Right
		list[i].Right, list[i].Left = l ,r

		if list[i].Left != nil {
			list = append(list, list[i].Left)
		}

		if list[i].Right != nil {
			list = append(list, list[i].Right)
		}
	}

	return root
}