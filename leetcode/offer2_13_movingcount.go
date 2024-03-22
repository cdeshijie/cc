/*package leetcode
type pair struct{
    x,y int
}
var dirctions =[]pair{{-1,0},{1,0},{0,1},{0,-1}}
func movingCount(m int, n int, k int) int {
    if k==0{
        return 1
    }
    re:=0
    vis:=make([][]bool,m)
    for i:=range vis{
        vis[i]=make([]bool,n)
    }
    var check func(i,j,k int)
    check=func(i,j,k int){
        if sumij(i,j)>k{
            return
        }
        re++
        vis[i][j]=true
        for _,dir:=range dirctions{
            if diri,dirj:=i+dir.x,j+dir.y;0<=diri&&diri<m&&0<=dirj&&dirj<n&&vis[diri][dirj]!=true{
                check(diri,dirj,k)
            }
        }
    }
    check(0,0,k)
    return re
}
func sumij(i,j int)int{
    sum:=0
    for i!=0{
        sum+=i%10
        i=i/10
    }
    for j!=0{
        sum+=j%10
        j=j/10
    }
    return sum
}*/