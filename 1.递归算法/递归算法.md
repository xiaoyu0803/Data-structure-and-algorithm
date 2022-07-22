## 递归算法的基本思想
* 自己调用自己
## 递归基本要素
  * 基线条件
    > 跳出递归或程序结束的条件，
  * 缩小规模
    > 缩小每次递归的数据量或者长度
## 力扣递归算法举例
  1. [力扣231题-2的幂](https://leetcode.cn/problems/power-of-two/)
  > 给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。
  >  [递归解法 : ](./力扣231.go)
  ```
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
  ```
2.[力扣206题-反转链表](https://leetcode.cn/problems/reverse-linked-list/)
> 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
> ![](./反转链表.jpg)
```
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```
