/*package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}
//假设未相交A部分长度为a，B长度部分为b，相交部分长度为c，让a走完A接着走b，b走完B接着走a，如果A和B长度相同且有公共结点，则没走完一个链表时就会遇到
//假如AB链表长度不同，假设A小于B，则A走a+c后，转移到B继续走，b此时继续走B，在a走了a+c+b时，b的次数为b+c+a,此时要么为nil，要么就是公共结点
//假如AB不想交，也可以把AB看做最后一个结点是公共nil结点，此时AB肯定有公共结点，不是nil就是中间相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {//看力扣解析，这个需要图和证明才好理解
	pa := headA
	pb := headB
	if pa == nil || pb == nil {
		return nil
	}
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}*/
