/*

题目: 输入一棵二叉树和一个整数, 打印出二叉树中节点值的和为输入整数的所有路径.

解答: 递归遍历, 记录当前的路径和路径所有值, 遇到叶子节点时, 就判断和是否为整数, 如果是的话就打印路径.

*/

package main

import "github.com/MrHuxu/types"

// FindPath ...
func FindPath(root *types.TreeNode, expectedSum int) {
	findPathRecursively(root, expectedSum, 0, []int{})
}

func findPathRecursively(root *types.TreeNode, expectedSum, currentSum int, path []int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && currentSum+root.Val == expectedSum {
		for _, item := range path {
			print(item, " ")
		}
		print(root.Val)
		println("")
	}

	findPathRecursively(root.Left, expectedSum, currentSum+root.Val, append([]int{}, append(path, root.Val)...))
	findPathRecursively(root.Right, expectedSum, currentSum+root.Val, append([]int{}, append(path, root.Val)...))
}

func main() {
	FindPath(types.BuildTree([]interface{}{10, 5, 12, 4, 7}), 22)
}
