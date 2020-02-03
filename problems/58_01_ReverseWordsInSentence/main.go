/*

题目: 输入一个英文句子, 反转句子中单词的顺序, 但单词内字符的顺序不变.

解答: 第一步翻转整个句子, 第二步翻转句子中每个单词.

*/

package main

// ReverseSentence ...
func ReverseSentence(sentence string) string {
	sentence = reverse(sentence)

	var result string

	start := -1
	for i, r := range sentence {
		if r != ' ' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				result += reverse(sentence[start:i])
				start = -1
			}

			result += sentence[i : i+1]
		}
	}
	if start != -1 {
		result += reverse(sentence[start:len(sentence)])
	}

	return result
}

func reverse(str string) string {
	if str == "" {
		return ""
	}

	return str[len(str)-1:len(str)] + reverse(str[0:len(str)-1])
}

func main() {
	println(ReverseSentence("I am a student."))
}
