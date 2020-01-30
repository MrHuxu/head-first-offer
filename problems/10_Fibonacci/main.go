/*

题目: 求斐波那契数列的第 n 项.

解答: 使用递归的话会有大量重复计算, 以及会有爆栈风险, 正确的做法应该是用两个临时变量一直往前迭代处理.

*/

package main

// Fibonacci ...
func Fibonacci(n int) int {
	num1 := 1
	num2 := 1

	if n == 1 || n == 2 {
		return num1
	}

	for i := 3; i <= n; i++ {
		tmp := num1 + num2
		num1 = num2
		num2 = tmp
	}

	return num2
}

func main() {
	println(Fibonacci(1))
	println(Fibonacci(2))
	println(Fibonacci(6))
	println(Fibonacci(11))
}
