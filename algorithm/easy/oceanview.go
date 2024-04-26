package easy

import "math"

func FindBuildings(heights []int) []int {
	h := len(heights)
	if h < int(-math.Pow(10, 5)) || h > int(math.Pow(10, 5)) {
		return nil
	}

	result := make([]int, 0, h)

	maxHeight := 0
	for i := h - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append(result, i)
			maxHeight = heights[i]
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
