/*

题目: 在一个 m*n 的棋盘上每一格都放有一个礼物, 每个礼物都有一个价值, 从左上角每次只能向右或向下移动, 直到达到棋盘右下角, 计算最多能拿到多少价值的礼物.

解答: DP 可解.

*/

package main

func getMaxValue(values [][]int) int {
	dp := make([][]int, len(values))
	for i := range dp {
		dp[i] = make([]int, len(values[i]))
	}

	for i, row := range values {
		for j, col := range row {
			var top int
			if i > 0 {
				top = dp[i-1][j]
			}
			var left int
			if j > 0 {
				left = dp[i][j-1]
			}

			dp[i][j] = col + max(top, left)
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	println(getMaxValue([][]int{
		[]int{1, 10, 3, 8},
		[]int{12, 2, 9, 6},
		[]int{5, 7, 4, 11},
		[]int{3, 7, 16, 5},
	}))
}
