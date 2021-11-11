/*******************************************************************************************
Problem 6. Knock out number. Given a (possibly non-square) matrix `matrix` and
a number x.  Remove all rows and columns that contain x and return the smaller
matrix. For example, if x appears in cell (4,5) and (4,8) then remove row 4 and
columns 5 and 8. Rows and columns start at index 0. Your function can "destroy"
the input matrix.
*******************************************************************************************/
func RemoveRowsColumns(matrix [][]int, x int) [][]int {
	var row []int //list to mark rows for deletion
	var col []int //list to mark cols for deletion
	var flag int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[1]); j++ {
			flag = 0
			if matrix[i][j] == x {
				for row_value := 0; row_value < len(row); row_value++ {
					if row[row_value] == i {
						flag = 1
						break
					}
				}
				if len(row) == 0 {
					row = append(row, i) // = true //matrix = removeRows(matrix,i)
				} else if flag == 0 {
					row = append(row, i)
				}
			}
		}
	}
	for j := 0; j < len(matrix); j++ {  //rows
		for i := 0; i < len(matrix[1]); i++ { //columns
			flag = 0
			if matrix[j][i] == x { //column, rows
				for col_value := 0; col_value < len(col); col_value++ {
					if col[col_value] == j {
						flag = 1
						break
					}
				}
				if len(col) == 0 {
					col = append(col, j) // = true //matrix = removeRows(matrix,i)
				} else if flag == 0 {
					col = append(col, j)
				}
			}
		}
	}

	for row_value := 0; row_value < len(row); row_value++ {
		matrix = RemoveRows(matrix, row[row_value]-row_value)
	}
	for column_value := 0; column_value < len(col); column_value++ {
		matrix = RemoveColumns(matrix, col[column_value]-column_value)
	}
	return matrix
}

//RemoveRows takes the matrix and the rows to append and return them
func RemoveRows(matrix [][]int, row int) [][]int {
	matrix = append(matrix[:row], matrix[row+1:]...)
	return matrix
}

////RemoveRows takes the matrix and the columns to append and return them
func RemoveColumns(matrix [][]int, column int) [][]int {
	for i := 0; i < len(matrix); i++ {
		matrix[i] = append(matrix[i][:column], matrix[i][column+1:]...)
	}
	return matrix
}

func InputQ6(){
	a := [][]int{  
	   {0, 3, 2, 1} ,   /*  initializers for row indexed by 0 */
	   {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
	   {8, 1, 10, 5},   /*  initializers for row indexed by 2 */
	   {1, 3, 1, 11},   /*  initializers for row indexed by 3 */
	}
	value := 0
	fmt.Println(RemoveRowsColumns(a,value))
}
