/*package leetcode

type pair struct{
    x int
    y int
}
var directions=[]pair{{-1,0},{1,0},{0,-1},{0,1}}
func exist(board [][]byte, word string) bool {
    row:=len(board)
    if row==0||len(word)==0{
        return false
    }
    col:=len(board[0])
    if col==0{
        return false
    }
    vis:=make([][]bool,row)
    for i:=range vis{
        vis[i]=make([]bool,col)
    }
    var check func(i, j, k int) bool
    check=func(i,j,k int)bool{
        if board[i][j]!=word[k]{
            return false
        }
        if k==len(word)-1{
            return true
        }
        vis[i][j]=true
        defer func(){vis[i][j]=false}()
        for _,dir:=range directions{
            if diri,dirj:=i+dir.x,j+dir.y;diri>=0&&diri<row&&dirj>=0&&dirj<col&&vis[diri][dirj]!=true{
                if check(diri,dirj,k+1)==true{
                    return true
                }
            }
        }
        return false
    }
    for i,v:=range board{
        for j:=range v{
            if check(i,j,0)==true{
                return true
            }
        }
    }
    return false
}*/
