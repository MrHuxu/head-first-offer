/*

题目: 找出字符串中第一个只出现一次的字符.

解答: 使用一个哈希表存储每个字符出现的次数, 经过两次遍历就可以获得第一个只有一次的字符, 空间换时间, 空间复杂度其实也就是 26 个字符的固定开销, 所以空间复杂度为 O(1), 时间复杂度为 O(n)j.

*/

package main

import "fmt"

// FirstNotRepeatingChar ...
func FirstNotRepeatingChar(str string) byte {
	m := make(map[byte]int)
	for i := range str {
		m[str[i]]++
	}
	for i := range str {
		if m[str[i]] == 1 {
			return str[i]
		}
	}
	return ' '
}

func main() {
	fmt.Printf("%c\n", FirstNotRepeatingChar("abaccdeff"))
}
