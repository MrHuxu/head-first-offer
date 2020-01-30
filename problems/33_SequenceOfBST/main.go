/*

题目: 输入一个整数数组, 判断该数组是不是某二叉搜索树的后续遍历结果.

解答: 二叉搜索树后续遍历的根节点在最末位置, 通过根节点可以把数组分成左子树和右子树遍历结果的两个部分, 再对两个部分递归判断即可.

*/

package main

// VerifySequenceOfBST ...
func VerifySequenceOfBST(sequence []int) bool {
	if len(sequence) == 0 || len(sequence) == 1 {
		return true
	}

	var div int
	for div = len(sequence) - 2; div >= 0; div-- {
		if sequence[div] < sequence[len(sequence)-1] {
			break
		}
	}

	for i := div; i >= 0; i-- {
		if sequence[i] >= sequence[len(sequence)-1] {
			return false
		}
	}

	return VerifySequenceOfBST(sequence[0:div+1]) && VerifySequenceOfBST(sequence[div+1:len(sequence)-1])
}

func main() {
	println(VerifySequenceOfBST([]int{5, 7, 6, 9, 11, 10, 8}))
	println(VerifySequenceOfBST([]int{7, 6, 4, 5}))
}
