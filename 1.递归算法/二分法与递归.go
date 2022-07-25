package main

import (
	"fmt"
	"log"
)

func main() {
	i := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(dichotomy(i, 2, 0, len(i)-1))
}
func dichotomy(i []int, n int, l, r int) int {

	mid := (r-l)/2 + l
	if l > r {
		return -1
	}
	log.Println(mid, "\n", i)
	if i[mid] == n {
		return mid
	} else if i[mid] > n {
		r = mid - 1
		return dichotomy(i, n, l, 90)
	} else if i[mid] < n {
		l = mid + 1
		return dichotomy(i, n, l, r)
	}
	return 0
}
