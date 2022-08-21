package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	nodeMap := make(map[int][]int)
	nodeMap[root.Val] = make([]int, 0)
	if root.Left != nil {
		nodeMap[root.Val] = append(nodeMap[root.Val], root.Left.Val)
	}

	if root.Right != nil {
		nodeMap[root.Val] = append(nodeMap[root.Val], root.Right.Val)
	}

	var reduce func(node *TreeNode, parent int)
	reduce = func(node *TreeNode, parent int) {
		if node == nil {
			return
		}

		nodeMap[node.Val] = []int{parent}

		if node.Left != nil {
			nodeMap[node.Val] = append(nodeMap[node.Val], node.Left.Val)
		}

		if node.Right != nil {
			nodeMap[node.Val] = append(nodeMap[node.Val], node.Right.Val)
		}

		reduce(node.Left, node.Val)
		reduce(node.Right, node.Val)
	}

	reduce(root.Left, root.Val)
	reduce(root.Right, root.Val)

	nodeStatus := make(map[int]bool)
	for k, _ := range nodeMap {
		nodeStatus[k] = false
	}
	nodeStatus[start] = true
	amountNum := 1
	time := 0
	amountNode := []int{start}

	for amountNum < len(nodeMap) {
		time++
		newAmountNode := make([]int, 0)
		for _, node := range amountNode {
			for _, j := range nodeMap[node] {
				if !nodeStatus[j] {
					amountNum++
					nodeStatus[j] = true
					newAmountNode = append(newAmountNode, j)
				}
			}
		}

		amountNode = newAmountNode
	}

	return time
}

func main() {

}
