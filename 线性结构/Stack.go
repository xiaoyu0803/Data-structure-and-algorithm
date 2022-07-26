package main

import (
	"fmt"
)

/*
 使用 slice 实现栈的操作
*/
type stacks []int

// 出栈操作
func (s *stacks) pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stacks is empty")
	}
	defer func() {
		*s = (*s)[:len(*s)-1]
	}()
	return (*s)[len(*s)-1], nil
}

//入栈操作
func (s *stacks) push(elem ...int) {
	*s = append(*s, elem...)
}

func StackeTest() {
	var s stacks
	s.push(1, 2, 3, 4, 5)
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
}
