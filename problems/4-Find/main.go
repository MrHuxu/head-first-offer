/*

题目: 在一个二维数组中, 每一行都按照从左到右递增的顺序排序, 每一列都按照从上到下递增的顺序排序. 请完成一个函数, 输入这样的一个二维数组和一个整数, 判断数组中是否含有该整数.

解答:

1. 通过二分查找找到目标数字在的行, 再通过二分查找找到目标在一行中的位置.

2. 我们从这个二维数组的右上角开始判断:

		- 如果当前位置的值大于 target, 则当前列都大于 target, target 只可能在当前列的左侧, 对列数 -1;
		- 如果当前位置的值小于 target, 则当前行都小于 target, target 只可能在当前行的下面, 对行数 +1;
		- 如果当前位置的值等于 target, 返回 true.

	 同理也可以从左下角开始判断.

*/

package main

// Find ...
func Find(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix) - 1
	for row <= len(matrix)-1 && col >= 0 {
		switch {
		case matrix[row][col] > target:
			col--

		case matrix[row][col] < target:
			row++

		default:
			return true
		}
	}

	return false
}

func main() {
	println(Find([][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}, 7))

	println(Find([][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}, 3))
}
