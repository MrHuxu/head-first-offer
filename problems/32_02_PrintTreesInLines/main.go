/*

题目: 分行从上到下打印二叉树.

解答: 使用一个数组不停记录每一行的节点数并打印出值, 当一行的节点数为 0 时结束循环返回即可.

*/

package main

import "github.com/MrHuxu/types"

// Print ...
func Print(root *types.TreeNode) {
	if root == nil {
		return
	}

	level := []*types.TreeNode{root}
	for len(level) != 0 {
		var nextLevel []*types.TreeNode
		for _, node := range level {
			print(node.Val, " ")

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		println("")

		level = nextLevel
	}
}

func main() {
	Print(types.BuildTree([]interface{}{8, 6, 10, 5, 7, 9, 11}))
}
