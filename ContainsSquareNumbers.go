/*******************************************************************************************
Problem 6. Write a function SquareNumbers that takes a list of integers and
returns the numbers in that list that are square numbers (can be written as x*x
for some integer x). The output list should be in the same order as the input list.

Example: If your input is []int{1,5,9,9,20} the returned value should be []int{1,9,9}.
*******************************************************************************************/
func SquareNumbers(a []int) []int {
	var b []int
	var number_float float64
	var squareroot_float float64
	var squareroot_int int

	for x := 0;x<len(a);x++ {
		number_float = float64(a[x])
		squareroot_float = math.Sqrt(number_float)
		squareroot_int = int(squareroot_float)
		if squareroot_float - float64(squareroot_int) == 0 { 
			b = append(b,a[x])
		}
	}
	
	return b
}
