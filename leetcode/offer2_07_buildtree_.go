/*package leetcode
//根据先序和中序构建二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {//根据根节点不断划分左子树和右子树，直到子树为叶节点
	if len(preorder) == 0 {
		return nil
	}
	i := 0
	for ; i < len(inorder); i++ {//找到根节点在中序的位置
		if inorder[i] == preorder[0] {
			break
		}
	}
	root := &TreeNode{Val: preorder[0],//构造根节点
		Left:  nil,
		Right: nil}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])//传入左子树的先序和中序
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])//传入右子树的先序和中序
	return root
}*/
