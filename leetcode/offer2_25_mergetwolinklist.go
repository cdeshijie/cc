/*package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}



func merge(l1 *ListNode,l2 *ListNode,l *ListNode)*ListNode{//递归合并，先按照顺序走但是不建立连接，后序逐渐返回时，再建立连接
    if l1==nil{
        return l2
    }
    if l2==nil{
        return l1
    }
    if l1.Val<=l2.Val{
        l=l1
        l1=l1.Next
    }else{
        l=l2
        l2=l2.Next
    }
    l.Next=merge(l1,l2,l)
    return l
}



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	left := l1
	right := l2
	head := &ListNode{}
	r := head
	for left != nil && right != nil {
		if left.Val <= right.Val {
			r.Next = left
			left = left.Next
		} else {
			r.Next = right
			right = right.Next
		}
		r = r.Next
	}
	if left != nil {
		r.Next = left
	} else {
		r.Next = right
	}
	return head.Next
}*/
