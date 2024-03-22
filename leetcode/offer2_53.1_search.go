/*package leetcode

func search(nums []int, target int) int { //寻找元素的两个边界
    left:=0
    right:=len(nums)-1
    for left<=right{//寻找右边界
        mid:=(left+right)/2
        if target<nums[mid]{
            right=mid-1
        }else{
            left=mid+1
        }
    }
    if right<0||nums[right]!=target{//循环完成后，right应该指向右边界，right小于0说明数组都大于target，如果nums[right]!=target，说明数组没有这个元素
        return 0
    }
    j:=right
    left=0
    for left<=right{//只重置left，因为到这里right一定等于target
        mid:=(left+right)/2
        if target>nums[mid]{
            left=mid+1
        }else{
            right=mid-1
        }
    }
    i:=left//相同道理
    return j-i+1
}*/