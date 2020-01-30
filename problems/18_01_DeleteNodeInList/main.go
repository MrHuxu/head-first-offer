/*

题目: 输入一个链表的头指针和一个节点指针, 定义一个函数在 O(1) 的时间内删除该节点.

解答:

常规的删除肯定要遍历到这个节点, 复杂度为 O(n), 其实只需要把当前节点后一个节点的值赋给要删除的节点, 再把后一个节点的 next 指针赋给要删除的节点即可.

删除节点为尾节点的情况需要单独遍历处理, 此时复杂度为 O(n), 但是平均复杂度仍然为 O(1).

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// DeleteNode ...
func DeleteNode(head *types.ListNode, toBeDeleted *types.ListNode) {
	if toBeDeleted.Next == nil {
		for head.Next.Next != nil {
			head = head.Next
		}
		head.Next = nil
		return
	}

	toBeDeleted.Val = toBeDeleted.Next.Val
	toBeDeleted.Next = toBeDeleted.Next.Next
}

func main() {
	list := types.BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println(list)
	node3 := list.Next.Next
	fmt.Println(node3.Val)
	DeleteNode(list, node3)
	fmt.Println(list)

	node1 := list
	fmt.Println(node1.Val)
	DeleteNode(list, node1)
	fmt.Println(list)

	node5 := list.Next.Next
	fmt.Println(node5.Val)
	DeleteNode(list, node5)
	fmt.Println(list)
}
