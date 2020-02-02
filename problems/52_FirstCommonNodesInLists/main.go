/*

题目: 输入两个链表, 找出他们的第一个公共节点.

解答:

两种方法:

1. 因为两个有公共节点的链表一定是一个 Y 型不是 X 型, 从公共链表之后都是一样的, 可以建立两个栈, 把两个链表的节点分别压栈, 然后一起弹栈, 如果弹出来的节点是一样的, 就更新公共节点, 最后返回即可, 这样需要额外空间.
2. 快慢指针法, 两个链表往后遍历得到长度, 然后长度长一点的链表先前进相差的长度, 然后一起遍历, 相遇的节点就是公共节点.

下面是第二种解法.

*/

package main

import "github.com/MrHuxu/types"

import "fmt"

// FindFirstCommonNode ...
func FindFirstCommonNode(head1, head2 *types.ListNode) *types.ListNode {
	var l1, l2 int

	tmp1 := head1
	for tmp1 != nil {
		l1++
		tmp1 = tmp1.Next
	}

	tmp2 := head2
	for tmp2 != nil {
		l2++
		tmp2 = tmp2.Next
	}

	if l1 > l2 {
		for l1 != l2 {
			head1 = head1.Next
			l1--
		}
	} else if l2 > l1 {
		for l2 != l1 {
			head2 = head2.Next
			l2--
		}
	}

	for head1 != head2 {
		head1 = head1.Next
		head2 = head2.Next
	}

	return head1
}

func main() {
	list1 := types.BuildList([]int{1, 2, 3})
	list2 := types.BuildList([]int{1, 2, 3, 4, 5})
	list3 := types.BuildList([]int{6, 7, 8})
	list1.Next.Next.Next = list3
	list2.Next.Next.Next.Next.Next = list3

	fmt.Println(FindFirstCommonNode(list1, list2))
}
