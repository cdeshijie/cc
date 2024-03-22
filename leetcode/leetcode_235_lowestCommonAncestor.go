/*package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {//先序找到pq的路径在判断
    var re *TreeNode=nil
	if root==nil{
        return re
    }
    wayp:=[]*TreeNode{}
    wayq:=[]*TreeNode{}
    for i:=root;i!=p;{
        wayp=append(wayp,i)
        if p.Val<i.Val{
            i=i.Left
        }else{
            i=i.Right
        }
    }
    wayp=append(wayp,p)
    for i:=root;i!=q;{
        wayq=append(wayq,i)
        if q.Val<i.Val{
            i=i.Left
        }else{
            i=i.Right
        }
    }
    wayq=append(wayq,q)
    for i:=0;i<len(wayp)&&i<len(wayq);i++{
        if wayp[i]==wayq[i]{
            re=wayp[i]
        }else{
            break
        }
    }
    return re
}*/