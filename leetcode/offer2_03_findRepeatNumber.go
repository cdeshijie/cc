/*package leetcode

func findRepeatNumber(nums []int) int { //原地交换，因为题干说明元素都小于切片长度n，所以一定有重复，让元素n去索引为n的位置，如果找到元素n的索引位置n的元素已经是n了，说明重复
	for i := 0; i < len(nums); {
		if nums[i] == i { //这里主注意并非每次都是i++，只有确定这里一定索引和值相等后才后移
			i++
			continue
		} else {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			}
		}
	}
	return -1
}

func findRepeatNumber2(nums []int) int { //用哈希表
	s := make(map[int]int, 0)
	for _, v := range nums {
		_, ok := s[v]
		if ok == true {
			return v
		} else {
			s[v] = 1
		}
	}
	return -1
}*/
