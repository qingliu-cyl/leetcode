package main

import "fmt"

func numTrees(n int) int {
	list := make([]int, n+1)
	list[0], list[1] = 1, 1

	// i为总长度
	for i := 2; i <= n; i++ {
		// j为root节点的位置
		for j := 1; j <= i; j++ {
			// 左子树的种类 * 右子树的种类
			// 左子树的数量等于j-1
			//右子树的数量等于i-j
			list[i] += list[j-1] * list[i-j]
		}
	}

	return list[n]
}

func main() {
	fmt.Println(numTrees(2))
}
