package threeSum

type sortedNums struct {
	a int
	b int
	c int
}

func threeSum(nums []int) [][]int {
	mapDuples := map[sortedNums]struct{}{}

	mapNums := make(map[int]int, len(nums))

	for _, num := range nums {
		mapNums[num]++
	}

	if mapNums[0] >= 3 {
		mapDuples[sortedNums{
			a: 0,
			b: 0,
			c: 0,
		}] = struct{}{}
	}

	for num, count := range mapNums {
		if count > 1 {
			if _, ok := mapNums[(num+num)*-1]; num != (num+num)*-1 && ok {
				mapDuples[sortThreeNums(num, num, (num+num)*-1)] = struct{}{}
			}
		}
		for num2, _ := range mapNums {
			if num != num2 && (num+num2)*-1 != num && (num+num2)*-1 != num2 {
				if _, ok := mapNums[(num+num2)*-1]; ok {
					mapDuples[sortThreeNums(num, num2, (num+num2)*-1)] = struct{}{}
				}
			}
		}
	}

	res := [][]int{}
	for n, _ := range mapDuples {
		res = append(res, []int{n.a, n.b, n.c})
	}

	return res
}

func sortThreeNums(a, b, c int) sortedNums {
	if a >= b && a >= c {
		if b >= c {
			return sortedNums{
				a: a,
				b: b,
				c: c,
			}
		}

		return sortedNums{
			a: a,
			b: c,
			c: b,
		}
	}

	if b >= a && b >= c {
		if a >= c {
			return sortedNums{
				a: b,
				b: a,
				c: c,
			}
		}

		return sortedNums{
			a: b,
			b: c,
			c: a,
		}
	}

	if c >= a && c >= b {
		if a >= b {
			return sortedNums{
				a: c,
				b: a,
				c: b,
			}
		}

		return sortedNums{
			a: c,
			b: b,
			c: a,
		}
	}

	return sortedNums{}
}
