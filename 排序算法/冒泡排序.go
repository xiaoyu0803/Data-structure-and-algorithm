package main

import "fmt"

func Bubble() {
	var bigen = []int{1, 3, 7, 5, 8, 4, 9, 2}
	for i := 0; i < len(bigen)-1; i++ {
		for j := 0; j < len(bigen)-1-i; j++ {
			if bigen[j] > bigen[j+1] {
				bigen[j], bigen[j+1] = bigen[j+1], bigen[j]
			}
		}
	}
	fmt.Println(bigen)
}
