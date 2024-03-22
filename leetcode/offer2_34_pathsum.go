/*package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	stack := []int{}
	var search func(root *TreeNode)
	search = func(root *TreeNode) {
		stack = append(stack, root.Val)
		if root.Left == nil && root.Right == nil {
			if sum(stack) == target {
				t := make([]int, len(stack))
				copy(t, stack)
				result = append(result, t) //如果是就添加入result
			}
		}
		if root.Left != nil {
			search(root.Left)
		}
		if root.Right != nil {
			search(root.Right)
		}
		stack = stack[:len(stack)-1]
	}
	search(root)
	return result
}

func sum(stack []int) int {
	if len(stack) == 0 {
		return 0
	}
	s := 0
	for i := 0; i < len(stack); i++ {
		s += stack[i]
	}
	return s
}
*/