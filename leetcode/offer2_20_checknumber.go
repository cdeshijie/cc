/* package main

import (
	"fmt"
	"strings"
)

func checkNumber(s string) bool {
	/*根据题意可以判断，字符串由空格，数字，+/-号，小数点，e/E以及剩下的字符共6种
	  其中只要非前5种字符就一定不是数字，我们从几种字符之间的关系和字符的位置进行判断字符串是否为数字，所以字符串起码要长度为2
	  首先对空字符串和单字符串进行处理；
	  然后判断字符串非法：
	  空格：遇到就非法
	  整数：头尾都合法，前面出现什么字符均合法，身后一位除了是+/-都合法
	  +/-：头合法，尾不合法，因为有e/E的存在前面出现+/-号，小数点，e/E都无法判断，后一位不能是+/-和e/E
	  .：.8 8. 8.8为合法数字，所以头尾都可以，前面不能出现e/E和.字符，后一位不能是+/-
	  e/E：头尾都不合法，前面必须出现过整数,前面出现e/E非法,后一位不能是.
	  其他字符：遇到就非法

	//这里不进行+-是否出现做判断，因为任何字符前面出现过+-号都不能判断出是否非法
	inted := false
	doted := false
	eEed := false
	if len(s) == 0 {
		return false
	}
	s1 := strings.Trim(s, " ")
	if len(s1) == 0 {
		return false
	}
	if len(s1) == 1 {
		if s1[0] >= '0' && s1[0] <= '9' {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] >= '0' && s1[i] <= '9' {
			if i != len(s1)-1 && (s[i+1] == '+' || s[i+1] == '-') {
				fmt.Printf("i: %v\n", i)
				return false
			}
			inted = true
		} else if s1[i] == '+' || s1[i] == '-' {
			if i == len(s1)-1 || i != 0 && s1[i-1] != 'e' && s1[i-1] != 'E' {
				fmt.Printf("i: %v\n", i)
				return false
			}
		} else if s1[i] == '.' {
			if doted == true || eEed == true || (i != len(s1)-1 && (s1[i+1] == '+' || s1[i+1] == '-')) {
				fmt.Printf("i: %v\n", i)
				return false
			}
			doted = true
		} else if s1[i] == 'e' || s1[i] == 'E' {
			if inted == false || eEed == true || i == 0 || i == len(s1)-1 || s[i+1] == '.' {
				fmt.Printf("i: %v\n", i)
				return false
			}
			eEed = true
		} else {
			return false
		}
	}
	if inted == false {
		return false
	}
	return true
}

func main() {
	tf := checkNumber(" 005047e+6")
	fmt.Printf("tf: %v\n", tf)
}
*/