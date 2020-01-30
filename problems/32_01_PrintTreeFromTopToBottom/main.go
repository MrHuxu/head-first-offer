/*

题目: 不分行地从上到下打印二叉树.

解答: 使用队列存储要遍历到节点, 逐个输出即可.

*/

package main

import "github.com/MrHuxu/types"

// PrintFromTopToBottom ...
func PrintFromTopToBottom(root *types.TreeNode) {
	if root == nil {
		return
	}

	queue := []*types.TreeNode{root}
	for len(queue) != 0 {
		head := queue[0]
		println(head.Val)

		queue = queue[1:len(queue)]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
}

func main() {
	PrintFromTopToBottom(types.BuildTree([]interface{}{8, 6, 10, 5, 7, 9, 11}))
}
