package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]int, 0)
	x := 0
	y := 0
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)
	for {
		if x < lenNums1 && y < lenNums2 {
			if nums1[x] < nums2[y] {
				res = append(res, nums1[x])
				x += 1
			} else {
				res = append(res, nums2[y])
				y += 1
			}
			continue
		}

		if x >= lenNums1 {
			res = append(res, nums2[y:]...)
			break
		} else {
			res = append(res, nums1[x:]...)
			break
		}
	}

	if len(res)%2 == 0 {
		return (float64(res[len(res)/2]) + float64(res[len(res)/2-1])) / 2
	} else {
		return float64(res[len(res)/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
