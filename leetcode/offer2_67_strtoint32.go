/* package main

import (
	"fmt"
	"math"
	"strings"
)

func strToInt(str string) int { //这道题题目描述不太好，很多情况没说清楚，但是合法的情况是非空首字符为+，-或者数字，然后开始计数，知道一个非数字的字母或者到结束，返回结果
	if len(str) == 0 {
		return 0
	}
	s := strings.TrimLeft(str, " ")
	result := 0
	sign := 1
	for k, v := range s {
		if v == '+' && k == 0 {
			continue
		} else if v == '-' && k == 0 {
			sign = -1
		} else if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else {
			break
		}
		fmt.Printf("result: %v\n", result)
		if result > math.MaxInt32 { //这里为了防止越界，int32要比字符串小很多，所以在累加过程中只要大于int32max就停止
			if sign == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
	}
	return result * sign
}

func main() {
	v := strToInt("9223372036854775808")
	fmt.Printf("v: %v\n", v)
}
*/