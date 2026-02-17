package romanToInt

func romanToInt(s string) int {
	var res int

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var currentLetter byte = 'I'

	for i := len(s) - 1; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[currentLetter] {
			res -= romanMap[s[i]]
		} else {
			res += romanMap[s[i]]
			currentLetter = s[i]
		}
	}

	return res
}
