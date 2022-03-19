package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	last := &TreeNode{}

	// 拼接右子树
	var sort func(node *TreeNode)
	sort = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil {
			node.Left = node.Right
			node.Right = nil
			// node.Left为nil, 说明已经左子树已经遍历结束， 此时的做左子树便是下一个要
			// 拼接的右子树的前驱节点
			if node.Left == nil {
				last = node
			}
			sort(node.Left)
		} else {
			// 先处理左子树
			sort(node.Left)
			// 将右子树拼接到前驱节点
			last.Left = node.Right
			node.Right = nil
			// 处理拼接后的子树
			sort(last.Left)
		}
	}
	sort(root)

	// 左节点换为右节点
	next := root
	for next != nil {
		next.Right = next.Left
		next.Left = nil
		next = next.Right
	}
}

func main() {
	flatten(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	})
}
