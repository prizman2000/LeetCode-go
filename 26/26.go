package removeDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	if nums == nil && len(nums) == 0 {
		return 0
	}

	uniqueCount := 1
	l := 0

	for i := 0; i < len(nums); i++ {
		if nums[l] != nums[i] {
			l += 1
			nums[l] = nums[i]
			uniqueCount++
		}
	}

	return uniqueCount
}
