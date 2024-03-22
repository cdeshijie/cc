/*package leetcode

//判断5张牌是不是顺子，0代表大小王可以是任何值，a不是14，这个题的判断条件在于最大最小值差是否小于5，且切片有没有重复值
//为了方便处理先进行排序

func isStraight(nums []int) bool {
	sort.Slice(nums,func(i,j int)bool{
	  return nums[i]<nums[j]
	})
	min:=nums[0]
	for i:=0;i<len(nums)-1;i++{
	  if nums[i]==0{
		min=nums[i+1]
		continue
	  }
	  if nums[i]==nums[i+1]{
		return false
	  }
	}
	return nums[len(nums)-1]-min<5
  }*/