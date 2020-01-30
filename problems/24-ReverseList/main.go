/*

题目: 输入一个链表的头结点, 反转该链表并输出反转后端头节点.

解答: 简单的链表遍历.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// ReverseList ...
func ReverseList(head *types.ListNode) *types.ListNode {
	dummyHead := new(types.ListNode)

	for head != nil {
		tmp := head.Next

		head.Next = dummyHead.Next
		dummyHead.Next = head

		head = tmp
	}

	return dummyHead.Next
}

func main() {
	fmt.Println(ReverseList(types.BuildList([]int{1, 2, 3, 5, 4})))
}
