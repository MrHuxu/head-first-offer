/*

题目: 给定一棵二叉搜索树, 请找出其中第 k 大调节点.

解答: 二叉搜索树的中序遍历是一个有序数组, 中序遍历找到第 k 个即可.

*/

package main

import "github.com/MrHuxu/types"

// KthNode ...
func KthNode(root *types.TreeNode, k int) *types.TreeNode {
	var cnt int

	var result *types.TreeNode
	var traverse func(*types.TreeNode)
	traverse = func(node *types.TreeNode) {
		if node == nil || result != nil {
			return
		}

		traverse(node.Left)

		cnt++
		if cnt == k {
			result = node
		}

		traverse(node.Right)
	}
	traverse(root)

	return result
}

func main() {
	println(KthNode(types.BuildTree([]interface{}{5, 3, 7, 2, 4, 6, 8}), 3).Val)
}
