package main

func prefixCount(words []string, pref string) int {
	res := 0
	prefLen := len(pref)
	for _, str := range words {
		if len(str) < prefLen {
			continue
		}


		if str[:prefLen] == pref {
			res++
		}
	}

	return res
}