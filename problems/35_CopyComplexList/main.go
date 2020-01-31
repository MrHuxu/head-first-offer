/*

题目: 复制一个复杂链表, 在复杂链表中, 每个节点除了有一个 next 指针指向下一个节点, 还有一个 sibling 指针指向链表中的任意节点或 nil.

解答:

这道题的一个难点就是怎么通过原来的节点 n 的 sibling 指针获得复制之后节点 n` 的 sibling` 指针, 一个做法是:

1. 建立一个原链表节点和新链表节点的 map, 我们可以通过 sibling 指针为 key 找到 sibling` 指针指向的节点, 需要额外 O(n) 的时间复杂度;
2. 复制的时候, 直接把新节点插入到原节点的后面, 这样 sibling` 指针直线的节点可以直接通过 sibling.next 找到, 最后把链表里的偶数节点抽出拼成一个新链表返回即可.

*/

package main

import "fmt"

import "strconv"

type listNode struct {
	val int

	next, sibling *listNode
}

func (l *listNode) String() (str string) {
	if l == nil {
		return "<nil>"
	}

	str += strconv.Itoa(l.val)
	for l.next != nil {
		str += "->" + strconv.Itoa(l.next.val)
		l = l.next
	}
	return
}

func cloneNodes(list *listNode) *listNode {
	if list == nil {
		return nil
	}

	tmp := list
	for tmp != nil {
		newNode := &listNode{val: tmp.val}
		newNode.next = tmp.next
		tmp.next = newNode

		tmp = tmp.next.next
	}

	tmp = list
	for tmp != nil {
		if tmp.sibling != nil {
			tmp.next.sibling = tmp.sibling.next
		}

		tmp = tmp.next.next
	}

	dummyHead := new(listNode)
	tmp = list.next
	tmp2 := dummyHead
	for tmp != nil {
		tmp2.next = tmp
		tmp2 = tmp2.next

		if tmp.next == nil {
			break
		}
		tmp = tmp.next.next
	}

	return dummyHead.next
}

func main() {
	node1 := &listNode{val: 1}
	node2 := &listNode{val: 2}
	node3 := &listNode{val: 3}
	node4 := &listNode{val: 4}
	node5 := &listNode{val: 5}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	node1.sibling = node3
	node2.sibling = node5
	node4.sibling = node2

	list := cloneNodes(node1)
	fmt.Println(list)
	fmt.Println(list.sibling)
	fmt.Println(list.next.sibling)
	fmt.Println(list.next.next.next.sibling)
}
