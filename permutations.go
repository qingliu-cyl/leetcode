package main

func permute(nums []int) [][]int {
	res := [][]int{}

	index := func(target int, list []int) bool {
		for _, i := range list {
			if i == target {
				return true
			}
		}

		return false
	}

	var dfs func(list []int)
	dfs = func(list []int) {
		if len(list) == len(nums) {
			newList := make([]int, len(list))
			copy(newList, list)
			res = append(res, list)
		}

		for _, i := range nums {
			if index(i, list) {
				continue
			}

			dfs(append(list, i))
		}
	}

	for _, i := range nums {
		dfs([]int{i})
	}

	return res
}

func main() {

}
