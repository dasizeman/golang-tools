package tools

// IntMax returns the larger of the two provided integers
func IntMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

// IntMin returns the smaller of the two provided integers
func IntMin(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
