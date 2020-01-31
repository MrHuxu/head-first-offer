/*

题目: 输入一个正整数数组, 把数组里所有数字拼接起来排成一个数, 打印出能拼接出的所有数字中最小的一个.

解答: 对于两个数 m 和 n, 如果拼接之后 mn<nm, 那么最后的结果里, m 一定在 n 前面, 这样我们可以直接使用原生的 sort 函数来排序求解.

*/

package main

import "sort"

import "strconv"

// PrintMinNumber ...
func PrintMinNumber(numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return compare(
			strconv.Itoa(numbers[i])+strconv.Itoa(numbers[j]),
			strconv.Itoa(numbers[j])+strconv.Itoa(numbers[i]),
		)
	})

	for _, number := range numbers {
		print(number)
	}
	println("")
}

func compare(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str1[i] < str2[i] {
			return true
		}
	}

	return false
}

func main() {
	PrintMinNumber([]int{3, 32, 321})
}
