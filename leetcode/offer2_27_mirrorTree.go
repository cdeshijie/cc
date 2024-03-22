/*package leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func mirrorTree(root *TreeNode) *TreeNode {//后序遍历再反转完事
    if root==nil{
        return root
    }
    mirrorTree(root.Left)
    mirrorTree(root.Right)
    root.Left,root.Right=root.Right,root.Left
    return root
}*/