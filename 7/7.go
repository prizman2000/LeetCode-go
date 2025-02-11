package reverseInteger

func reverse(x int) int {
	var res = 0

	for radix := 1; x != 0; radix++ {
		if x%10 == 0 && res == 0 {
			x = x / 10
			continue
		}

		res *= 10
		res += x % 10

		if res > 214748364 || res < -214748364 {
			return 0
		}

		x = x / 10
	}

	return res
}
