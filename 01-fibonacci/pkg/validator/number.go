package validator

// Compare value in range from min to max
func InRange(n int, min int, max int) bool {
	return min <= n && n <= max
}
