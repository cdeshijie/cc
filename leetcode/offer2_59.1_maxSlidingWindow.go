/*package leetcode

func maxSlidingWindow(nums []int, k int) []int {//原理看b站卡尔讲的视频，具体就是在新建队列时，把队列里小于新入列元素的值去除，因为新元素在队列时在新元素前还小于新元素的永远不会成为最大值
    s:=make([]int,1,k)//cap限制为k后内存击败直接提升30多
    max:=make([]int,0,len(nums)-k+1)
    push:=func(i int){
        for len(s)!=0&&s[len(s)-1]<i{
            s=s[:len(s)-1]
        }
        s=append(s,i)
    }
    for i:=0;i<len(nums);i++{
        push(nums[i])
        if i<k-1{
            continue
        }
        if i==k-1{
            max=append(max,s[0])
            continue
        }
        if nums[i-k]==s[0]{
            s=s[1:]
        }
        max=append(max,s[0])
    }
    return max
}*/
