/* package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int { //使用栈
	stack := []int{}
	if head == nil {
		return stack
	}
	for t := head; t != nil; t = t.Next {
		stack = append(stack, t.Val)
	}
	left := 0
	right := len(stack) - 1
	for left < right {
		stack[left], stack[right] = stack[right], stack[left]
		left++
		right--
	}
	return stack
}

func reversePrint2(head *ListNode) []int { //使用递归
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}

func main() {

}
*/