/*package leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func isSymmetric(root *TreeNode) bool {
    if root==nil{//空结点认为为对称
        return true
    }
    return compare(root.Left,root.Right)
}

func compare(A *TreeNode,B *TreeNode) bool{
    if A==nil&&B==nil{//对应结点都是空结点
        return true
    }
    if A!=nil&&B!=nil{
        if A.Val!=B.Val{
            return false
        }
        return compare(A.Left,B.Right)&&compare(A.Right,B.Left)//A结点左比B结点的右，A结点右比B结点左
    }
    return false//两个结点一个空，一个不为空
}*/