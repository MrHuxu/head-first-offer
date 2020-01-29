/*

题目: 实现一个函数用来匹配包含 `.` 和 `*` 的正则表达式.

解答: 两种做法, 一种是递归, 一种是 DP, 这里使用递归解法, 也可以使用二维数组进行动态规划求解.

*/

package main

func match(str string, pattern string) bool {
	switch {
	case str == "" && pattern == "":
		return true

	case pattern == "":
		return false

	case str == "":
		if len(pattern) >= 2 && pattern[1] == '*' {
			return match(str, pattern[2:len(pattern)])
		}
		return false

	default:
		switch pattern[len(pattern)-1] {
		case '.':
			return match(str[0:len(str)-1], pattern[0:len(pattern)-1])

		case '*':
			var idx int
			for idx = len(str) - 1; idx >= 0; idx-- {
				if str[idx] != pattern[len(pattern)-2] {
					break
				}
			}
			return match(str, pattern[0:len(pattern)-2]) || match(str[0:idx+1], pattern[0:len(pattern)-2])

		default:
			return match(str[0:len(str)-1], pattern[0:len(pattern)-1])
		}
	}
}

func main() {
	println(match("aaa", "a.a"))
	println(match("aaa", "ab*ac*a"))
	println(match("aaa", "aa.a"))
	println(match("aaa", "ab*a"))
	println(match("", "."))
	println(match("", "b*"))
}
