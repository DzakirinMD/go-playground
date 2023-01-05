package main

// Add @return sum of 2 number
func Add(x, y int) (res int) {
	return x + y
}

// Subtract @return difference of 2 number
func Subtract(x, y int) (res int) {
	return x - y
}

// IntMin @ return a if a less than b
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
