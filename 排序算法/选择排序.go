package main

import "fmt"

func main() {
	SelectSort()
}
func SelectSort() {
	var arr = []int{1, 3, 7, 6, 5, 4, 2, 89, 10, 9}
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}
