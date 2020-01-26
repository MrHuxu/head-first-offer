/*

题目: 输入一个链表的头结点, 从尾到头打印出链表每个节点的值.

解答: 递归打印当前节点除开头节点后面的部分, 再打印头节点的值, 结果就是从尾到头输出的.

*/

package main

import (
	"github.com/MrHuxu/types"
)

// PrintListReversinglyRecursively ...
func PrintListReversinglyRecursively(head *types.ListNode) {
	if head == nil {
		return
	}

	PrintListReversinglyRecursively(head.Next)
	println(head.Val)
}

func main() {
	list := types.BuildList([]int{1, 34, 4, 2, 12})
	PrintListReversinglyRecursively(list)
}
