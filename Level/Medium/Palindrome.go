package medium

// LongestPalindrome finds the longest palindrome in a string using the expand around center approach.
func LongestPalindrome(s string) string {
	results := ""
	for i := 0; i < len(s); i++ {
		odd := ExpandOroundCenter(s, i, i)
		if len(odd) > len(results) {
			results = odd
		}
		even := ExpandOroundCenter(s, i, i+1)
		if len(even) > len(results) {
			results = even
		}
	}
	return results
}

func ExpandOroundCenter(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// LongestPalindromeBruteForce finds the longest palindrome in a string using brute force.
func LongestPalindromeBruteForce(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	curr := ""
	max := 0
	results := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			curr = string(s[i:j])
			str := reverseString(curr)
			if curr == str {
				if len(curr) > max {
					max = len(curr)
					results = curr
				}
			}
		}
	}
	return results
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
