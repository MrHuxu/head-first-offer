/*

题目: 地上有一个 rows*cols 的方格, 机器人从 (0,0) 开始可以向四个方向走, 其中如果两个坐标的数位和大于 threshold, 那么这个格不能到达, 求这个机器人能取到多少格.

解答: 和 12 题类似, 还是回溯法.

*/

package main

func movingCount(threshold, rows, cols int) int {
	visited := make([][]bool, rows)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, cols)
	}

	var result int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result += movingCountHelper(threshold, i, j, rows, cols, visited)
		}
	}

	return result
}

func movingCountHelper(threshold, i, j, rows, cols int, visited [][]bool) int {
	if !validate(i, j, threshold) {
		return 0
	}

	if i < 0 || i >= rows || j < 0 || j >= cols || visited[i][j] {
		return 0
	}

	visited[i][j] = true
	return movingCountHelper(threshold, i-1, j, rows, cols, visited) +
		movingCountHelper(threshold, i+1, j, rows, cols, visited) +
		movingCountHelper(threshold, i, j-1, rows, cols, visited) +
		movingCountHelper(threshold, i, j+1, rows, cols, visited) + 1
}

func validate(threshold, i, j int) bool {
	var sum int

	for i != 0 {
		sum += i % 10
		if sum > threshold {
			return false
		}
		i /= 10
	}

	for j != 0 {
		sum += j % 10
		if sum > threshold {
			return false
		}
		j /= 10
	}

	return true
}

func main() {
	println(movingCount(3, 10, 10))
}
