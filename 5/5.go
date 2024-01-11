package longestPalindromicSubstr

import "strings"

func longestPalindrome(s string) string {
	transformed := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(transformed)
	p := make([]int, n)
	center, right := 0, 0
	
	for i := 1; i < n - 1; i++ {
		if right > 1 {
			p[i] = min(right - i, p[2 * center - i])
		}
		for transformed[i + 1 + p[i]] == transformed[i - 1 - p[i]] {
			p[i]++
		}
		if i + p[i] > right {
			center, right = i, i + p[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i, v := range p {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}
	return s[(centerIndex - maxLen) / 2 : (centerIndex + maxLen) / 2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}