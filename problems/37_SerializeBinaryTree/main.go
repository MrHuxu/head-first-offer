/*

题目: 实现两个函数, 分别用来序列化和反序列化二叉树.

解答: 主要是看怎么处理树里的空位, LeetCode 上采用的是层序遍历, 使用 null 表示空位, 下面这两种是通过前序遍历, 使用 -999 表示空位的方式.

*/

package main

import (
	"fmt"

	"github.com/MrHuxu/types"
)

type data struct {
	list []int
}

func serialize(root *types.TreeNode) *data {
	if root == nil {
		return &data{[]int{-999}}
	}

	result := &data{[]int{root.Val}}

	left := serialize(root.Left)
	right := serialize(root.Right)

	result.list = append(result.list, append(left.list, right.list...)...)
	return result
}

func deserialize(d *data, result *types.TreeNode) {
	result.Val = d.list[0]
	d.list = d.list[1:len(d.list)]

	if d.list[0] != -999 {
		result.Left = new(types.TreeNode)
		deserialize(d, result.Left)
	} else {
		d.list = d.list[1:len(d.list)]
	}

	if d.list[0] != -999 {
		result.Right = new(types.TreeNode)
		deserialize(d, result.Right)
	} else {
		d.list = d.list[1:len(d.list)]
	}
}

func main() {
	tree := types.BuildTree([]interface{}{1, 2, 3, 4, nil, 5, 6})

	serializeResult := serialize(tree)
	fmt.Println(serializeResult)
	deserializeResult := new(types.TreeNode)
	deserialize(serializeResult, deserializeResult)
	fmt.Println(deserializeResult)
}
