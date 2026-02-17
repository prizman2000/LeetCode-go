package maxArea

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxLeft := height[left]
	maxRight := height[right]
	ma := minInt(height[left], height[right]) * (len(height) - 1)

	for left != right {
		if height[left] < height[right] {
			left++

			if height[left] <= maxLeft {
				continue
			}

			maxLeft = height[left]
		} else {
			right--

			if height[right] <= maxRight {
				continue
			}

			maxRight = height[right]
		}

		if minInt(height[left], height[right])*(right-left) > ma {
			ma = minInt(height[left], height[right]) * (right - left)
		}
	}

	return ma
}

func minInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}
