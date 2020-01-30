/*

题目: 输入一个矩阵, 按照从外向里以顺时针的顺序打印出每一个数字.

解答: 每圈分为上下和左右四个方向循环打印即可.

*/

package main

// PrintMatrixClockwisely ...
func PrintMatrixClockwisely(numbers [][]int) {
	for i := 0; i < (len(numbers)+1)/2; i++ {
		printCircle(numbers, i)
	}
}

func printCircle(numbers [][]int, level int) {
	for i := level; i < len(numbers[level])-level; i++ {
		println(numbers[level][i])
	}

	for i := level + 1; i < len(numbers)-level-1; i++ {
		println(numbers[i][len(numbers[level])-1-level])
	}

	for i := len(numbers[level]) - 1 - level; i >= level; i-- {
		println(numbers[len(numbers)-1-level][i])
	}

	for i := len(numbers) - level - 2; i >= level+1; i-- {
		println(numbers[i][level])
	}
}

func main() {
	PrintMatrixClockwisely([][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
		[]int{13, 14, 15, 16},
		[]int{13, 14, 15, 16},
	})
}
