/*

题目: 如果一个链表中包含环, 实现一个函数找到环的入口节点.

解答:

1. 首先使用快慢指针法, 快指针一次遍历前进两个, 慢指针一次前进一个, 找到环里面的一个节点;
2. 再使用这个节点获得环的长度 length;
3. 再使用快慢指针法, 快指针比慢指针先出发 length 个位置, 同时遍历, 相遇的位置就是环的入口.

*/

package main

import "github.com/MrHuxu/types"

// MeetingNode ...
func MeetingNode(head *types.ListNode) *types.ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next
	for fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}

	nodeInList := fast
	listLen := 1
	tmp := nodeInList.Next
	for tmp != nodeInList {
		listLen++
		tmp = tmp.Next
	}

	fast = head
	slow = head
	for i := 0; i < listLen; i++ {
		slow = slow.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func main() {
	list := types.BuildList([]int{1, 2, 3, 4, 5})
	node5 := list.Next.Next.Next.Next
	node3 := list.Next.Next
	node5.Next = node3
	println(list.Next.Next.Next.Next.Val, list.Next.Next.Next.Next.Next.Val)

	println(MeetingNode(list).Val)
}
