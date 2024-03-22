/*package leetcode

func minArray(numbers []int) int { //从尾到头找逆序，没有说明顺序，返回第一个元素
	n := len(numbers) - 1
	if n == 0 {
		return numbers[0]
	}
	for n > 0 {
		if numbers[n] < numbers[n-1] {
			return numbers[n]
		}
		n--
	}
	return numbers[0]
}*/
