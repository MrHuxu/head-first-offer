/*

题目: 输入一个链表, 输出该链表中倒数第 k 个节点.

解答: 双指针法, 需要处理链表长度小于 k 的情况.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// FindKthToTail ...
func FindKthToTail(head *types.ListNode, k uint) *types.ListNode {
	if head == nil || k == 0 {
		return nil
	}

	node1 := head
	node2 := head

	for i := 0; i < int(k); i++ {
		if node2 == nil {
			return nil
		}
		node2 = node2.Next
	}

	for node2 != nil {
		node1 = node1.Next
		node2 = node2.Next
	}

	return node1
}

func main() {
	list := types.BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println(FindKthToTail(list, 3))
	fmt.Println(FindKthToTail(list, 4))
	fmt.Println(FindKthToTail(list, 1))
	fmt.Println(FindKthToTail(list, 5))

	fmt.Println(FindKthToTail(list, 6))
	fmt.Println(FindKthToTail(list, 0))
}
