package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var search func(nodes []*TreeNode)
	search = func(nodes []*TreeNode) {
		treeList := make([]*TreeNode, 0) // 保存当前层级的子节点
		resList := make([]int, 0)        // 保存当前层级的val

		// 子节点加入队列
		for _, node := range nodes {
			if node.Left != nil {
				treeList = append(treeList, node.Left)
			}
			if node.Right != nil {
				treeList = append(treeList, node.Right)
			}
			resList = append(resList, node.Val)
		}
		// 添加当前层级的值
		res = append(res, resList)
		if len(treeList) > 0 {
			search(treeList)
		}
	}
	if root != nil {
		search([]*TreeNode{root})
	}

	return res
}

func main() {

}
