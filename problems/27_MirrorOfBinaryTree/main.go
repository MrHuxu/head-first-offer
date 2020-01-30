/*

题目: 完成一个函数, 输入一棵二叉树, 输出它的镜像.

解答: 递归可解.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

// MirrorRecursively ...
func MirrorRecursively(root *types.TreeNode) {
	if root == nil {
		return
	}

	MirrorRecursively(root.Left)
	MirrorRecursively(root.Right)

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
}

func main() {
	tree := types.BuildTree([]interface{}{8, 6, 10, 5, 7, 9, 11})
	fmt.Println(tree)
	MirrorRecursively(tree)
	fmt.Println(tree)
}
