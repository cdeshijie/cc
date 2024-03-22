/*package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool { //题目不难，顺序为先序遍历，之后比较当前然后比较左子树和目标，右子树和目标
	if A == nil || B == nil {
		return false
	}
	return campare(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func campare(A *TreeNode, B *TreeNode) bool { //判断两个树是否相等
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return campare(A.Left, B.Left) && campare(A.Right, B.Right)
}*/
