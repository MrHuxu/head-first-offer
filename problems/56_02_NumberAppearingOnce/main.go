/*

题目: 数组中除了一个数字只出现一次, 别的都出现了三次, 找出这个只出现一次的数字.

解答: 对数组中数字的每个位进行累加, 如果是 3 的倍数, 那么只出现一次的数字这位是 0, 否则就是 1.

*/

package main

// FindNumberAppearingOnce ...
func FindNumberAppearingOnce(numbers []int) int {
	var result int

	var cnt int
	for {
		var hasLarger bool

		var sum int
		for _, number := range numbers {
			if number>>cnt != 0 {
				hasLarger = true

				sum += (number >> cnt) & 1
			}
		}

		result |= sum % 3 << cnt
		cnt++

		if !hasLarger {
			break
		}
	}

	return result
}

func main() {
	println(FindNumberAppearingOnce([]int{2, 1, 2, 2}))
	println(FindNumberAppearingOnce([]int{2, 1, 2, 2, 4, 4, 4, 1, 1, 3, 5, 6, 5, 5, 6, 6}))
}
