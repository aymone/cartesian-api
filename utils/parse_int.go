package utils

import "strconv"

// ParseInt is a util method to parse a string number p to int
// and assure that p is not empty string returning a bool value
// for validation purposes, so string zero "0" return true and n=0
func ParseInt(p string) (bool, int) {
	var n int
	if p == "" {
		return false, n
	}

	n, err := strconv.Atoi(p)
	if err != nil {
		return false, n
	}

	return true, n
}
