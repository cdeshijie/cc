/*package leetcode
func twoSum(nums []int, target int) []int {
    left,right:=0,len(nums)-1
    for left<right{
        if nums[left]+nums[right]<target{
            left+=1
        }else if nums[left]+nums[right]>target{
            right-=1
        }else{
            return []int{nums[left],nums[right]}
        }
    }
    return []int{}
}*/