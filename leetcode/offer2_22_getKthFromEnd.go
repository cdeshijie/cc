/*package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	target, _ := returntarget(head, k) //这里只要target
	return target
}

func returntarget(node *ListNode, i int) (*ListNode, int) { //递归到最后一个结点，然后函数每结束一次，i就减一，如果没有指定结点，target一直就是nil，有的话即i==时，改变target
	if node == nil {
		return nil, i - 1
	}
	target, k := returntarget(node.Next, i)
	if k == 0 {
		target = node
	}
	return target, k - 1
}*/

