/*package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	result := make([]int, 0)
	queue = append(queue, root)
	for i := 0; i < len(queue); i++ {
		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}
		result = append(result, queue[i].Val)
	}
	return result
}*/
