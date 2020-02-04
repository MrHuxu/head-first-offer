/*

题目: 从扑克牌中随机抽 5 张牌, 判断是不是一个顺子. 2~5 为数字本身, A 为 1, J 为 11, Q 为 12, K 为 13, 大小王可以看做任意数字.

解答:

可以对数字使用快排进行递增排序, 那么为顺子的条件为, 0 的个数大于等于空缺的个数, 且没有不为 0 的重复牌, 下面使用这种解法.

但是因为数字的离散的规模比较小, 其实可以建立一个大小为 14 的数组, 直接通过这个数组来保存数字个数做排序.

*/

package main

import "sort"

// IsContinuous ...
func IsContinuous(numbers []int) bool {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	var numberOf0 int
	for i, number := range numbers {
		if number == 0 {
			numberOf0++
			continue
		}

		if i == 0 || numbers[i-1] == 0 || i == 4 {
			continue
		}

		if number != numbers[i-1]+1 {
			numberOf0 -= number - numbers[i-1] - 1
		}
	}

	return numberOf0 >= 0
}

func main() {
	println(IsContinuous([]int{1, 2, 3, 4, 5}))
	println(IsContinuous([]int{0, 0, 3, 5, 7}))
	println(IsContinuous([]int{0, 0, 1, 5, 7}))
}
