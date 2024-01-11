package medianOfTwoSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int
	len1 := len(nums1)
	len2 := len(nums2)

	merged := make([]int, len1 + len2)

	for i < len1 || j < len2 {
		if j >= len2 || (i < len1 && nums1[i] <= nums2[j]) {
			merged[i + j] = nums1[i]
			i++
		} else {
			merged[i + j] = nums2[j]
			j++
		}
	}

	if len(merged) % 2 == 1 {
		return float64(merged[(len(merged) - 1) / 2])
	} else {
		return float64(merged[len(merged) / 2] + merged[(len(merged) - 1) / 2]) / 2
	}
}