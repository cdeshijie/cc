/*package leetcode

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for nums[left]%2 != 0 && left < right {
			left += 1
		}
		for nums[right]%2 == 0 && left < right {
			right -= 1
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}*/
