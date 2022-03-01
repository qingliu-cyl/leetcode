package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 字母异位字符串， 排序之后一定是相同的
	resMap := make(map[string][]string)
	for _, str := range strs {
		newStrSlice := []byte(str)
		sort.Slice(newStrSlice, func(i, j int) bool {
			return newStrSlice[i] < newStrSlice[j]
		})
		newStr := string(newStrSlice)
		resMap[newStr] = append(resMap[newStr], str)
	}
	res := make([][]string, 0)
	for _, list := range resMap {
		res = append(res, list)
	}
	return res
}
