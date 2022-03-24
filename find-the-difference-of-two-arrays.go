package main

//找出两数组的不同
func findDifference(nums1 []int, nums2 []int) [][]int {
	key1 := make(map[int]struct{})
	key2 := make(map[int]struct{})

	res1 := make([]int, 0)
	res2 := make([]int, 0)

	for _, i := range nums2 {
		key2[i] = struct{}{}
	}

	for _, i := range nums1 {
		key1[i] = struct{}{}
	}

	for k, _ := range key2 {
		if _, ok := key1[k]; !ok {
			res2 = append(res2, k)
		}
	}

	for k, _ := range key1 {
		if _, ok := key2[k]; !ok {
			res1 = append(res1, k)
		}
	}

	res := [][]int{res1, res2}
	return res
}

func main() {

}
