/*

题目: 假设把某股票的价格按照实践先后顺序存储在数组中, 请问买卖一次可能获得的最大利润是多少.

解答: 使用一个变量存储遍历位置之前到最小值, 然后和当前值对比, 找出最大差值.

*/

package main

// MaxDiff ...
func MaxDiff(numbers []int) int {
	var result int

	prevMin := numbers[0]
	for _, number := range numbers {
		if number > prevMin {
			if number-prevMin > result {
				result = number - prevMin
			}
		} else {
			prevMin = number
		}
	}

	return result
}

func main() {
	println(MaxDiff([]int{9, 11, 8, 5, 7, 12, 16, 14}))
}
