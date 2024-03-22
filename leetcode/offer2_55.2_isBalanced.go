/*package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool { //判断是否为平衡二叉树，第一种方法，先序遍历，每次看看左右子树高度差
	if root == nil {
		return true
	}
	if abs(high(root.Left)-high(root.Right)) > 1 {
		return false
	}
	if isBalanced(root.Left) == false {
		return false
	}
	if isBalanced(root.Right) == false {
		return false
	}
	return true
}
func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(high(root.Left), high(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}*/
