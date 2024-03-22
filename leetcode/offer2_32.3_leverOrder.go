/*package leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
	re:=[][]int{}
 if root==nil{
	 return re
 }
 level:=1//偶数从左到右
 queue:=[]*TreeNode{root}
 tre:=[]int{root.Val}
 re=append(re,tre)
 for len(queue)!=0{
	 j:=len(queue)
	 tre=[]int{}
	 for i:=0;i<j;i++{
		 t:=queue[0]
		 queue=queue[1:]
		 if t.Left!=nil{
			 queue=append(queue,t.Left)
			 tre=append(tre,t.Left.Val)
		 }
		 if t.Right!=nil{
			 queue=append(queue,t.Right)
			 tre=append(tre,t.Right.Val)
		 }
	 }
	 if len(tre)!=0{
		 if level%2==1{
			 for i:=0;i<len(tre)/2;i++{
			 tre[i],tre[len(tre)-1-i]=tre[len(tre)-1-i],tre[i]
			 }
		 }
		 re=append(re,tre)
	 }
	 level++
 }
 return re
}*/