/*

题目: 字符串当左旋转操作是把字符串前面的若干个字符转移到字符串当尾部.

解答: 可以复用上一题的 reverse 函数, 旋转 n 位可以转换成, 先把字符串整体翻转, 再把最后 n 位和前面的字符串分别翻转.

*/

package main

// LeftRotateString ...
func LeftRotateString(str string, n int) string {
	str = reverse(str)
	return reverse(str[0:len(str)-n]) + reverse(str[len(str)-n:len(str)])
}

func reverse(str string) string {
	if str == "" {
		return ""
	}

	return str[len(str)-1:len(str)] + reverse(str[0:len(str)-1])
}

func main() {
	println(LeftRotateString("abcdefg", 2))
}
