/*

题目: 找出数组中出现次数超过一半的数字.

解答: 我们使用一个变量表示数组中的一个数字, 一个变量表示次数, 当我们遍历到下一个数字时

1. 如果和之前保存的数字不同, 则次数减 1;
2. 相同则数字加 1;
3. 如果次数为 0, 则把数字遍历为当前遍历到的数字, 因为有一个数字的个数超过一半, 所以最后会更新成这个数字.

*/

package main

// MoreThanHalfNum ...
func MoreThanHalfNum(numbers []int) int {
	result := numbers[0]

	var times int
	for _, number := range numbers {
		if times == 0 {
			result = number
			times = 1
		} else {
			if number == result {
				times++
			} else {
				times--
			}
		}
	}

	return result
}

func main() {
	println(MoreThanHalfNum([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
