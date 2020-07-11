package algo

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数
func FindNumberIn2DArray(matrix [][]int, target int) bool {

	var x,y int
	for i := 0; i < len(matrix); i++ {
		if target >= matrix[i][len(matrix[i])] {
			x = i
			break
		}
	}

	return true
}
