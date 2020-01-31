/*

题目: 输入一个字符串, 打印出该字符串中字符的所有排列.

解答:

可以把字符串看成两部分, 即第一个字符和后面的字符串, 整个解答可以看做两步:

1. 求出第一个字符后面字符串的全排列;
2. 把第一个字符逐个和后面的字符串交换, 并且求全排列.

*/

package main

// Permutation ...
func Permutation(str string) {
	if str == "" {
		return
	}

	permutation([]byte(str), 0)
}

func permutation(str []byte, begin int) {
	if begin == len(str) {
		println(string(str))
	}

	for i := begin; i < len(str); i++ {
		tmp := str[i]
		str[i] = str[begin]
		str[begin] = tmp

		permutation(str, begin+1)

		tmp = str[i]
		str[i] = str[begin]
		str[begin] = tmp
	}
}

func main() {
	Permutation("abc")
}
