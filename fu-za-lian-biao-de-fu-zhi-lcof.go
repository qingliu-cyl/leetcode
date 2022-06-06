package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var nodeMap map[*Node]*Node
func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if newNode, ok := nodeMap[node]; ok {
		return newNode
	}

	newNode := &Node{
		Val: node.Val,
	}
	nodeMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)

	return newNode
}

func copyRandomList(head *Node) *Node {
	nodeMap = make(map[*Node]*Node)
	return deepCopy(head)
}

func main() {

}
