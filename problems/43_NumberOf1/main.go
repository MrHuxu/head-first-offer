/*

题目: 输入一个整数 n, 求 1~n 这 n 个数字中的十进制表示中 1 出现的次数.

解答: 不要一个个算, 而是直接算出一批数字中出现的 1 的个数, 具体解法直接看书 P222 吧.

*/

package main

import "strconv"

import "math"

// NumbersOf1Between1AndN ...
func NumbersOf1Between1AndN(n int) int {
	return numberOf1(strconv.Itoa(n))
}

func numberOf1(strN string) int {
	if strN == "" {
		return 0
	}

	first := strN[0] - '0'

	if len(strN) == 1 && first == 0 {
		return 0
	}
	if len(strN) == 1 && first > 0 {
		return 1
	}

	var numFirstDigit int
	if first > 1 {
		numFirstDigit += int(math.Pow10(len(strN) - 1))
	} else {
		tmp, _ := strconv.Atoi(strN[1:len(strN)])
		numFirstDigit += tmp + 1
	}

	numOtherDigits := int(first) * (len(strN) - 1) * int(math.Pow10(len(strN)-2))

	numRecursive := numberOf1(strN[1:len(strN)])

	return numFirstDigit + numOtherDigits + numRecursive
}

func main() {
	println(NumbersOf1Between1AndN(21345))
	println(NumbersOf1Between1AndN(55))
}
