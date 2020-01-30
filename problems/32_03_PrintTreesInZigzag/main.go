/*

题目: 之字形打印二叉树

解答: 增加一个布尔值来表示遍历方向, 同样逐层遍历即可.

*/

package main

import "github.com/MrHuxu/types"

// Print ...
func Print(root *types.TreeNode) {
	if root == nil {
		return
	}

	level := []*types.TreeNode{root}
	fromLeft := true
	for len(level) != 0 {
		var nextLevel []*types.TreeNode

		if fromLeft {
			for i := 0; i < len(level); i++ {
				print(level[i].Val, " ")

				if level[i].Left != nil {
					nextLevel = append(nextLevel, level[i].Left)
				}
				if level[i].Right != nil {
					nextLevel = append(nextLevel, level[i].Right)
				}
			}
		} else {
			for i := len(level) - 1; i >= 0; i-- {
				print(level[i].Val, " ")

				if level[i].Right != nil {
					nextLevel = append([]*types.TreeNode{level[i].Right}, nextLevel...)
				}
				if level[i].Left != nil {
					nextLevel = append([]*types.TreeNode{level[i].Left}, nextLevel...)
				}
			}
		}
		println("")

		level = nextLevel
		fromLeft = !fromLeft
	}
}

func main() {
	Print(types.BuildTree([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
}
