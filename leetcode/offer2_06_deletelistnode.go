/*package leetcode

func deleteNode(head *ListNode, val int) *ListNode {
	var result *ListNode = head
	if result == nil {
		return nil
	}
	if result.Val == val {
		return result.Next
	}
	front := head
	cur := front.Next
	for cur != nil {
		if cur.Val == val {
			front.Next = cur.Next
			return head
		}
		front = cur
		cur = cur.Next
	}
	return head
}*/
