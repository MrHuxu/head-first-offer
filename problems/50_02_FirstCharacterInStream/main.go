/*

题目: 字符流中第一个只出现一次的字母.

解答:

思想还是和上题类似, 使用一个 occurrence 数组记录下所有字符的位置(假设 256 可以放下所有的 ASCII 字符).

每次从流中读进一个字符, 就更新 occurrence 对应的值:

1. 如果原本的值为 0, 那么就更新为当前下标 +1;
2. 如果原本的值不为 0, 那么这个字符已经出现过了, 就更新为 -1.

最后找出 occurrence 数组中所有值大于 0 的最小值对应的字符即可.

*/

package main

import "fmt"

// FirstApperingOnce ...
func FirstApperingOnce(str string) byte {
	occurrence := make([]int, 256)

	for i := range str {
		if occurrence[str[i]] > 0 {
			occurrence[str[i]] = -1
		} else {
			occurrence[str[i]] = i + 1
		}
	}

	var ch byte
	var idx int
	for i, o := range occurrence {
		if o > 0 {
			if idx == 0 || o < idx {
				idx = o
				ch = byte(i)
			}
		}
	}

	return ch
}

func main() {
	fmt.Printf("%c\n", FirstApperingOnce("go"))
	fmt.Printf("%c\n", FirstApperingOnce("google"))
}
