/*package leetcode
func numWays(n int) int {//f(n)=f(n-1)+f(n-2)
    if n<=1{
        return 1
    }
    p,q,r:=1,1,0
    for i:=2;i<=n;i++{
        r=(p+q)%(1e9+7)
        p=q
        q=r
    }
    return r
}*/