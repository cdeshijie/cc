/*package leetcode

import "strings"

func reverseWords(s string) string {
	str := strings.Trim(s, " ")  //去除两边空格
	words := strings.Fields(str) //按照空格划分字符串为字符串数组，这里的单词之间的空格即使是多个也会拆分成不带空格的单词，split需要指定按照哪个字符串进行拆分，所以一个空格和两个空格是不同的
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ") //将字符串数组里的字符串按照某个字符串连接起来
}

func reverseWords2(s string) string { //先处理首尾空格，再划分为字符串数组，最后拼接
	if len(s) == 0 {
		return s
	}
	left, right := 0, len(s)-1
	for left < right && s[left] == ' ' {
		left++
	}
	for left < right && s[right] == ' ' {
		right--
	}
	words := []string{}
	s = s[left : right+1]                     //去掉首尾空格
	for i, j := len(s)-1, len(s)-1; i >= 0; { //双指针，从后向前遍历，i先动，找到第一个空格，此时ij之间即为单词，i继续前移找到第一个非空格，跟过来，重新循环
		for i >= 0 && s[i] != ' ' { //在用到数组时注意越界，所以索引的判断尽量写在前面
			i--
		}
		words = append(words, s[i+1:j+1])
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	s = ""
	for i := 0; i < len(words); i++ { //用空格连起来，因为是从后往前遍历，所以单词之间的翻转已经完成
		s = s + words[i] + " "
	}
	return s[:len(s)-1] //去除最后一个空格并输出
}*/
