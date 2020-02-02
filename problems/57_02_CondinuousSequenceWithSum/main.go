/*

题目: 输入一个正整数 s, 打印出所有和为 s 的连续正数序列(至少有两个数).

解答: 滑动窗口.

*/

package main

// FindContinuousSequence ...
func FindContinuousSequence(target int) {
	front := 1
	sum := 0
	for tail := 1; tail <= target/2+1; tail++ {
		sum += tail

		if sum == target {
			for i := front; i <= tail; i++ {
				print(i, " ")
			}
			println("")
		} else if sum > target {
			for sum > target {
				sum -= front
				front++
			}

			if sum == target {
				for i := front; i <= tail; i++ {
					print(i, " ")
				}
				println("")
			}
		}

	}
}

func main() {
	FindContinuousSequence(15)
}
