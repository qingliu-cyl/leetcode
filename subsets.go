package main

// 动态规划法， res[i] = res[0] +  res[i-1] + （res[0] +  res[i-1]）* nums[j]
func subsetsDP(nums []int) [][]int {
	res := make([][]int, 0)

	dp := func(n int) {
		for _, i := range res {
			newList := make([]int, len(i))
			copy(newList, i)
			res = append(res, append(newList, n))
		}
	}

	for _, num := range nums {
		// 本身肯定不重复
		dp(num)
		res = append(res, []int{num})
	}

	// 空集是任何集合的子集
	return append(res, []int{})
}

func subsetsDFS(nums []int) [][]int {
	res := make([][]int, 1)

	var dfs func(idx int, list []int)
	dfs = func(idx int, list []int) {
		newList := make([]int, len(list))
		copy(newList, list)
		newList = append(newList, nums[idx])
		res = append(res, newList)

		for i := idx + 1; i < len(nums); i++ {
			dfs(i, newList)
		}
	}

	for idx, _ := range nums {
		dfs(idx, []int{})
	}

	return res
}
