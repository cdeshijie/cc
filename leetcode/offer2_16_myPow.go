/*package leetcode


func myPow(x float64, n int) float64 {//每次递归可用的变量是n，x和已经计算好的y
    if n>=0{
        return f(x,n)
    }else{
        return 1/f(x,-n)
    }
}
func f(x float64,n int) float64{
    if n==0{
        return 1
    }
    y:=f(x,n/2)
    if n%2==0{
        return y*y
    }else{
        return y*y*x
    }
}*/