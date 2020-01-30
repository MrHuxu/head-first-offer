/*

题目: 实现一个函数, 判断一棵二叉树是不是对称的.

解答: 递归可解.

*/

package main

import "github.com/MrHuxu/types"

func isSymmetrical(root *types.TreeNode) bool {
	return isSymmetricalRecursively(root, root)
}

func isSymmetricalRecursively(left, right *types.TreeNode) bool {
	switch {
	case left == nil && right == nil:
		return true

	case left == nil, right == nil:
		return false

	default:
		return left.Val == right.Val &&
			isSymmetricalRecursively(left.Left, right.Right) &&
			isSymmetricalRecursively(left.Right, right.Left)
	}
}

func main() {
	println(isSymmetrical(types.BuildTree([]interface{}{8, 6, 6, 5, 7, 7, 5})))
	println(isSymmetrical(types.BuildTree([]interface{}{8, 6, 5, 7, 7})))
}
