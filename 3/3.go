package lengthOfLongestSubstring

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
    hMap := make(map[uint8]int)
    var left, right, res int

    for right < len(s) {
        hMap[s[right]] += 1
		
        for hMap[s[right]] > 1 {
            hMap[s[left]] -= 1
            left += 1
        }
        res = int(math.Max(float64(res), float64(right - left + 1)))
		right += 1
    }
    return res
}