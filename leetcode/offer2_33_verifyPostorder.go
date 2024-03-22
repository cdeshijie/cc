/*package leetcode

func verifyPostorder(postorder []int) bool {
    if len(postorder)<=2{//如果序列小于等于2，此时一定是某个二叉排序树的后序遍历
        return true
    }
    m:=0
    for postorder[m]<postorder[len(postorder)-1]{//找到第一个比根节点大的值，从此值开始为右子树，之前为左子树，因为循环条件，这里可以确定左子树都小于根节点
        m++
    }
    n:=m
    for postorder[n]>postorder[len(postorder)-1]{//判断右子树的值是否都大于根节点
        n++
    }
    return n==len(postorder)-1&&verifyPostorder(postorder[:m])&&verifyPostorder(postorder[m:len(postorder)-1])//n的判断是右子树是否有值小于根节点，有就是false，无序进入深次递归
}
//判断该后序遍历是否为某二叉搜索树的后序遍历，后序遍历最右边的结点就是根节点，其实就是看左右子树是否为二叉排序树，左右是，在判断左子树是否都小于根节点，右子树是否都大于根节点*/