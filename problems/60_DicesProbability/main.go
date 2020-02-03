/*

题目: 把 n 个骰子扔在地上, 所有骰子朝上一面的点数之和为 s, 输入 n, 打印出 s 的所有可能的值出现的概率.

解答: 还是使用筛法类似的做法, 不是暴力统计次数, 而是每次增加一个骰子可能出现的数字的次数, 最后统计计算概率.

*/

package main

import "fmt"

// PrintProbability ...
func PrintProbability(number int) {
	values := []int{0, 1, 1, 1, 1, 1, 1}
	for i := 2; i <= number; i++ {
		nextValues := make([]int, i*6+1)
		for j := (i - 1) * 1; j <= (i-1)*6; j++ {
			for k := 1; k <= 6; k++ {
				nextValues[k+j] += values[j]
			}
		}

		values = nextValues
	}

	var sum int
	for i := number; i <= number*6; i++ {
		sum += values[i]
	}
	for i := number; i <= number*6; i++ {
		fmt.Printf("%f\n", float64(values[i])/float64(sum))
	}
}

func main() {
	PrintProbability(1)
	PrintProbability(2)
	PrintProbability(3)
	PrintProbability(6)
}
