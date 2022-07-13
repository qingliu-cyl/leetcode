package main

func isNumber(s string) bool {
	status := []map[string]int{
		{" ":0, "s": 1, "d":2, ".":4},
		{"d":2, ".":4},
		{"d":2, ".":3, "e":5, " ":8},
		{"d":3, "e":5, " ": 8},
		{"d":3},
		{"s":6, "d":7},
		{"d":7},
		{"d":7, " ":8},
		{" ":8},
	}

	p := 0
	for idx, _ := range s {
		t := "?"
		if []byte("0")[0] <= s[idx] && s[idx] <= []byte("9")[0]{
			t = "d"
		} else if string(s[idx]) == "+" || string(s[idx]) == "-" {
			t = "s"
		} else if string(s[idx]) == "e" || string(s[idx]) == "E" {
			t = "e"
		} else if string(s[idx]) == "." || string(s[idx]) == " "{
			t = string(s[idx])
		}

		find := func() bool {
			for k, _ := range status[p] {
				if k == t {
					return true
				}
			}
			return false
		}

		if !find() {
			return false
		}
		p = status[p][t]
	}

	return p == 2 || p == 3 || p == 7 || p == 8
}