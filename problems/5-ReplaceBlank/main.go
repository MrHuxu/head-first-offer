/*

题目: 给定一个字符串和其字符数组的总长度, 把字符串当空格转换成 %20 三个字符

解答: 从前向后遍历, 需要提前安排各个字符的位置, 复杂度为 O(n^2), 可以换个思路, 从后向前遍历, 复杂度为 O(n)

*/

package main

// ReplaceBlank ...
func ReplaceBlank(str []byte, len int) {
	var originalLength int
	for i := 0; i < len; i++ {
		if str[i] != 0 {
			originalLength++
		} else {
			break
		}
	}

	for i := len - 1; i >= 0; i-- {
		if str[originalLength-1] == ' ' {
			str[i] = '0'
			i--
			str[i] = '2'
			i--
			str[i] = '%'
		} else {
			str[i] = str[originalLength-1]
		}

		originalLength--
	}
}

func main() {
	bytes := []byte{'i', ' ', 'a', 'm', 0, 0}
	ReplaceBlank(bytes, 6)
	println(string(bytes))
}
