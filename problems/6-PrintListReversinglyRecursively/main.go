/*

题目: 输入一个链表的头结点, 从尾到头打印出链表每个节点的值.

解答: 递归打印当前节点除开头节点后面的部分, 再打印头节点的值, 结果就是从尾到头输出的.

*/

package main

import (
	"github.com/MrHuxu/leetcode150/problems/utils"
)

// PrintListReversinglyRecursively ...
func PrintListReversinglyRecursively(pHead *utils.ListNode) {
	if pHead == nil {
		return
	}

	PrintListReversinglyRecursively(pHead.Next)
	println(pHead.Val)
}

func main() {
	list := utils.BuildList([]int{1, 34, 4, 2, 12})
	PrintListReversinglyRecursively(list)
}
