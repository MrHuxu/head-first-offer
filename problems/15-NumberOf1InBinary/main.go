/*

题目: 求一个十进制数转换成二进制后 1 的个数.

解答:

两个解法:

1. 直接把数字左移然后按位与 1, 这样的时间复杂度为 O(n), n 为二进制数字的尾数;
2. 有一个规律就是, n&(n-1) 得到的结果会把 n 最低的 1 变成 0, 只要重复这个操作直到 n 变为 0 即可, 这样的时间复杂度为 O(n), n 为二进制数字中 1 的个数.

*/

package main

// NumberOf1Solution1 ...
func NumberOf1Solution1(n int) int {
	var result int

	for n != 0 {
		if n&1 == 1 {
			result++
		}
		n = n >> 1
	}

	return result
}

// NumberOf1Solution2 ...
func NumberOf1Solution2(n int) int {
	var result int

	for n != 0 {
		n = n & (n - 1)
		result++
	}

	return result
}

func main() {
	println(NumberOf1Solution1(9))
	println(NumberOf1Solution2(9))

	println(NumberOf1Solution1(62))
	println(NumberOf1Solution2(62))

	println(NumberOf1Solution1(89))
	println(NumberOf1Solution2(89))
}
