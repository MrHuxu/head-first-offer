/*

题目: 输入一棵二叉树的根节点, 判断这棵树是不是平衡二叉树.

解答: 递归处理.

*/

package main

import "github.com/MrHuxu/types"

// IsBalanced ...
func IsBalanced(root *types.TreeNode) bool {
	result, _ := isBalancedHelper(root)
	return result
}

func isBalancedHelper(root *types.TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftDepth := isBalancedHelper(root.Left)
	isRightBalanced, rightDepth := isBalancedHelper(root.Right)

	return isLeftBalanced && isRightBalanced && abs(leftDepth-rightDepth) <= 1, max(leftDepth, rightDepth) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(IsBalanced(types.BuildTree([]interface{}{1, 2, 3, 4, 5, nil, 6, nil, nil, 7})))
	println(IsBalanced(types.BuildTree([]interface{}{1, 2, 3, 4, 5, nil, 6, nil, nil, 7, nil, 8})))
}
