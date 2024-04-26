package easy

// Given a two-digit integer n. Return the sum of its digits.
// For n = 29, the output should be solution(n) = 11.
func SumIntegerDigits(n int) int {
	result := 0
	if n != 0 {
		a := n / 10
		b := n % 10
		result = a + b
	}
	return result
}
