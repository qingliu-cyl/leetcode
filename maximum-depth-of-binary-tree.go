package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max := 0
	var getDepth func(node *TreeNode, depth int)
	getDepth = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		if node.Left == nil && node.Right == nil {
			if max < depth {
				max = depth
			}
			return
		}
		getDepth(node.Left, depth)
		getDepth(node.Right, depth)
	}

	getDepth(root, 0)
	return max
}

func main() {

}
