/*package leetcode
func reversePairs(nums []int) int {//分治算法，小的块之间先算逆序对，合并有序后再算大的块之间的逆序对
    return f(0,len(nums)-1,nums)
}
//L = [8, 12, 16, 22, 100]   R = [9, 26, 55, 64, 91]  M = []
//     |                          |
//   lPtr                       rPtr

//以上图为例，当lptr小于等于rptr代表此时没有逆序对，反之，lptr及以后都大于rptr，均为逆序对

func f(start int,end int,s []int)int{
    if start>=end{//这里不用==判断是因为可能s为[]，此时start=0，end=-1
        return 0
    }
    mid:=(start+end)/2
    re:=f(start,mid,s)+f(mid+1,end,s)//开始分治
    t:=[]int{}//这里写入排序好的元素
    i:=start
    j:=mid+1
    for i<=mid&&j<=end{//这个for循环就把所有逆序对找到了
        if s[i]<=s[j]{
            t=append(t,s[i])
            i++
        }else{
            re+=mid-i+1
            t=append(t,s[j])
        j++
        }
    }
    for i<=mid{
        t=append(t,s[i])
        i++
    }
    for j<=end{
        t=append(t,s[j])
        j++
    }
    i=start
    j=0
    for i<=end{//将排好序的元素写入s，方便下次合并使用
        s[i]=t[j]
        i++
        j++
    }
    return re
}*/