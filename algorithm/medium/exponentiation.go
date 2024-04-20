package medium

func MyPow(x float64, n int) float64 {
	if x == 1 || n == 0 {
		return 1
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	results := 1.0
	for n > 0 {
		if n%2 == 1 {
			results *= x
		}
		x *= x
		n /= 2
	}
	return results
}
