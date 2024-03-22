/*package leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func kthLargest(root *TreeNode, k int) int {//第一种做法，因为是二叉排序树，所以直接中序遍历然后此时即是有序，返回想要的值
    result:=[]int{}
    var check func(root *TreeNode)
    check=func(root *TreeNode){
        if root.Left!=nil{
            check(root.Left)
        }
        result=append(result,root.Val)
        if root.Right!=nil{
            check(root.Right)
        }
    }
    check(root)
    return result[len(result)-k]
}*/
/*func kthLargest(root *TreeNode, k int) int {//第二种，先序遍历所有结点，然后内部冒泡排序，返回想要的值，离谱的是这样竟然要快
    result:=[]int{}
    var check func(root *TreeNode)
    check=func(root *TreeNode){
        result=append(result,root.Val)
        if root.Left!=nil{
            check(root.Left)
        }
        if root.Right!=nil{
            check(root.Right)
        }
    }
    check(root)
    maopao(result)
    return result[len(result)-k]
}
func maopao(s []int){
    for i:=0;i<len(s)-1;i++{
        flag:=0
        for j:=0;j<len(s)-1-i;j++{
            if s[j]>s[j+1]{
                s[j],s[j+1]=s[j+1],s[j]
                flag=1
            }
        }
        if flag==0{
            break
        }
    }
}*/