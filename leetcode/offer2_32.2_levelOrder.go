/*package leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	re := [][]int{[]int{root.Val}}
	for len(queue) != 0 {
		j := len(queue)
		tq := []int{}
		for i := 0; i < j; i++ {
			t := queue[0]
			queue = queue[1:]
			if t.Left != nil {
				queue = append(queue, t.Left)
				tq = append(tq, t.Left.Val)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
				tq = append(tq, t.Right.Val)
			}
		}
		if len(tq) != 0 {
			re = append(re, tq)
		}
	}
	return re
}*/