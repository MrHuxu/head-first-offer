/*

题目: 0~n-1 这 n 个数字排成一个圆圈, 从数字 0 开始, 每次从这个圆圈里删除第 m 个数字, 求出这个圆圈里最后一个数字.

解答: 最糙的解法就是用一个环来模拟删除过程, 但是其实有数学解法:

1. n == 1 时, 有 f(n,m)=0;
2. n > 1 时, 有 f(n,m)=[f(n-1,m)+m]%n.

*/

package main

// LastRemaining ...
func LastRemaining(n, m int) int {
	dp := make([]int, n+1)

	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + m) % i
	}

	return dp[n]
}

func main() {
	println(LastRemaining(5, 3))
}
