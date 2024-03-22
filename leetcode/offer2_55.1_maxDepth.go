/*package leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }
    high,re:=0,0
    var check func(root *TreeNode)
    check=func(root *TreeNode){
        if root==nil{
            return
        }
        high++
        if high>re{
            re=high
        }
        check(root.Left)
        check(root.Right)
        high--
    }
    check(root)
    return re
}*/