package main

import (
	"fmt"
)

/*
整型 链表 的简单实现
*/
type link struct {
	Val  int
	Next *link
}

//读取链表元素到切片
func (l *link) Read() []int {
	var linkList []int
	for {
		if l.Next == nil {
			linkList = append(linkList, l.Val)
			break
		}
		linkList = append(linkList, l.Val)
		l = l.Next
	}
	return linkList
}

// 向尾部添加元素
func (l *link) AddTail(val ...int) {
	for {
		if l.Next == nil {
			for _, v := range val {
				l.Next = &link{
					Val:  v,
					Next: nil,
				}
				l = l.Next
			}
			break
		}
		l = l.Next
	}
}

// 删除中间元素
func (l *link) delete(index int) error {

	for {
		if l.Next == nil && index != 0 {
			return fmt.Errorf("index invalid")
		}
		index--
		fmt.Println(l.Val)
		if index == 0 && l.Next != nil {
			l.Next = l.Next.Next
			break
		} else if index == 0 && l.Next == nil {
			l = &link{}
			break
		} else if index == -1 && l.Next != nil {

			a := &link{0, l.Next}
			*l = *a
			break
		}
		l = l.Next

	}
	return nil
}

func main() {
	var a = new(link)
	a.Val = 8
	a.AddTail(1, 2, 3)
	fmt.Print(a.Read())
	a.delete(0)

	fmt.Println(a.Read())
}
