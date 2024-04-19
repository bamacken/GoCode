package medium

import (
	"math"
	"strconv"
	"strings"
)

// ReverseInteger reverses an integer.
func Reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}

	num, temp := 0, x
	if temp < 0 {
		temp *= -1
	}

	str, rsl := strings.Split(strconv.Itoa(x), "-"), ""
	n := len(str) - 1
	for i := 0; i < len(str[n]); i++ {
		rsl += strconv.Itoa(temp / (int(math.Pow10(i))) % 10)
	}
	num, _ = strconv.Atoi(rsl)

	if x < 0 {
		num *= -1
	}

	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}

	return num
}
