package ac_004

//二分查找，搜索二位矩阵中的值，左至右递增，上到下递增

//1.暴力解法
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	for _, v1 := range matrix {
		for _, v2 := range v1 {
			if target == v2 {
				return true
			}
		}
	}
	return false
}

//二分查找
func findNumberIn2DArray2(matrix [][]int,target int) bool  {
	if len(matrix) == 0 {
		return false
	}
	//初始化到最左下角的元素位置
	rows := len(matrix)
	cols := len(matrix[0])
	i,j := rows-1,0
	//i只能减,j只能增，进而转化为二分查找
	for i >=0 && j < cols {
		if matrix[i][j] < target {
			j ++
		}else if matrix[i][j] > target{
			i --
		}else {
			return true
		}
	}
	return false
}
