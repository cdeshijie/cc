/*package leetcode
func maxSales(sales []int) int {
    max:=sales[0]
    for i:=1;i<len(sales);i++{
        if sales[i]+sales[i-1]>sales[i]{
            sales[i]+=sales[i-1]
        }
        if sales[i]>max{
            max=sales[i]
        }
    }
    fmt.Println(sales)
    return max
}*/