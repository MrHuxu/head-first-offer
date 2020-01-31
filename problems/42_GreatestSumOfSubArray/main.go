/*

题目: 输入一个整型数组, 输出连续子数组的最大和.

解答:

动态规划, 声明数组 `dp []int` 使用 dp[i] 表示以第 i 个数字结尾的子数组的最大值

*/

package main

// FindGreatestSumOfSubArray ...
func FindGreatestSumOfSubArray(data []int) int {
	dp := make([]int, len(data))

	for i, d := range data {
		if i == 0 {
			dp[i] = d
		} else {
			if dp[i-1] <= 0 {
				dp[i] = d
			} else {
				dp[i] = dp[i-1] + d
			}
		}
	}

	return dp[len(dp)-1]
}

func main() {
	println(FindGreatestSumOfSubArray([]int{1, -2, 3, 10, -4, 7, 2, 5}))
}
