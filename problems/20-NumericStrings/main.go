/*

题目: 实现一个函数用来判断字符串是否表示数值, 包括整数和小数, "+100", "5e2", "-123", "3.1416" 及 "-1E-16" 都是合法数值.

解答:

合法的字符串符合这两种模式, A.[.[B]][e|EC] 和 .B[e|EC], 其中:

A, C: 有符号整数
B: 无符号整数

按这样进行匹配即可.

*/

package main

func isNumeric(str string) bool {
	if str == "" {
		return false
	}

	var idx int
	var numeric, rightPart bool

	if idx, numeric = scanInteger(str, 0); numeric && idx == len(str) {
		return true
	}

	if str[idx] == '.' {
		idx, rightPart = scanUnsignedInteger(str, idx+1)
		numeric = rightPart || numeric
	} else if str[idx] == 'e' || str[idx] == 'E' {
		idx, rightPart = scanInteger(str, idx+1)
		numeric = rightPart && numeric
	}

	return numeric && idx == len(str)
}

func scanInteger(str string, idx int) (int, bool) {
	if idx == len(str) {
		return idx + 1, false
	}

	if str[idx] == '+' || str[idx] == '-' {
		return scanUnsignedInteger(str, idx+1)
	}
	return scanUnsignedInteger(str, idx)
}

func scanUnsignedInteger(str string, idx int) (int, bool) {
	if idx == len(str) {
		return idx + 1, false
	}

	newIdx := idx
	for newIdx < len(str) && str[newIdx] >= '0' && str[newIdx] <= '9' {
		newIdx++
	}

	return newIdx, newIdx > idx
}

func main() {
	println(isNumeric("+100"))
	println(isNumeric("5e2"))
	println(isNumeric("-123"))
	println(isNumeric("3.1416"))
	println(isNumeric("-1E-16"))

	println(isNumeric("12e"))
	println(isNumeric("1a3.14"))
	println(isNumeric("1.2.3"))
	println(isNumeric("+-5"))
	println(isNumeric("12e+5.4"))
}
