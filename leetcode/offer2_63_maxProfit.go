/*package leetcode
func maxProfit(prices []int) int {//暴力双循环
    max:=0
    for i:=0;i<len(prices)-1;i++{
        for j:=i+1;j<len(prices);j++{
            if prices[j]-prices[i]>max{
                max=prices[j]-prices[i]
            }
        }
    }
    return max
}

func maxProfit(prices []int) int {//贪心
    if len(prices)<=1{
        return 0
    }
    min:=prices[0]
    re:=0
    for i:=1;i<len(prices);i++{
        if re<prices[i]-min{
            re=prices[i]-min
        }
        if prices[i]<min{
            min=prices[i]
        }
    }
    return re
}*/