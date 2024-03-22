/*package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var re *TreeNode = nil
	if root == nil {
		return root
	}
	pathp := []*TreeNode{}
	pathq := []*TreeNode{}
	search(root, p, &pathp)
	search(root, q, &pathq)
	for i := 0; i < len(pathp) && i < len(pathq); i++ {
		if pathp[i] == pathq[i] {
			re = pathp[i]
		} else {
			break
		}
	}
	return re
}
func search(root, target *TreeNode, path *[]*TreeNode) bool { //存储路径
	if root == nil {
		return false
	}
	*path = append(*path, root)
	//如果当前结点为目标或者左右子树已经找到则直接返回，不用再递归
	if root == target || search(root.Left, target, path) == true || search(root.Right, target, path) == true {
		return true
	}
	*path = (*path)[:len(*path)-1] //这里说明目标不在这条路径，删除结点
	return false
}*/
