/*package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode { //申请头结点，然后经典头插法
	if head == nil {
		return nil
	}
	var headnode ListNode
	headnode.Next = nil
	t := head
	for t != nil {
		r := t.Next
		t.Next = headnode.Next
		headnode.Next = t
		t = r
	}
	return headnode.Next
}
func reverseList2(head *ListNode) *ListNode { //用一个pre指针，省了一个结点空间，新开一个以pre为首的指针，思想依然是头插法
	var pre, cur, next *ListNode = nil, head, nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}*/
