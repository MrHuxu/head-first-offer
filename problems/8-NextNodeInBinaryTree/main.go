/*

题目: 输入一个二叉树节点, 返回其在中序遍历中的下一个节点.

解答:

1. 如果这个节点右子树不为空, 那么它的下一个节点就是右子树中的最左节点.
2. 如果这个节点右子树为空, 但是它是父节点的左子节点, 那么它的下一个节点就是它的父节点.
3. 如果这个节点右子树为空, 且为父节点的右子节点, 那么就沿着父节点指针一直向上遍历, 直到找到一个节点是其父节点的左子节点, 那么下一个遍历到节点就是这个节点的父节点.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/leetcode150/problems/utils"
)

// GetNext ...
func GetNext(node *utils.TreeNode) *utils.TreeNode {
	if node == nil {
		return nil
	}

	var next *utils.TreeNode

	if node.Right != nil {
		tmp := node.Right

		for tmp.Left != nil {
			tmp = tmp.Left
		}

		next = tmp
	} else if node.Parent != nil {
		current := node
		parent := node.Parent

		for parent != nil && current == parent.Right {
			current = parent
			parent = parent.Parent
		}

		next = parent
	}

	return next
}

func main() {
	nodeA := &utils.TreeNode{Val: 1}
	nodeB := &utils.TreeNode{Val: 2}
	nodeC := &utils.TreeNode{Val: 3}
	nodeD := &utils.TreeNode{Val: 4}
	nodeE := &utils.TreeNode{Val: 5}
	nodeF := &utils.TreeNode{Val: 6}
	nodeG := &utils.TreeNode{Val: 7}
	nodeH := &utils.TreeNode{Val: 8}
	nodeI := &utils.TreeNode{Val: 9}

	nodeA.Left = nodeB
	nodeA.Right = nodeC
	nodeB.Parent = nodeA
	nodeB.Left = nodeD
	nodeB.Right = nodeE
	nodeC.Parent = nodeA
	nodeC.Left = nodeF
	nodeC.Right = nodeG
	nodeD.Parent = nodeB
	nodeE.Parent = nodeB
	nodeE.Left = nodeH
	nodeE.Right = nodeI
	nodeF.Parent = nodeC
	nodeG.Parent = nodeC
	nodeH.Parent = nodeE
	nodeI.Parent = nodeE

	fmt.Println(GetNext(nodeE).Val)
	fmt.Println(GetNext(nodeG))
}
