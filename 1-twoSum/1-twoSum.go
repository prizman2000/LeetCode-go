package twoSum

func twoSum(nums []int, target int) []int {
    var hashMap = make(map[int]int, len(nums))

	for i, first := range nums {
		second := target - first

		if j, ok := hashMap[second]; ok {
			return []int{j, i}
		}

		if _, ok := hashMap[first]; !ok {
			hashMap[first] = i
		}
	}
	return nil
}