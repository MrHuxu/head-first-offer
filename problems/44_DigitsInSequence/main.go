/*

题目: 数字以 012345678910111213141516... 的格式序列化到一个字符串中, 求任意 n 位(从 0 开始计数)对应的数字.

解答: 如果一个一个数字数的话速度太慢, 可以按照 0~9, 10~99, 100~999 ... 这样的顺序一次找一批数字, 具体算法见代码.

*/

package main

import (
	"math"
	"strconv"
)

// digitAtIndex ...
func digitAtIndex(n int) int {
	digits := 1
	for n > digits*countOfIntegers(digits) {
		n -= digits * countOfIntegers(digits)
		digits++
	}

	multi := n / digits
	extra := n % digits

	if extra == 0 {
		return nthDigitOfNumber(beginNumber(digits)+multi, digits-1)
	}

	return nthDigitOfNumber(beginNumber(digits)+multi+1, extra)
}

func beginNumber(n int) int {
	if n == 1 {
		return 0
	}

	return int(math.Pow10(n - 1))
}

func countOfIntegers(n int) int {
	if n == 1 {
		return 10
	}

	return 9 * int(math.Pow10(n-1))
}

func nthDigitOfNumber(number, n int) int {
	return int(strconv.Itoa(number)[n] - '0')
}

func main() {
	println(digitAtIndex(5))
	println(digitAtIndex(1001))
}
