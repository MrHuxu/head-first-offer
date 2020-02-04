/*

题目: 求 1 累加到 n, 要求不能使用乘除法, for, while, if, else, switch, case 等关键字.

解答: 我们可以使用递归求解, 但是普通的递归需要用到判断语句来结束, 我们可以转换一下思路, 也就是递归结束和没结束其实是调用两个不同的函数, 我们可以使用一个 map 来分别调用这两个函数, 达到判断递归结束的目的.

*/

package main

var m = make(map[bool]func(int) int)

// Sum ...
func Sum(n int) int {
	m[true] = terminate
	m[false] = processing

	return processing(n)
}

func processing(n int) int {
	return n + m[n-1 == 0](n-1)
}

func terminate(n int) int {
	return 0
}

func main() {
	println(Sum(4))
}
