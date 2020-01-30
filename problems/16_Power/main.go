/*

题目: 实现一个 Power 函数, 求 base 的 exponent 次方.

解答:

exponent 为 0 和 1 可以提前返回, 并且需要主要 exponent 为负的情况.

并且求 n 次方可以转换成两个 n/2 次方的乘积, 把 O(n) 的累乘操作, 优化成 O(logn).

*/

package main

// Power ...
func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}

	if exponent == 1 {
		return base
	}

	absExponent := uint(exponent)
	if exponent < 0 {
		absExponent = uint(-exponent)
	}

	result := powerHelper(base, absExponent)

	if exponent < 0 {
		return 1 / result
	}

	return result
}

func powerHelper(base float64, exponent uint) float64 {
	if exponent == 0 {
		return 1
	}

	if exponent == 1 {
		return base
	}

	result := powerHelper(base, exponent>>1)
	result *= result

	if exponent%2 == 1 {
		result *= base
	}

	return result
}

func main() {
	println(Power(2, 3))
	println(Power(2, -2))
}
