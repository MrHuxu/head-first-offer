/*

题目: 找出一个递增数组的数组里任意一对和为 s 的两个数字.

解答: 双指针逼近法.

*/

package main

import "fmt"

// FindNumbersWithSum ...
func FindNumbersWithSum(data []int, sum int) (int, int) {
	front := 0
	tail := len(data) - 1
	for front < tail {
		if data[front]+data[tail] > sum {
			tail--
		} else if data[front]+data[tail] < sum {
			front++
		} else {
			return data[front], data[tail]
		}
	}

	return data[front], data[tail]
}

func main() {
	fmt.Println(FindNumbersWithSum([]int{1, 2, 4, 7, 11, 15}, 15))
}
