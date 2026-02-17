package stringToIntegerAtoi

var mapper = map[byte]int{
	48: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
}

func myAtoi(s string) int {
	var res, numStarts int
	sign := 1

	for numStarts = 0; numStarts < len(s); numStarts++ {
		if _, isDigit := mapper[s[numStarts]]; isDigit {
			break
		}

		if s[numStarts] == 32 {
			continue
		}

		if s[numStarts] == 45 {
			sign = -1
			numStarts++
			break
		}

		if s[numStarts] == 43 {
			numStarts++
			break
		}

		if s[numStarts] == 48 {
			continue
		}

		return 0
	}

	for i := numStarts; i < len(s); i++ {
		if s[i] == 48 && res == 0 {
			continue
		}
		if digit, isDigit := mapper[s[i]]; isDigit {
			res = res*10 + digit
			if 2147483648 <= res {
				if sign == -1 {
					return -2147483648
				} else {
					return 2147483647
				}
			}

		} else {
			break
		}
	}

	return sign * res
}
