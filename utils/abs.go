package utils

// Abs returns the absulute number for integer
// Golang standard lib has math.Abs but it receives a float64 number as input.
// The implementation is simple, so i prefered this way instead of
// converting int to float64 and return as int again
func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
