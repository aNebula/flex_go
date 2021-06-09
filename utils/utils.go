package utils

/**
Function to return the absolute value of an input int.
*/
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/**
Function to return the smallest integer out of 2 input integers.
*/
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
