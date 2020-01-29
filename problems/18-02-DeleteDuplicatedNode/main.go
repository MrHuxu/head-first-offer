/*

题目: 删除链表中重复的节点.

解答: 很麻烦的链表遍历题.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// DeleteDuplication ...
func DeleteDuplication(head *types.ListNode) {
	dummyHead := &types.ListNode{Next: head}

	pre := dummyHead
	cur := dummyHead.Next
	for cur != nil {
		tmp := cur.Next
		cnt := 1
		for tmp != nil {
			if tmp.Val == cur.Val {
				tmp = tmp.Next
				cnt++
			} else {
				break
			}
		}

		if cnt < 2 {
			pre = pre.Next
			cur = pre.Next
			continue
		}

		pre.Next = tmp
		cur = pre.Next

		if pre == nil {
			break
		}
	}
}

func main() {
	list := types.BuildList([]int{1, 2, 3, 3, 4, 4, 5, 6, 6, 6})
	fmt.Println(list)
	DeleteDuplication(list)
	fmt.Println(list)
}
