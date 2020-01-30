/*

题目: 输入两棵二叉树 A 和 B, 判断 B 是不是 A 的子结构.

解答: 递归判断, 注意根节点和非根节点处理的不同.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// HasSubtree ...
func HasSubtree(root1, root2 *types.TreeNode) bool {
	return hasSubtreeHelper(root1, root2, true)
}

func hasSubtreeHelper(root1, root2 *types.TreeNode, isRoot bool) bool {
	switch {
	case root1 == nil && root2 == nil:
		return true

	case root2 == nil:
		return true

	case root1 == nil:
		return false

	default:
		has := root1.Val == root2.Val &&
			hasSubtreeHelper(root1.Left, root2.Left, false) &&
			hasSubtreeHelper(root1.Right, root2.Right, false)

		if isRoot {
			has = has ||
				hasSubtreeHelper(root1.Left, root2, true) ||
				hasSubtreeHelper(root1.Right, root2, true)
		}

		return has
	}
}

func main() {
	tree1 := types.BuildTree([]interface{}{8, 8, 7, 9, 2, nil, nil, nil, nil, 4, 7})
	fmt.Println(tree1)
	tree2 := types.BuildTree([]interface{}{8, 9, 2})
	fmt.Println(tree2)
	tree3 := types.BuildTree([]interface{}{2, nil, 7})
	fmt.Println(tree3)
	tree4 := types.BuildTree([]interface{}{8, 7, 7})
	fmt.Println(tree4)

	println(HasSubtree(tree1, tree2))
	println(HasSubtree(tree1, tree3))
	println(HasSubtree(tree1, tree4))
}
