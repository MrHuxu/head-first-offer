/*

题目: 输入一棵二叉搜索树, 将该二叉树搜索树转换成一个排序的双向链表.

解答: 二叉搜索树的中序遍历就是一个有序链表, 这里我们需要使用两个指针来表示双向链表的头节点和尾节点.

1. 根据二叉搜索树的特性, 头节点直接从根节点一直向左遍历就可以得到;
2. 每次中序遍历把二叉树的节点挂到链表尾节点上.

最后返回链表头节点即可.

*/

package main

import "github.com/MrHuxu/types"

// Convert ...
func Convert(root *types.TreeNode) *types.TreeNode {
	head := root
	for head != nil && head.Left != nil {
		head = head.Left
	}

	var tail *types.TreeNode
	var traverse func(*types.TreeNode)
	traverse = func(node *types.TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		node.Left = tail
		if tail != nil {
			tail.Right = node
		}
		tail = node
		traverse(node.Right)
	}
	traverse(root)

	return head
}

func main() {
	head := Convert(types.BuildTree([]interface{}{10, 6, 14, 4, 8, 12, 16}))
	for head != nil {
		println(head.Val)
		head = head.Right
	}
}
