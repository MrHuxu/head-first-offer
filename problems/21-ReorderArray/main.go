/*

题目: 输入一个整数数组, 实现一个函数来调整该数组中数字的顺序, 使得所有的奇数位于前半部分, 偶数位于后半部分.

解答: 使用双指针法, 比较好的实现是把判断奇偶的逻辑抽出来, 这样可以实现任意条件把数组分成两部分.

*/

package main

import "fmt"

// ReorderOddEven ...
func ReorderOddEven(data []int, check func(int) bool) {
	left := 0
	right := len(data) - 1
	for left < right {
		if check(data[left]) {
			left++
			continue
		}

		if !check(data[right]) {
			right--
			continue
		}

		tmp := data[left]
		data[left] = data[right]
		data[right] = tmp
	}
}

func isEven(n int) bool {
	return n%2 == 1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ReorderOddEven(arr, isEven)
	fmt.Println(arr)
}
