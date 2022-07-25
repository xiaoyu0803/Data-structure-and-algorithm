package main

func isPowerOfTwo(n int) bool {
	if n == 2 || n == 1 {
		return true
	} else if n == 0 {
		return false
	} else if n%2 > 0 {
		return false
	}
	return isPowerOfTwo(n / 2)
}
