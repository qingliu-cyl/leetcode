package main

import "fmt"

var res [][]int

//candidates 原始数组
//idx 待计算数组元素的索引
//list 可能的匹配结果
//sum  当前list的和
//target 目标值
func dfs(candidates []int, idx int, list []int, sum int, target int) {
	if idx == len(candidates) {
		return
	}
	if sum+candidates[idx] == target {
		newList := make([]int, len(list))
		copy(newList, list)
		res = append(res, append(newList, candidates[idx]))
		return
	} else if sum+candidates[idx] > target {
		return
	} else {
		for i := idx; i < len(candidates); i++ {
			dfs(candidates, i, append(list, candidates[idx]), sum+candidates[idx], target)
		}
	}

}

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	for i := 0; i < len(candidates); i++ {
		dfs(candidates, i, []int{}, 0, target)
	}

	return res
}

func main() {
	fmt.Println(combinationSum([]int{4, 8}, 32))
}
