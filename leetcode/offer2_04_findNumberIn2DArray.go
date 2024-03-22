/*package leetcode
func findNumberIn2DArray(matrix [][]int, target int) bool {//这里二维矩阵行列有序，所以可以逆时针旋转矩阵45度，此时矩阵类似于二叉排序树
    if len(matrix)==0||len(matrix[0])==0{
        return false
    }
    //row:=len(matrix)
    row:=0
    col:=len(matrix[0])-1
    points:=matrix[row][col]
    for{
        if points==target{
            return true
        }
        if points<target{
            row++
            if row>len(matrix)-1{
                return false
            }
        }else{
            col--
            if col<0{
                return false
            }
        }
        points=matrix[row][col]
    }
}*/