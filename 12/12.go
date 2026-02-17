package intToRoman

import (
	"strings"
)

func intToRoman(num int) string {
	var res string

	idx := 1
	p := map[int][]string{
		1: {"I", "V", "X"},
		2: {"X", "L", "C"},
		3: {"C", "D", "M"},
	}

	for ; num > 0; num = num / 10 {
		last := num % 10

		switch last {
		case 9:
			res = p[idx][0] + p[idx][2] + res
		case 4:
			res = p[idx][0] + p[idx][1] + res
		case 0, 1, 2, 3:
			res = strings.Repeat(p[idx][0], last) + res
		case 5, 6, 7, 8:
			res = p[idx][1] + strings.Repeat(p[idx][0], last-5) + res
		}

		if idx < 3 {
			idx++
		} else {
			res = strings.Repeat(p[idx][2], num/10) + res
			break
		}
	}

	return res
}
