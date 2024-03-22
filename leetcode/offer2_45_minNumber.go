/*package leetcode

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) string {//我的做法，全排列字符串数组，把每次得到一个结果就和最小值比较，取较小值
	if len(nums) <= 1 {
		if len(nums) == 1 {
			return strconv.Itoa(nums[0])
		} else {
			return ""
		}
	}
	var re, t string
	s := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {//把整数变为字符串
		s[i] = strconv.Itoa(nums[i])
		re = re + s[i]
	}
	var f func([]string)
	visit := make([]bool, len(s))
	f = func(s []string) {//全排列
		if len(t) == len(re) {
			if compare(re, t) == true { //比较大小
				re = t
			}
		}
		for i := range s {
			if visit[i] == true {
				continue
			}
			t += s[i]
			visit[i] = true
			f(s)
			t = t[:len(t)-len(s[i])]
			visit[i] = false
		}
	}
	f(s)
	return re
}
func compare(a, b string) bool {//字符串比较
	t1 := strings.TrimLeft(a, "0")
	t2 := strings.TrimLeft(b, "0")
	if len(t1) != len(t2) {
		if len(t1) > len(t2) {
			return true
		} else {
			return false
		}
	}
	return t1 > t2
}

func minNumber(nums []int) string {//对于字符串而言，xy均为int，在xy字符串化后,如果x+y>y+x,则认为x>y。这里使用冒泡排序
    re:=""
    if len(nums)==0{
        return re
    }
    for i:=0;i<len(nums)-1;i++{
        flag:=0
        for j:=0;j<len(nums)-1-i;j++{
            if strconv.Itoa(nums[j])+strconv.Itoa(nums[j+1])>strconv.Itoa(nums[j+1])+strconv.Itoa(nums[j]){
                nums[j],nums[j+1]=nums[j+1],nums[j]
                flag=1
            }
        }
        if flag==0{
            break
        }
    }
    for i:=range nums{
        re+=strconv.Itoa(nums[i])
    }
    return re
}

func minNumber(nums []int) string {//另一种排序的写法，使用go自带的sort包，对切片进行排序，nums是要排序的切片，func里是切片两个元素比较的规则,x<y从小到大，x>y从大到小
    //将整数数组按字符串形式排序
    sort.Slice(nums, func(i, j int) bool {
        x := fmt.Sprintf("%d%d", nums[i], nums[j])
        y := fmt.Sprintf("%d%d", nums[j], nums[i])
        return x < y
    })

    res := ""
    for i:=0; i<len(nums); i++ {
        res += fmt.Sprintf("%d", nums[i])
    }
    return res
}


*/
