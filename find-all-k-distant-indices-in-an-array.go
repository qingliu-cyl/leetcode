package main

func findKDistantIndices(nums []int, key int, k int) []int {
	kList := make([]int, 0)
	res := make([]int, 0)
	for idx, v := range nums {
		if v == key {
			kList = append(kList, idx)
		}
	}
	abs := func(v int) int {
		if v >= 0 {
			return v
		} else {
			return -1 * v
		}
	}
	judge := func(v int) {
		for _, kIdx := range kList {
			if abs(v-kIdx) <= k {
				res = append(res, v)
				return
			}
		}
	}

	for idx, _ := range nums {
		judge(idx)
	}

	return res
}

func main() {

}
