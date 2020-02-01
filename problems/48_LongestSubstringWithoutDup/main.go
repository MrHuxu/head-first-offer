/*

题目: 请从字符串中找出一个最长的不包含重复字符的子字符串.

解答: 滑动窗口法.

*/

package main

func longestSubstringWithoutDuplication(str string) int {
	if str == "" {
		return 0
	}

	result := 1

	var front, tail, length int
	window := make(map[byte]int)
	for tail < len(str) {
		if window[str[tail]] == 0 {
			window[str[tail]]++
			tail++
			length++

			if length > result {
				result = length
			}
		} else {
			for front < tail {
				window[str[front]]--
				front++
				length--

				if window[str[front]] == 1 {
					break
				}
			}
		}
	}

	return result
}

func main() {
	println(longestSubstringWithoutDuplication("arabcacfr"))
}
