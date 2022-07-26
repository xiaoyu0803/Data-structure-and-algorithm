package main

import "fmt"

/*
用 slice 模拟 queue
*/
type queue []int

// 队列添加元素
func (q *queue) push(elem ...int) {
	*q = append(*q, elem...)
}

// 取出队列元素
func (q *queue) poll() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	defer func() {
		*q = (*q)[1:]
	}()
	return (*q)[0], nil
}

// 判断长度
func (q *queue) length() int {
	return len(*q)
}

func queueTest() {
	var a queue
	a.push(1, 2, 3)
	fmt.Println(a)
	a.poll()
	fmt.Println(a.poll())
	fmt.Println(a.length())
}
