/*

题目: 输入某二叉树的前序遍历和中序遍历的结果, 重建该二叉树.

解答: 前序遍历结果第一项给根节点的值, 通过这个值可以把中序遍历的结果分为左右子树中序遍历的结果, 再分别对左右子树进行递归构建即可.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/leetcode150/problems/utils"
)

// Construct ...
func Construct(preorder, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var rootPos int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootPos = i
			break
		}
	}

	root := &utils.TreeNode{Val: preorder[0]}
	root.Left = Construct(preorder[1:rootPos+1], inorder[0:rootPos])
	root.Right = Construct(preorder[rootPos+1:len(preorder)], inorder[rootPos+1:len(inorder)])

	return root
}

func main() {
	fmt.Println(Construct([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}))
}
