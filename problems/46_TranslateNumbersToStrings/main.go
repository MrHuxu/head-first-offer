/*

题目: 给定一个数字, 我们按照 0 翻译成 'a', 1 翻译成 'b' ... 25 翻译成 'z' 的方式, 总共有多少种翻译方法.

解答: 每一位数字, 既可以由当前位进行翻译, 也可以和前一位组合进行翻译, 就是一个简单的动态规划问题.

*/

package main

import (
	"strconv"
)

// GetTranslation ...
func GetTranslation(number int) int {
	return getTranslation(strconv.Itoa(number))
}

func getTranslation(str string) int {
	if len(str) == 0 {
		return 0
	}
	if len(str) == 1 {
		return 1
	}

	dp := make([]int, len(str))

	dp[0] = 1
	if (str[0]-'0')*10+(str[1]-'0') <= 25 {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(str); i++ {
		dp[i] = dp[i-1]

		if (str[i-1]-'0')*10+(str[i]-'0') <= 25 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(dp)-1]
}

func main() {
	println(GetTranslation(12258))
}
