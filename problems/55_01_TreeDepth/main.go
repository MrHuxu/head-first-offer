/*

题目: 二叉树的深度.

解答: 递归遍历.

*/

package main

import "github.com/MrHuxu/types"

// TreeDepth ...
func TreeDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(TreeDepth(root.Left), TreeDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(TreeDepth(types.BuildTree([]interface{}{1, 2, 3, 4, 5, nil, 6, nil, nil, 7})))
}
