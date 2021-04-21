package validator

// Compare value in range from min to max
func InRange(value int, min int, max int) bool {
	if value < min || value > max {
		return false
	}
	return true
}
