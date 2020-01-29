/*

题目: 长度给 length 的绳子, 把绳子剪成 m(m>1) 段, 没段长度为 k[0], k[1] ... k[m-1], 求 k[0]*k[1]*...*k[m-1] 可能的最大值.

解答:

两种解法:

1. DP
2. 每次都尝试用 3 来减, 这样的收益最大, 最后如果剩 4 的话, 就剪成两个 2 因为 2*2>3, 这样减就可以了.

*/

package main

import "math"

func maxProductAfterCuttingSolution1(length int) int {
	switch length {
	case 0, 1:
		return 0

	case 2:
		return 1

	case 3:
		return 2
	}

	dp := make([]int, length+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= length; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[j]*dp[i-j], dp[i])
		}
	}

	return dp[length]
}

func maxProductAfterCuttingSolution2(length int) int {
	switch length {
	case 0, 1:
		return 0

	case 2:
		return 1

	case 3:
		return 2
	}

	timesOf3 := length / 3

	if length-timesOf3*3 == 1 {
		timesOf3--
	}

	timesOf2 := (length - timesOf3*3) / 2

	return int(math.Pow(3, float64(timesOf3))) * int(math.Pow(2, float64(timesOf2)))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	println(maxProductAfterCuttingSolution1(8))
	println(maxProductAfterCuttingSolution2(8))
}
