package autocomplete

import "unicode/utf8"

const minLengthThreshold = 32

func ComputeDistance(a, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}
	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}

	if a == b {
		return 0
	}

	s1 := []rune(a)
	s2 := []rune(b)

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	n := len(s1)
	m := len(s2)

	dp := make([]int, n+3)

	for i := 1; i < len(dp); i++ {
		dp[i] = i
	}

	_ = dp[n]
	for i := 1; i <= m; i++ {
		prev := i
		for j := 1; j <= n; j++ {
			curr := dp[j-1] // if s1[j-1] == s2[i-1]
			if s2[i-1] != s1[j-1] {
				insertCost := dp[j-1] + 1
				replaceCost := prev + 1
				deleteCost := dp[j] + 1
				curr = min(min(insertCost, deleteCost), replaceCost)
			}
			dp[j-1] = prev
			prev = curr
		}
		dp[n] = prev
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
