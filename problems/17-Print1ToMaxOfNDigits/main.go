/*

题目: 输入数字 n, 打印出从 1 到最大的 n 位数.

解答:

复杂的做法是用字符串来模拟大数, 简单的做法就是, 其实 1 到 n 位最大数, 就是 n 个 0-9 的全排列, 去掉前面是 0 的结果.

*/

package main

import "fmt"

// Print1ToMaxOfNDigits ..
func Print1ToMaxOfNDigits(n int) {
	print1ToMaxOfNDigitsHelper(nil, n)
}

func print1ToMaxOfNDigitsHelper(bytes []byte, n int) {
	if len(bytes) == n {
		printNumber(bytes)
		return
	}

	for i := 0; i <= 9; i++ {
		tmp := append(bytes, '0'+byte(i))
		print1ToMaxOfNDigitsHelper(tmp, n)
	}
}

func printNumber(bytes []byte) {
	var hasNonZero bool
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '0' {
			if hasNonZero {
				fmt.Printf("%c", bytes[i])
			} else if i == len(bytes)-1 {
				fmt.Print("0")
			}
		} else {
			hasNonZero = true
			fmt.Printf("%c", bytes[i])
		}
	}
	fmt.Println()
}

func main() {
	Print1ToMaxOfNDigits(2)
	Print1ToMaxOfNDigits(3)
}
