package maths

// abs returns the absolute value of an integer.
func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
