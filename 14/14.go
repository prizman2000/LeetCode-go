package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[0]); j++ {
			if j >= len(strs[i]) {
				strs[0] = strs[0][:j]
				break
			}

			if strs[0][j] != strs[i][j] {
				strs[0] = strs[0][:j]
			}
		}
	}

	return strs[0]
}
