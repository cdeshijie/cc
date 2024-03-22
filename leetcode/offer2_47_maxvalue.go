/*package leetcode

dp(i,j)=   
frame(i,j),i=0,j=0
frame(i,j)+dp(i,j−1),i=0,j!=0
frame(i,j)+dp(i−1,j),i!=0,j=0
frame(i,j)+max[dp(i−1,j),dp(i,j−1),i!=0,j!=0
​dp矩阵的值分为四种情况
  

func jewelleryValue(frame [][]int) int {//动态规划，dp矩阵的值dp[i][j]仅与dp[i-1][j]和dp[i][j-1]有关，所以可以求最大值矩阵dp
    row:=len(frame)
    col:=len(frame[0])
    if row==1&&col==1{
        return frame[0][0]
    }
    //fmt.Println(row,col)
    for i:=0;i<row;i++{
        for j:=0;j<col;j++{
            if i==0&&j==0{
                continue
            }else if i==0&&j!=0{
                frame[i][j]=frame[i][j-1]+frame[i][j]
            }else if i!=0&&j==0{
                frame[i][j]=frame[i-1][j]+frame[i][j]
            }else{
                if frame[i][j-1]>frame[i-1][j]{
                    frame[i][j]=frame[i][j-1]+frame[i][j]
                }else{
                    frame[i][j]=frame[i-1][j]+frame[i][j]
                }
            }
        }
    }
    return frame[row-1][col-1]
}*/