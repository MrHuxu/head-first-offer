/*

题目: 判断在一个矩阵中是否存在一条包含字符串中所有字符的路径.

解答: 回溯法, 每次向前面四个方向的其中一个方向前进, 如果不符合就会到上一个节点更换另一个方向遍历, 最后就可以拿到结果.

*/

package main

func hasPath(matrix [][]byte, str string) bool {
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if hasPathHelper(matrix, i, j, 0, str, visited) {
				return true
			}
		}
	}

	return false
}

func hasPathHelper(matrix [][]byte, i, j, k int, str string, visited [][]bool) bool {
	if k == len(str) {
		return true
	}

	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) || visited[i][j] {
		return false
	}

	if matrix[i][j] != str[k] {
		return false
	}

	visited[i][j] = true
	result := hasPathHelper(matrix, i-1, j, k+1, str, visited) ||
		hasPathHelper(matrix, i+1, j, k+1, str, visited) ||
		hasPathHelper(matrix, i, j-1, k+1, str, visited) ||
		hasPathHelper(matrix, i, j+1, k+1, str, visited)
	visited[i][j] = false

	return result
}

func main() {
	println(hasPath([][]byte{
		[]byte{'a', 'b', 't', 'g'},
		[]byte{'c', 'f', 'c', 's'},
		[]byte{'j', 'd', 'e', 'h'},
	}, "bfce"))

	println(hasPath([][]byte{
		[]byte{'a', 'b', 't', 'g'},
		[]byte{'c', 'f', 'c', 's'},
		[]byte{'j', 'd', 'e', 'h'},
	}, "abfb"))
}
