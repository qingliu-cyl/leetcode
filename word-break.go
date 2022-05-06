package main

func wordBreak(s string, wordDict []string) bool {
	resMap := make(map[string]bool)
	for _, word := range wordDict {
		resMap[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && resMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func main() {

}
