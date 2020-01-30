/*

题目: 输入两个递增排序的链表, 合并这两个链表并使新链表中的节点仍然是递增排序的.

解答: 简单的链表遍历.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// Merge ...
func Merge(head1, head2 *types.ListNode) *types.ListNode {
	dummyHead := new(types.ListNode)
	tail := dummyHead

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tail.Next = head1
			head1 = head1.Next
		} else {
			tail.Next = head2
			head2 = head2.Next
		}

		tail = tail.Next
	}

	if head1 != nil {
		tail.Next = head1
	} else {
		tail.Next = head2
	}

	return dummyHead.Next
}

func main() {
	fmt.Println(Merge(
		types.BuildList([]int{1, 3, 5, 8}),
		types.BuildList([]int{2, 4, 6, 7}),
	))
}
