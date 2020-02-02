/*

题目: 数组中除了两个数字仅出现一次, 别的数字都出现了两次, 找出这两个数字.

解答:

可以尝试把数组分成两个部分, 让第一个部分的异或的结果为第一个数字, 第二个部分的异或的结果为第二个数字.

我们先把数字整体做异或, 结果肯定不为 0, 我们结果中选取一个为 1 的位, 这两个数字在这一位上肯定一个是 1 一个是 0, 然后根据这一位把数组分成两部分, 那么这一位位 1 的一部分异或的结果就是这位为 1 的那个数, 剩下一部分异或结果就是另一个数.

*/

package main

import "fmt"

// FindNumsAppearOnce ...
func FindNumsAppearOnce(data []int) (int, int) {
	var num int
	for _, d := range data {
		num ^= d
	}

	flag := 1
	for num&flag == 0 {
		flag = flag << 1
	}

	var arr1, arr2 []int
	for _, n := range data {
		if n&flag != 0 {
			arr1 = append(arr1, n)
		} else {
			arr2 = append(arr2, n)
		}
	}

	var num1, num2 int
	for _, n := range arr1 {
		num1 ^= n
	}
	for _, n := range arr2 {
		num2 ^= n
	}

	return num1, num2
}

func main() {
	fmt.Println(FindNumsAppearOnce([]int{2, 4, 3, 6, 3, 2, 5, 5}))
}
