/*package leetcode

func spiralOrder(matrix [][]int) []int { //顺序打印矩阵，定义矩阵四角
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row := len(matrix)
	col := len(matrix[0])
	result := make([]int, 0, row*col)
	top := 0
	left := 0
	right := col - 1
	bottom := row - 1
	for top <= bottom && left <= right { //循环打印
		for cols := left; cols <= right; cols++ { //从左到右
			result = append(result, matrix[top][cols])
		}
		for rows := top + 1; rows <= bottom; rows++ { //从上到下，注意第一个已经被打印，这里rows=top+1
			result = append(result, matrix[rows][right])
		}
		if right-left >= 1 && bottom-top >= 1 { //保证矩阵不是一条，而是有多行，如果矩阵只剩竖着或者横着的一条，上面两行足以打印
			for cols := right - 1; cols >= left; cols-- {
				result = append(result, matrix[bottom][cols])
			}
			for rows := bottom - 1; rows > top; rows-- {
				result = append(result, matrix[rows][left])
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return result
}*/
